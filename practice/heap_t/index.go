package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	heap := NewMinHeap()
	fmt.Println("每次请输入一个整数，输入 exit 或 quit 结束：")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if isExitCommand(input) {
			return
		}

		value, err := parseInteger(input)
		if err != nil {
			fmt.Printf("输入错误：%v\n", err)
			continue
		}

		index := heap.Add(value)
		fmt.Printf("新元素 %d 的数组下标：%d，当前小根堆数组：%v\n", value, index, heap.Values())
		fmt.Printf("小根堆结构：\n%s\n", heap.TreeString())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取输入失败：%v\n", err)
	}
}

// MinHeap 表示使用数组存储的小根堆。
type MinHeap struct {
	values []int
}

// NewMinHeap 创建一个空的小根堆。
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Values 返回小根堆数组数据的副本，避免调用方修改堆内部数据。
func (h *MinHeap) Values() []int {
	return append([]int(nil), h.values...)
}

// Add 将整数添加到小根堆，并返回该新元素最终所在的数组下标。
func (h *MinHeap) Add(value int) int {
	h.values = append(h.values, value)
	return h.siftUp(len(h.values) - 1)
}

// TreeString 将小根堆按层级二叉树的形式转换为字符串。
func (h *MinHeap) TreeString() string {
	if len(h.values) == 0 {
		return "（空堆）"
	}

	height := h.height()
	nodeWidth := h.maxNodeWidth()
	leafCount := 1 << (height - 1)
	slotWidth := nodeWidth + 3
	canvasWidth := leafCount * slotWidth
	lines := make([]string, 0, height*2-1)

	for level := 0; level < height; level++ {
		nodeLine := newCanvasLine(canvasWidth)
		firstIndex := (1 << level) - 1
		nodeCount := min(1<<level, len(h.values)-firstIndex)
		span := 1 << (height - level - 1)

		for position := 0; position < nodeCount; position++ {
			index := firstIndex + position
			value := strconv.Itoa(h.values[index])
			center := (2*position + 1) * span * slotWidth / 2
			nodeLine.write(center-len(value)/2, value)
		}
		lines = append(lines, nodeLine.String())

		if level == height-1 {
			continue
		}

		branchLine := newCanvasLine(canvasWidth)
		for position := 0; position < nodeCount; position++ {
			parentIndex := firstIndex + position
			center := (2*position + 1) * span * slotWidth / 2
			branchOffset := max(1, span*slotWidth/4)
			leftChild := parentIndex*2 + 1
			rightChild := parentIndex*2 + 2

			if leftChild < len(h.values) {
				branchLine.write(center-branchOffset, "/")
			}
			if rightChild < len(h.values) {
				branchLine.write(center+branchOffset, "\\")
			}
		}
		lines = append(lines, branchLine.String())
	}

	return strings.Join(lines, "\n")
}

// height 返回小根堆包含的层数。
func (h *MinHeap) height() int {
	height := 0
	for nodeCount := len(h.values); nodeCount > 0; nodeCount >>= 1 {
		height++
	}
	return height
}

// maxNodeWidth 返回堆中节点整数文本的最大宽度。
func (h *MinHeap) maxNodeWidth() int {
	width := 1
	for _, value := range h.values {
		width = max(width, len(strconv.Itoa(value)))
	}
	return width
}

// canvasLine 表示一行可按指定位置写入字符的文本画布。
type canvasLine struct {
	data []byte
}

// newCanvasLine 创建指定宽度的空白文本画布。
func newCanvasLine(width int) *canvasLine {
	return &canvasLine{data: []byte(strings.Repeat(" ", width))}
}

// write 将文本写入画布的指定位置。
func (line *canvasLine) write(position int, text string) {
	copy(line.data[position:], text)
}

// String 返回去除行尾空格后的画布内容。
func (line *canvasLine) String() string {
	return strings.TrimRight(string(line.data), " ")
}

// siftUp 将 index 位置的新元素上浮到符合小根堆规则的位置。
func (h *MinHeap) siftUp(index int) int {
	for index > 0 {
		parent := (index - 1) / 2
		if h.values[parent] <= h.values[index] {
			break
		}

		h.values[parent], h.values[index] = h.values[index], h.values[parent]
		index = parent
	}
	return index
}

// isExitCommand 判断输入是否为退出命令。
func isExitCommand(input string) bool {
	command := strings.ToLower(strings.TrimSpace(input))
	return command == "exit" || command == "quit"
}

// parseInteger 将输入文本解析为单个整数。
func parseInteger(input string) (int, error) {
	if len(strings.Fields(input)) != 1 {
		return 0, fmt.Errorf("每次只能输入一个整数")
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("%q 不是有效整数", input)
	}
	return value, nil
}
