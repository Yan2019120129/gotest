package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gotest/practice/file_t/read_t/custom"
	customoptimized "gotest/practice/file_t/read_t/custom_optimized"
	"gotest/practice/file_t/read_t/traverse"
	"io"
	"os"
	"time"
)

const (
	// defaultMethod 使用五分片流式读取方案作为默认实现。
	defaultMethod = "custom_optimized"
	// defaultLogPath 表示从 gotest 模块根目录执行时的默认访问日志路径。
	defaultLogPath = "practice/file_t/read_t/logs20/access.log"
)

// commandConfig 保存命令行解析后的日志处理参数。
type commandConfig struct {
	method    string
	filePath  string
	startTime time.Time
	endTime   time.Time
	topK      int
}

// nodeCount 统一三种实现返回的节点统计结果。
type nodeCount struct {
	nodeID string
	count  int64
}

// commandRunner 根据命令行选择的实现执行日志统计。
type commandRunner struct {
	config commandConfig
}

// application 负责分派读取与生成日志两个命令。
type application struct {
	arguments   []string
	output      io.Writer
	errorOutput io.Writer
	now         func() time.Time
}

// main 创建应用并将运行错误输出到标准错误流。
func main() {
	app := newApplication(os.Args[1:], os.Stdout, os.Stderr, time.Now)
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// newApplication 创建命令行应用。
func newApplication(arguments []string, output io.Writer, errorOutput io.Writer, now func() time.Time) *application {
	return &application{
		arguments:   arguments,
		output:      output,
		errorOutput: errorOutput,
		now:         now,
	}
}

// Run 根据子命令执行日志读取或日志生成；省略子命令时保持读取兼容行为。
func (a *application) Run() error {
	arguments := a.arguments
	if len(arguments) == 0 {
		return a.runReader(arguments)
	}

	switch arguments[0] {
	case "generate":
		return a.runGenerator(arguments[1:])
	case "read":
		return a.runReader(arguments[1:])
	case "help", "-h", "--help":
		printRootUsage(a.output)
		return nil
	default:
		return a.runReader(arguments)
	}
}

// runReader 解析读取参数并输出指定时间范围内的节点统计结果。
func (a *application) runReader(arguments []string) error {
	config, err := parseCommandConfig(arguments, a.now())
	if errors.Is(err, flag.ErrHelp) {
		printReadUsage(a.output)
		return nil
	}
	if err != nil {
		printReadUsage(a.errorOutput)
		return fmt.Errorf("参数错误: %w", err)
	}

	startedAt := a.now()
	nodes, err := (&commandRunner{config: config}).Run(context.Background())
	if err != nil {
		return fmt.Errorf("处理日志失败: %w", err)
	}

	fmt.Fprintf(a.output, "方法: %s\n", config.method)
	fmt.Fprintf(a.output, "时间范围: %s ~ %s\n", config.startTime.Format(time.RFC3339), config.endTime.Format(time.RFC3339))
	fmt.Fprintf(a.output, "耗时: %s，节点数: %d\n", a.now().Sub(startedAt), len(nodes))
	return nil
}

// parseCommandConfig 将命令行参数转换为经过校验的读取配置。
func parseCommandConfig(arguments []string, now time.Time) (commandConfig, error) {
	defaultStart, defaultEnd := defaultTimeRange(now)
	flags := flag.NewFlagSet("read_t read", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	method := flags.String("method", defaultMethod, "读取方法：custom、custom_optimized 或 traverse")
	filePath := flags.String("file", defaultLogPath, "访问日志文件路径")
	start := flags.String("start", defaultStart.Format(time.RFC3339), "统计起始时间，RFC3339 格式")
	end := flags.String("end", defaultEnd.Format(time.RFC3339), "统计结束时间，RFC3339 格式")
	topK := flags.Int("k", 10, "返回的节点数量")
	if err := flags.Parse(arguments); err != nil {
		return commandConfig{}, err
	}
	if flags.NArg() > 0 {
		return commandConfig{}, fmt.Errorf("不支持的位置参数: %s", flags.Arg(0))
	}

	startTime, err := time.Parse(time.RFC3339, *start)
	if err != nil {
		return commandConfig{}, fmt.Errorf("-start 必须是 RFC3339 时间: %w", err)
	}
	endTime, err := time.Parse(time.RFC3339, *end)
	if err != nil {
		return commandConfig{}, fmt.Errorf("-end 必须是 RFC3339 时间: %w", err)
	}
	if *filePath == "" {
		return commandConfig{}, errors.New("-file 不能为空")
	}
	if !startTime.Before(endTime) {
		return commandConfig{}, errors.New("-start 必须早于 -end")
	}
	if *topK <= 0 {
		return commandConfig{}, errors.New("-k 必须大于 0")
	}

	config := commandConfig{
		method:    *method,
		filePath:  *filePath,
		startTime: startTime,
		endTime:   endTime,
		topK:      *topK,
	}
	if !config.isSupportedMethod() {
		return commandConfig{}, fmt.Errorf("不支持的方法 %q，可选值为 custom、custom_optimized、traverse", config.method)
	}

	return config, nil
}

// defaultTimeRange 保持原示例中当前年月 29 日的默认 UTC 时间范围。
func defaultTimeRange(now time.Time) (time.Time, time.Time) {
	start := time.Date(now.Year(), now.Month(), 29, 13, 32, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month(), 29, 13, 55, 0, 0, time.UTC)
	return start, end
}

// isSupportedMethod 判断命令行指定的方法是否可用。
func (c commandConfig) isSupportedMethod() bool {
	switch c.method {
	case "custom", "custom_optimized", "traverse":
		return true
	default:
		return false
	}
}

// Run 按选择的方法调用对应的 TopKNodes 实现并统一返回结果。
func (r *commandRunner) Run(ctx context.Context) ([]nodeCount, error) {
	config := r.config
	switch config.method {
	case "custom":
		nodes, err := custom.TopKNodes(ctx, config.filePath, config.startTime.Unix(), config.endTime.Unix(), config.topK)
		return convertCustomNodes(nodes), err
	case "custom_optimized":
		nodes, err := customoptimized.TopKNodes(ctx, config.filePath, config.startTime.Unix(), config.endTime.Unix(), config.topK)
		return convertOptimizedNodes(nodes), err
	case "traverse":
		nodes, err := traverse.TopKNodes(ctx, config.filePath, config.startTime.Unix(), config.endTime.Unix(), config.topK)
		return convertTraverseNodes(nodes), err
	default:
		return nil, fmt.Errorf("不支持的方法 %q", config.method)
	}
}

// convertCustomNodes 转换 custom 包的统计结果。
func convertCustomNodes(nodes []custom.NodeCount) []nodeCount {
	result := make([]nodeCount, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, nodeCount{nodeID: node.NodeID, count: node.Count})
	}
	return result
}

// convertOptimizedNodes 转换 custom_optimized 包的统计结果。
func convertOptimizedNodes(nodes []customoptimized.NodeCount) []nodeCount {
	result := make([]nodeCount, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, nodeCount{nodeID: node.NodeID, count: node.Count})
	}
	return result
}

// convertTraverseNodes 转换 traverse 包的统计结果。
func convertTraverseNodes(nodes []traverse.NodeCount) []nodeCount {
	result := make([]nodeCount, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, nodeCount{nodeID: node.NodeID, count: node.Count})
	}
	return result
}

// printRootUsage 输出应用提供的子命令。
func printRootUsage(writer io.Writer) {
	fmt.Fprintln(writer, "用法：go run ./practice/file_t/read_t [read|generate] [参数]")
	fmt.Fprintln(writer, "  read      读取访问日志；省略子命令时默认执行此操作")
	fmt.Fprintln(writer, "  generate  生成访问、信息和错误日志")
}

// printReadUsage 输出读取日志所需的参数。
func printReadUsage(writer io.Writer) {
	fmt.Fprintln(writer, "用法：go run ./practice/file_t/read_t read [参数]")
	fmt.Fprintln(writer, "  -method custom|custom_optimized|traverse")
	fmt.Fprintln(writer, "  -file 日志文件路径")
	fmt.Fprintln(writer, "  -start RFC3339 起始时间")
	fmt.Fprintln(writer, "  -end RFC3339 结束时间")
	fmt.Fprintln(writer, "  -k Top-K 节点数量")
}
