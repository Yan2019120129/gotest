package traverse

import (
	"bufio"
	"container/heap"
	"context"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type NodeCount struct {
	NodeID string
	Count  int64
}

func TopKNodes(
	ctx context.Context,
	filePath string,
	startTime int64,
	endTime int64,
	k int,
) ([]NodeCount, error) {
	if filePath == "" {
		return nil, nil
	}

	fileIo, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileIo.Close()

	size := 64 * 1024
	scanner := bufio.NewReaderSize(fileIo, size)
	topM := make(map[string]NodeCount)
	startTimeTmp := time.Unix(startTime, 0)
	endTimeTmp := time.Unix(endTime, 0)
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		node, err := Worker(strings.Clone(strings.TrimSpace(line)))
		if err != nil {
			continue
		}
		if node.Time.After(startTimeTmp) && node.Time.Before(endTimeTmp) {
			v := topM[node.Name]
			v.NodeID = node.Name
			v.Count++
			topM[node.Name] = v
		}
	}

	return selectTopK(topM, k), nil
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

type Node struct {
	Time    time.Time
	Name    string
	Status  int
	RspTime string
}

// Worker 工作模式处理字节
// 2026-07-29T12:32:56.147020891Z INFO [access] access_time=2026-07-29T12:32:56.147020891Z node_id=api-node-24 request_id=de9099b31cb250f6 remote_addr=10.10.164.204 method=PUT path="/api/v1/login" status=204 bytes=94940 duration_ms=387 user_agent="Mozilla/5.0"
func Worker(line string) (Node, error) {
	//time.Sleep(1 * time.Microsecond)
	if line == "" {
		return Node{}, errors.New("nil")
	}

	fields := strings.SplitAfter(line, " ")
	if len(fields) < 11 {
		return Node{}, errors.New("not enough fields")
	}

	_, timeStr, found := strings.Cut(fields[3], "=")
	timeStr = strings.TrimSpace(timeStr)
	if !found || timeStr == "" {
		return Node{}, errors.New("not found field time")
	}

	nodeTime, err := time.Parse(time.RFC3339Nano, strings.TrimSpace(timeStr))
	if err != nil {
		return Node{}, err
	}

	_, name, found := strings.Cut(fields[4], "=")
	name = strings.TrimSpace(name)
	if !found || name == "" {
		return Node{}, errors.New("not found field node")
	}

	_, statusStr, found := strings.Cut(fields[9], "=")
	statusStr = strings.TrimSpace(statusStr)
	if !found || statusStr == "" {
		return Node{}, errors.New("not found field node")
	}

	status, err := strconv.ParseInt(statusStr, 10, 64)
	if err != nil {
		return Node{}, err
	}

	_, rspTimeStr, found := strings.Cut(fields[11], "=")
	rspTimeStr = strings.TrimSpace(rspTimeStr)
	if !found || rspTimeStr == "" {
		return Node{}, errors.New("not found field rsp time")
	}

	return Node{
		Time:    nodeTime,
		Name:    name,
		Status:  int(status),
		RspTime: rspTimeStr,
	}, nil
}
