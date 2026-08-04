package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	defaultLogSizeGB      float64 = 20
	bytesPerGB                    = int64(1024 * 1024 * 1024)
	defaultStatusInterval         = time.Second
	writeBufferSize               = 1024 * 1024
	defaultDataDirectory          = "practice/file_t"
)

var simulatedNodeIDs = []string{
	"api-node-01", "api-node-02", "api-node-03", "api-node-04", "api-node-05",
	"api-node-06", "api-node-07", "api-node-08", "api-node-09", "api-node-10",
	"api-node-11", "api-node-12", "api-node-13", "api-node-14", "api-node-15",
	"api-node-16", "api-node-17", "api-node-18", "api-node-19", "api-node-20",
	"api-node-21", "api-node-22", "api-node-23", "api-node-24", "api-node-25",
	"api-node-26", "api-node-27", "api-node-28", "api-node-29", "api-node-30",
}

// simulatorConfig 保存日志模拟器的运行参数。
type simulatorConfig struct {
	outputDir      string
	totalSizeBytes int64
	overwrite      bool
	statusInterval time.Duration
	seed           int64
	statusWriter   io.Writer
}

// logSpec 描述一种日志的文件名、容量和日志级别。
type logSpec struct {
	fileName string
	target   int64
	kind     string
}

// logFileWriter 负责一个日志文件的缓冲写入和字节统计。
type logFileWriter struct {
	spec    logSpec
	path    string
	file    *os.File
	writer  *bufio.Writer
	written int64
}

// logSimulator 按线上常见格式生成三类日志文件。
type logSimulator struct {
	config simulatorConfig
	random *rand.Rand
	now    time.Time
}

