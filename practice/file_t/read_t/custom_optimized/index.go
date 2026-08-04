package custom_optimized

import (
	"bufio"
	"bytes"
	"container/heap"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	// partitionCount 表示固定的文件逻辑分片数量。
	partitionCount = 5
	// readBufferSize 表示边界定位和顺序读取时使用的缓冲区大小。
	readBufferSize = 64 * 1024
	// maxLineSize 限制单条日志最大长度，避免异常日志行占用无限内存。
	maxLineSize = 1024 * 1024
)

// NodeCount 表示一个节点在指定时间范围内的日志数量。
type NodeCount struct {
	NodeID string
	Count  int64
}

// Node 表示从单条访问日志中提取的统计字段。
type Node struct {
	Time time.Time
	Name string
}

// filePartition 表示一个已对齐到完整日志行的文件字节区间。
type filePartition struct {
	start int64
	end   int64
}

// logProcessor 负责协调分片读取、并发解析和统计汇总。
type logProcessor struct {
	workerCount int
	bufferSize  int
	maxLineSize int
}

// newLogProcessor 创建使用当前机器可用 CPU 数量的日志处理器。
func newLogProcessor() *logProcessor {
	workerCount := runtime.GOMAXPROCS(0)
	if workerCount < 1 {
		workerCount = 1
	}

	return &logProcessor{
		workerCount: workerCount,
		bufferSize:  readBufferSize,
		maxLineSize: maxLineSize,
	}
}

// TopKNodes 统计指定时间范围内出现次数最多的 k 个节点。
func TopKNodes(
	ctx context.Context,
	filePath string,
	startTime int64,
	endTime int64,
	k int,
) ([]NodeCount, error) {
	if ctx == nil {
		return nil, errors.New("上下文不能为空")
	}
	if filePath == "" {
		return nil, errors.New("日志文件路径不能为空")
	}
	if k <= 0 {
		return []NodeCount{}, nil
	}

	return newLogProcessor().topKNodes(ctx, filePath, startTime, endTime, k)
}

// topKNodes 执行分片读取、并发解析、节点统计和 Top-K 排序。
func (p *logProcessor) topKNodes(
	ctx context.Context,
	filePath string,
	startTime int64,
	endTime int64,
	k int,
) ([]NodeCount, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	if fileInfo.Size() == 0 {
		return []NodeCount{}, nil
	}

	partitions, err := p.createPartitions(filePath, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	processContext, cancel := context.WithCancel(ctx)
	defer cancel()

	lines := make(chan string, p.workerCount*2)
	errorsC := make(chan error, 1)
	var reportOnce sync.Once
	reportError := func(err error) {
		if err == nil {
			return
		}
		reportOnce.Do(func() {
			errorsC <- err
			cancel()
		})
	}

	var readerGroup sync.WaitGroup
	for _, partition := range partitions {
		if partition.start >= partition.end {
			continue
		}
		readerGroup.Add(1)
		go func(partition filePartition) {
			defer readerGroup.Done()
			if err := p.streamPartition(processContext, filePath, partition, lines); err != nil && !errors.Is(err, context.Canceled) {
				reportError(err)
			}
		}(partition)
	}
	go func() {
		readerGroup.Wait()
		close(lines)
	}()

	results := p.process(processContext, lines)
	start := time.Unix(startTime, 0)
	end := time.Unix(endTime, 0)
	counts := make(map[string]NodeCount)
	for node := range results {
		if node.Time.After(start) && node.Time.Before(end) {
			value := counts[node.Name]
			value.NodeID = node.Name
			value.Count++
			counts[node.Name] = value
		}
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}
	select {
	case err := <-errorsC:
		return nil, err
	default:
	}

	return selectTopK(counts, k), nil
}

// createPartitions 将文件按字节均分，并把每个边界校准到完整日志行。
func (p *logProcessor) createPartitions(filePath string, fileSize int64) ([]filePartition, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	partitions := make([]filePartition, partitionCount)
	for index := 0; index < partitionCount; index++ {
		rawStart := fileSize * int64(index) / partitionCount
		rawEnd := fileSize * int64(index+1) / partitionCount

		start := rawStart
		if index > 0 {
			start, err = p.alignStart(file, rawStart, fileSize)
			if err != nil {
				return nil, err
			}
		}

		end := rawEnd
		if index < partitionCount-1 {
			end, err = p.alignEnd(file, rawEnd, fileSize)
			if err != nil {
				return nil, err
			}
		}

		partitions[index] = filePartition{start: start, end: end}
	}

	return partitions, nil
}

// alignStart 将分片起点移到当前位置所在日志行的下一行起始位置。
func (p *logProcessor) alignStart(file *os.File, offset int64, fileSize int64) (int64, error) {
	if offset <= 0 || offset >= fileSize {
		return offset, nil
	}

	previous, err := p.readByteAt(file, offset-1)
	if err != nil {
		return 0, err
	}
	if previous == '\n' {
		return offset, nil
	}

	return p.findNextLineEnd(file, offset, fileSize)
}

// alignEnd 将分片终点扩展到当前位置所在日志行的结束位置。
func (p *logProcessor) alignEnd(file *os.File, offset int64, fileSize int64) (int64, error) {
	if offset <= 0 || offset >= fileSize {
		return offset, nil
	}

	previous, err := p.readByteAt(file, offset-1)
	if err != nil {
		return 0, err
	}
	if previous == '\n' {
		return offset, nil
	}

	return p.findNextLineEnd(file, offset, fileSize)
}

// findNextLineEnd 从指定偏移量向后查找换行符后的字节偏移量。
func (p *logProcessor) findNextLineEnd(file *os.File, offset int64, fileSize int64) (int64, error) {
	buffer := make([]byte, p.bufferSize)
	for offset < fileSize {
		n, err := file.ReadAt(buffer, offset)
		if n > 0 {
			if index := bytes.IndexByte(buffer[:n], '\n'); index >= 0 {
				return offset + int64(index+1), nil
			}
			offset += int64(n)
		}
		if errors.Is(err, io.EOF) {
			return fileSize, nil
		}
		if err != nil {
			return 0, err
		}
	}

	return fileSize, nil
}

// readByteAt 读取指定偏移量的一个字节，用于判断分片边界是否已位于换行处。
func (p *logProcessor) readByteAt(file *os.File, offset int64) (byte, error) {
	buffer := []byte{0}
	_, err := file.ReadAt(buffer, offset)
	if err != nil {
		return 0, err
	}
	return buffer[0], nil
}

// streamPartition 使用独立文件句柄流式读取一个已对齐分片中的完整日志行。
func (p *logProcessor) streamPartition(
	ctx context.Context,
	filePath string,
	partition filePartition,
	lines chan<- string,
) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Seek(partition.start, io.SeekStart); err != nil {
		return err
	}
	limitedReader := &io.LimitedReader{R: file, N: partition.end - partition.start}
	scanner := bufio.NewScanner(limitedReader)
	scanner.Buffer(make([]byte, p.bufferSize), p.maxLineSize)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case lines <- scanner.Text():
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("分片 [%d, %d) 读取失败: %w", partition.start, partition.end, err)
	}

	return nil
}

