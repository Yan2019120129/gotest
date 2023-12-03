package composite

import (
	"fmt"
)

type Folder struct {
	DirComposite []DirComposite
	File         *File
}

// Search 查找文件
func (fd *Folder) Search(key string) {
	for _, Dir := range fd.DirComposite {
		Dir.Search(key)
	}
}

// Show 显示文件
func (fd *Folder) Show() {
	fmt.Println("Dir-detail---------------->:", fd.File)
	for _, dirComposite := range fd.DirComposite {
		dirComposite.Show()
	}
}

// Add 添加文件
func (fd *Folder) Add(f DirComposite) {
	fd.DirComposite = append(fd.DirComposite, f)
}

// Remove 删除文件
func (fd *Folder) Remove(key string) bool {
	for i, composite := range fd.DirComposite {
		if composite.Remove(key) {
			fd.DirComposite = append(fd.DirComposite[:i], fd.DirComposite[i+1:]...)
		}
	}
	return false
}