// runGenerator 解析生成参数并在收到终止信号时安全结束生成任务。
func (a *application) runGenerator(arguments []string) error {
	config, err := parseConfig(arguments, a.output)
	if errors.Is(err, flag.ErrHelp) {
		printGenerateUsage(a.output)
		return nil
	}
	if err != nil {
		printGenerateUsage(a.errorOutput)
		return fmt.Errorf("参数错误: %w", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := newLogSimulator(config).run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("生成日志失败: %w", err)
	}
	return nil
}

// parseConfig 将命令行参数转换为经过校验的模拟器配置。
func parseConfig(arguments []string, statusWriter io.Writer) (simulatorConfig, error) {
	flags := flag.NewFlagSet("read_t generate", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	outputDir := flags.String("output-dir", "", "日志输出目录，默认根据日志总大小自动生成")
	sizeGB := flags.Float64("size-gb", defaultLogSizeGB, "日志总大小（GiB，可为小数，例如 0.5）")
	overwrite := flags.Bool("overwrite", false, "允许覆盖已有日志文件")
	statusInterval := flags.Duration("status-interval", defaultStatusInterval, "进度刷新周期，例如 1s")
	seed := flags.Int64("seed", time.Now().UnixNano(), "随机种子")
	if err := flags.Parse(arguments); err != nil {
		return simulatorConfig{}, err
	}
	if flags.NArg() != 0 {
		return simulatorConfig{}, fmt.Errorf("不支持的位置参数: %s", strings.Join(flags.Args(), " "))
	}
	maxLogSizeGB := float64(math.MaxInt64 / bytesPerGB)
	if math.IsNaN(*sizeGB) || math.IsInf(*sizeGB, 0) || *sizeGB <= 0 || *sizeGB > maxLogSizeGB {
		return simulatorConfig{}, fmt.Errorf("size-gb 必须大于 0 且不超过 %g，可使用小数，例如 0.5", maxLogSizeGB)
	}
	totalSizeBytes := int64(math.Round(*sizeGB * float64(bytesPerGB)))
	if totalSizeBytes <= 0 {
		return simulatorConfig{}, errors.New("size-gb 换算后的日志总大小必须大于 0 字节")
	}
	resolvedOutputDir := *outputDir
	if strings.TrimSpace(resolvedOutputDir) == "" {
		resolvedOutputDir = defaultOutputDir(*sizeGB)
	}

	config := simulatorConfig{
		outputDir:      resolvedOutputDir,
		totalSizeBytes: totalSizeBytes,
		overwrite:      *overwrite,
		statusInterval: *statusInterval,
		seed:           *seed,
		statusWriter:   statusWriter,
	}
	return config, config.validate()
}

// defaultOutputDir 根据日志总大小生成 read_t 目录中的默认日志目录。
func defaultOutputDir(sizeGB float64) string {
	sizeName := strconv.FormatFloat(sizeGB, 'f', -1, 64)
	return filepath.Join(defaultDataDirectory, "logs"+sizeName)
}

// validate 检查配置是否可安全用于生成日志。
func (c simulatorConfig) validate() error {
	if strings.TrimSpace(c.outputDir) == "" {
		return errors.New("output-dir 不能为空")
	}
	if c.totalSizeBytes <= 0 {
		return errors.New("日志总大小必须大于 0")
	}
	if c.statusInterval <= 0 {
		return errors.New("status-interval 必须大于 0")
	}
	if c.statusWriter == nil {
		return errors.New("状态输出不可为空")
	}
	return nil
}

// newLogSimulator 使用配置创建一个可复现的日志模拟器。
func newLogSimulator(config simulatorConfig) *logSimulator {
	return &logSimulator{
		config: config,
		random: rand.New(rand.NewSource(config.seed)),
		now:    time.Now().UTC().Add(-24 * time.Hour),
	}
}

// run 创建日志文件、持续写入记录，并定期报告每个文件的 MB 大小。
func (s *logSimulator) run(ctx context.Context) (runErr error) {
	if err := s.config.validate(); err != nil {
		return err
	}
	if err := os.MkdirAll(s.config.outputDir, 0o755); err != nil {
		return fmt.Errorf("创建日志目录: %w", err)
	}

	specs := s.buildSpecs()
	if err := s.ensureTargetsAreSafe(specs); err != nil {
		return err
	}

	writers := make([]*logFileWriter, 0, len(specs))
	defer func() {
		for _, writer := range writers {
			if err := writer.close(); runErr == nil && err != nil {
				runErr = err
			}
		}
		s.writeStatus("最终", writers)
	}()

	for _, spec := range specs {
		writer, err := newLogFileWriter(s.config.outputDir, spec, s.config.overwrite)
		if err != nil {
			return err
		}
		writers = append(writers, writer)
	}

	s.writeStatus("开始", writers)
	nextStatusAt := time.Now().Add(s.config.statusInterval)

	for !allFilesComplete(writers) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		for _, writer := range writers {
			if err := s.writeNextRecord(writer); err != nil {
				return err
			}
		}
		if time.Now().Before(nextStatusAt) {
			continue
		}
		if err := flushWriters(writers); err != nil {
			return err
		}
		s.writeStatus("进行中", writers)
		nextStatusAt = time.Now().Add(s.config.statusInterval)
	}

	return nil
}

// buildSpecs 按访问、信息、错误日志的 80%、15%、5% 容量比例创建目标。
func (s *logSimulator) buildSpecs() []logSpec {
	accessSize := s.config.totalSizeBytes * 80 / 100
	infoSize := s.config.totalSizeBytes * 15 / 100
	return []logSpec{
		{fileName: "access.log", target: accessSize, kind: "access"},
		{fileName: "info.log", target: infoSize, kind: "info"},
		{fileName: "error.log", target: s.config.totalSizeBytes - accessSize - infoSize, kind: "error"},
	}
}

// ensureTargetsAreSafe 防止在未明确授权时覆盖已有日志。
func (s *logSimulator) ensureTargetsAreSafe(specs []logSpec) error {
	if s.config.overwrite {
		return nil
	}
	for _, spec := range specs {
		path := filepath.Join(s.config.outputDir, spec.fileName)
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("日志文件已存在: %s；如需覆盖请使用 --overwrite", path)
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("检查日志文件 %s: %w", path, err)
		}
	}
	return nil
}

// newLogFileWriter 打开一个待写入日志文件，并以截断方式从头生成内容。
func newLogFileWriter(outputDir string, spec logSpec, overwrite bool) (*logFileWriter, error) {
	path := filepath.Join(outputDir, spec.fileName)
	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !overwrite {
		flags |= os.O_EXCL
	}
	file, err := os.OpenFile(path, flags, 0o644)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件 %s: %w", path, err)
	}
	return &logFileWriter{
		spec:   spec,
		path:   path,
		file:   file,
		writer: bufio.NewWriterSize(file, writeBufferSize),
	}, nil
}

