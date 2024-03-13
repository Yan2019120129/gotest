package main

import (
	"fmt"
	"gotest/base/designpatterns_t/composite"
	"time"
)

func main() {
	nowTime := time.Now()
	file := &composite.File{Name: "yuzhou.jpg", Type: "jpg", Size: 100, Date: int(nowTime.Unix())}
	file1 := &composite.File{Name: "diqui.jpg", Type: "jpg", Size: 10, Date: int(nowTime.Unix())}
	file2 := &composite.File{Name: "yuqui.png", Type: "png", Size: 11, Date: int(nowTime.Unix())}

	dir := &composite.Folder{File: &composite.File{Name: "picture", Type: "file", Size: 300, Date: int(nowTime.Unix())}}
	dir.Add(file)
	dir.Add(file1)
	dir.Add(file2)

	file3 := &composite.File{Name: "c++.pdf", Type: "pdf", Size: 65, Date: int(nowTime.Unix())}
	file4 := &composite.File{Name: "go.pdf", Type: "pdf", Size: 33, Date: int(nowTime.Unix())}
	file5 := &composite.File{Name: "java.pdf", Type: "pdf", Size: 55, Date: int(nowTime.Unix())}

	dir1 := &composite.Folder{File: &composite.File{Name: "learn", Type: "file", Size: 300, Date: int(nowTime.Unix())}}
	dir1.Add(file3)
	dir1.Add(file4)
	dir1.Add(file5)

	dir2 := composite.Folder{File: &composite.File{Name: "documnet", Type: "file", Size: 500, Date: 600}}

	dir2.Add(dir)
	dir2.Add(dir1)

	//dir2.Search("diqui.jpg")
	dir2.Show()
	fmt.Println("---------------------")
	dir2.Remove("java.pdf")
	fmt.Println("---------------------")
	dir2.Show()

}
