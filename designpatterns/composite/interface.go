package composite

type DirComposite interface {
	// Show 显示文件夹和文件
	Show()
	// Search 查找文件
	Search(key string)

	Remove(key string) bool
}