// writeNextRecord 向未完成的文件写入一条完整且不超出目标容量的日志。
func (s *logSimulator) writeNextRecord(writer *logFileWriter) error {
	remaining := writer.spec.target - writer.size()
	if remaining == 0 {
		return nil
	}

	record := s.nextRecord(writer.spec.kind)
	minimumFinalSize := int64(len(s.finalRecord(writer.spec.kind, 0)))
	if remaining <= int64(len(record))+minimumFinalSize {
		record = s.finalRecord(writer.spec.kind, int(remaining-minimumFinalSize))
	}
	return writer.write(record)
}

// nextRecord 生成一条带有常见线上字段的普通日志记录。
func (s *logSimulator) nextRecord(kind string) []byte {
	s.now = s.now.Add(time.Duration(s.random.Intn(900)+100) * time.Microsecond)
	timestamp := s.now.Format(time.RFC3339Nano)
	requestID := fmt.Sprintf("%016x", s.random.Uint64())
	nodeID := s.randomNodeID()

	switch kind {
	case "access":
		method := []string{"GET", "GET", "POST", "PUT", "DELETE"}[s.random.Intn(5)]
		path := []string{"/api/v1/orders", "/api/v1/users/me", "/api/v1/login", "/healthz", "/assets/app.js"}[s.random.Intn(5)]
		status := s.accessStatus()
		return []byte(fmt.Sprintf("%s INFO [access] access_time=%s node_id=%s request_id=%s remote_addr=10.%d.%d.%d method=%s path=%q status=%d bytes=%d duration_ms=%d user_agent=%q\n", timestamp, timestamp, nodeID, requestID, s.random.Intn(255), s.random.Intn(255), s.random.Intn(255), method, path, status, s.random.Intn(128000)+128, s.random.Intn(900)+1, "Mozilla/5.0"))
	case "info":
		component := []string{"http", "worker", "cache", "database", "scheduler"}[s.random.Intn(5)]
		message := []string{"请求处理完成", "缓存命中", "任务执行完成", "连接池状态正常", "配置热加载完成"}[s.random.Intn(5)]
		status := []int{200, 200, 201, 202, 204}[s.random.Intn(5)]
		return []byte(fmt.Sprintf("%s INFO access_time=%s node_id=%s service=order-api component=%s request_id=%s status=%d duration_ms=%d message=%q\n", timestamp, timestamp, nodeID, component, requestID, status, s.random.Intn(500)+1, message))
	default:
		component := []string{"database", "auth", "gateway", "rate-limit", "worker"}[s.random.Intn(5)]
		errorMessage := []string{"upstream request timeout", "database connection reset", "token verification failed", "rate limit exceeded", "message consume failed"}[s.random.Intn(5)]
		return []byte(fmt.Sprintf("%s ERROR access_time=%s node_id=%s service=order-api component=%s request_id=%s status=%d duration_ms=%d caller=internal/%s.go:%d error=%q retryable=%t\n", timestamp, timestamp, nodeID, component, requestID, s.errorStatus(), s.random.Intn(3000)+10, component, s.random.Intn(180)+20, errorMessage, s.random.Intn(100) < 70))
	}
}

// finalRecord 生成可填充的末尾日志，确保文件大小精确且最后一行仍为合法日志。
func (s *logSimulator) finalRecord(kind string, paddingLength int) []byte {
	level := "INFO"
	status := 200
	if kind == "error" {
		level = "ERROR"
		status = 500
	}
	message := "生成完成"
	if paddingLength > 0 {
		message += strings.Repeat("x", paddingLength)
	}
	timestamp := s.now.Format(time.RFC3339Nano)
	return []byte(fmt.Sprintf("%s %s access_time=%s node_id=api-node-01 service=order-api component=generator status=%d duration_ms=1 message=%q\n", timestamp, level, timestamp, status, message))
}

