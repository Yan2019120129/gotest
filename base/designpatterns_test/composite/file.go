package composite

import (
	"fmt"
	"strings"
)

type File struct {
	Name string // 文件名

	Type string // 文件类型

	Size int // 文件大小

	Date int // 创建日期
}

// Search 查找文件
func (f *File) Search(key string) {
	if strings.Contains(f.Name, key) {
		fmt.Println("file-detail-->：", f)
	}
}

// Show 显示全部文件
func (f *File) Show() {
	fmt.Println("file-detail->：", f)
}

func (f *File) Remove(key string) bool {
	if strings.Contains(f.Name, key) {
		return true
	}
	return false
}