// process 启动固定数量的 worker，并把合法节点数据汇集到结果通道。
func (p *logProcessor) process(ctx context.Context, lines <-chan string) <-chan Node {
	results := make(chan Node, p.workerCount*2)
	var workerGroup sync.WaitGroup
	for index := 0; index < p.workerCount; index++ {
		workerGroup.Add(1)
		go func() {
			defer workerGroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case line, ok := <-lines:
					if !ok {
						return
					}

					node, err := Worker(line)
					if err != nil {
						continue
					}
					select {
					case <-ctx.Done():
						return
					case results <- node:
					}
				}
			}
		}()
	}

	go func() {
		workerGroup.Wait()
		close(results)
	}()

	return results
}

// nodeMinHeap 保存当前 Top-K 中最差的节点，堆顶始终最容易被替换。
type nodeMinHeap []NodeCount

// Len 返回堆中节点数量。
func (h nodeMinHeap) Len() int {
	return len(h)
}

// Less 定义小根堆顺序：数量更小或同数量下节点名更大者优先位于堆顶。
func (h nodeMinHeap) Less(left, right int) bool {
	return isWorse(h[left], h[right])
}

// Swap 交换堆中的两个节点。
func (h nodeMinHeap) Swap(left, right int) {
	h[left], h[right] = h[right], h[left]
}

// Push 向堆中加入一个节点。
func (h *nodeMinHeap) Push(value any) {
	*h = append(*h, value.(NodeCount))
}

// Pop 删除并返回堆尾节点，供 container/heap 维护堆结构使用。
func (h *nodeMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// isWorse 判断 left 是否比 right 更不应保留在 Top-K 中。
func isWorse(left, right NodeCount) bool {
	if left.Count != right.Count {
		return left.Count < right.Count
	}
	return left.NodeID > right.NodeID
}

// isBetter 判断 left 是否比 right 更应保留在 Top-K 中。
func isBetter(left, right NodeCount) bool {
	if left.Count != right.Count {
		return left.Count > right.Count
	}
	return left.NodeID < right.NodeID
}

// selectTopK 使用容量为 k 的小根堆筛选节点，避免对全部节点进行排序。
func selectTopK(counts map[string]NodeCount, k int) []NodeCount {
	if k <= 0 || len(counts) == 0 {
		return []NodeCount{}
	}

	selected := make(nodeMinHeap, 0, min(k, len(counts)))
	heap.Init(&selected)
	for _, node := range counts {
		if selected.Len() < k {
			heap.Push(&selected, node)
			continue
		}
		if isBetter(node, selected[0]) {
			selected[0] = node
			heap.Fix(&selected, 0)
		}
	}

	result := make([]NodeCount, len(selected))
	copy(result, selected)
	sort.Slice(result, func(left, right int) bool {
		return isBetter(result[left], result[right])
	})
	return result
}

// Worker 从访问日志中提取访问时间和节点标识。
func Worker(line string) (Node, error) {
	//time.Sleep(1 * time.Microsecond)
	if line == "" {
		return Node{}, errors.New("日志行为空")
	}

	_, fields, found := strings.Cut(line, "access_time=")
	if !found {
		return Node{}, errors.New("未找到访问时间字段")
	}
	timeString, fields, found := strings.Cut(fields, " ")
	if !found || timeString == "" {
		return Node{}, errors.New("访问时间字段无效")
	}
	nodeTime, err := time.Parse(time.RFC3339Nano, timeString)
	if err != nil {
		return Node{}, err
	}

	_, fields, found = strings.Cut(fields, "node_id=")
	if !found {
		return Node{}, errors.New("未找到节点字段")
	}
	nodeName, _, _ := strings.Cut(fields, " ")
	if nodeName == "" {
		return Node{}, errors.New("节点字段无效")
	}

	return Node{Time: nodeTime, Name: nodeName}, nil
}