// randomNodeID 从模拟服务器节点池中随机选择一个节点名称。
func (s *logSimulator) randomNodeID() string {
	return simulatedNodeIDs[s.random.Intn(len(simulatedNodeIDs))]
}

// accessStatus 按常见线上比例返回 HTTP 状态码。
func (s *logSimulator) accessStatus() int {
	roll := s.random.Intn(1000)
	switch {
	case roll < 920:
		return []int{200, 200, 200, 201, 204}[s.random.Intn(5)]
	case roll < 985:
		return []int{400, 401, 403, 404, 429}[s.random.Intn(5)]
	default:
		return []int{500, 502, 503}[s.random.Intn(3)]
	}
}

// errorStatus 返回适用于错误日志的客户端或服务端失败状态码。
func (s *logSimulator) errorStatus() int {
	return []int{400, 401, 403, 404, 429, 500, 502, 503}[s.random.Intn(8)]
}

// writeStatus 以 MB 为单位输出每个日志文件和总量的当前大小。
func (s *logSimulator) writeStatus(stage string, writers []*logFileWriter) {
	var total int64
	parts := make([]string, 0, len(writers))
	for _, writer := range writers {
		size := writer.size()
		total += size
		parts = append(parts, fmt.Sprintf("%s=%.2fMB", writer.spec.fileName, float64(size)/(1024*1024)))
	}
	percent := float64(total) / float64(s.config.totalSizeBytes) * 100
	fmt.Fprintf(s.config.statusWriter, "[%s] %s total=%.2fMB/%.2fMB (%.2f%%)\n", stage, strings.Join(parts, " "), float64(total)/(1024*1024), float64(s.config.totalSizeBytes)/(1024*1024), percent)
}

// write 将一条完整记录写入缓冲区，并更新已写入字节数。
func (w *logFileWriter) write(record []byte) error {
	n, err := w.writer.Write(record)
	w.written += int64(n)
	if err != nil {
		return fmt.Errorf("写入日志文件 %s: %w", w.path, err)
	}
	if n != len(record) {
		return fmt.Errorf("写入日志文件 %s: %w", w.path, io.ErrShortWrite)
	}
	return nil
}

// size 返回当前已写入缓冲区的日志字节数。
func (w *logFileWriter) size() int64 {
	return w.written
}

// flush 将缓冲区内容写入文件，使状态输出对应实际文件大小。
func (w *logFileWriter) flush() error {
	if err := w.writer.Flush(); err != nil {
		return fmt.Errorf("刷新日志文件 %s: %w", w.path, err)
	}
	return nil
}

// close 刷新缓冲数据、同步文件并关闭文件描述符。
func (w *logFileWriter) close() (closeErr error) {
	if err := w.flush(); err != nil {
		closeErr = err
	}
	if err := w.file.Sync(); err != nil && closeErr == nil {
		closeErr = fmt.Errorf("同步日志文件 %s: %w", w.path, err)
	}
	if err := w.file.Close(); err != nil && closeErr == nil {
		closeErr = fmt.Errorf("关闭日志文件 %s: %w", w.path, err)
	}
	return closeErr
}

// flushWriters 刷新所有日志文件，保证进度数据已写入磁盘文件。
func flushWriters(writers []*logFileWriter) error {
	for _, writer := range writers {
		if err := writer.flush(); err != nil {
			return err
		}
	}
	return nil
}

// allFilesComplete 判断所有目标日志是否均已达到精确容量。
func allFilesComplete(writers []*logFileWriter) bool {
	for _, writer := range writers {
		if writer.size() != writer.spec.target {
			return false
		}
	}
	return true
}

// printGenerateUsage 输出生成日志所需的参数。
func printGenerateUsage(writer io.Writer) {
	fmt.Fprintln(writer, "用法：go run ./practice/file_t generate [参数]")
	fmt.Fprintln(writer, "  -size-gb 日志总大小（GiB，默认 20）")
	fmt.Fprintln(writer, "  -output-dir 日志输出目录（默认 practice/file_t/logs<大小>）")
	fmt.Fprintln(writer, "  -overwrite 允许覆盖已有日志文件")
	fmt.Fprintln(writer, "  -status-interval 进度刷新周期")
	fmt.Fprintln(writer, "  -seed 随机种子")
}
