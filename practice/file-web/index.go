package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func main() {
	// 默认端口
	port := "8080"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}

	// 创建上传目录
	os.MkdirAll(uploadDir, os.ModePerm)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", filesHandler)

	// 文件下载
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(uploadDir))))

	fmt.Println("Server started at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

// 首页（上传 + 文件列表入口）
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<html>
	<head>
		<title>文件服务</title>
	</head>
	<body>
		<h2>上传文件</h2>
		<form enctype="multipart/form-data" action="/upload" method="post">
			<input type="file" name="file"/>
			<input type="submit" value="上传"/>
		</form>

		<h2>文件列表</h2>
		<a href="/files">查看文件</a>
	</body>
	</html>
	`
	t, _ := template.New("index").Parse(tpl)
	t.Execute(w, nil)
}

// 上传处理
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "上传失败", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, "保存失败", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}

// 文件列表
func filesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "读取目录失败", http.StatusInternalServerError)
		return
	}

	tpl := `
	<html>
	<head>
		<title>文件列表</title>
	</head>
	<body>
		<h2>文件列表</h2>
		<ul>
		{{range .}}
			<li>
				<a href="/download/{{.Name}}">{{.Name}}</a>
			</li>
		{{end}}
		</ul>
		<a href="/">返回首页</a>
	</body>
	</html>
	`

	t, _ := template.New("files").Parse(tpl)
	t.Execute(w, files)
}
