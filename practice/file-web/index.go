package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const uploadDir = "./uploads"
const messageFile = "./data/messages.json"

type Message struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var (
	messages []Message
	msgMutex sync.RWMutex
)

func main() {

	port := "8080"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}

	os.MkdirAll(uploadDir, os.ModePerm)
	os.MkdirAll("./data", os.ModePerm)

	loadMessages()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", filesHandler)

	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/chat/send", sendMessageHandler)

	http.HandleFunc("/chat/edit", editMessagePage)
	http.HandleFunc("/chat/update", updateMessageHandler)

	http.HandleFunc("/chat/delete", deleteMessageHandler)

	http.Handle(
		"/download/",
		http.StripPrefix(
			"/download/",
			http.FileServer(http.Dir(uploadDir)),
		),
	)

	fmt.Println("Server started at http://localhost:" + port)

	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}

// 首页（上传 + 文件列表入口）
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<html>
	<head>
		<title>文件服务</title>
	</head>
	<body>
		<h2>聊天室</h2>
		<a href="/chat">进入聊天室</a>
		<hr>
		<h2>上传文件</h2>
		<form enctype="multipart/form-data" action="/upload" method="post">
			<input type="file" name="file"/>
			<input type="submit" value="上传"/>
		</form>
		<hr>
		<h2>文件列表</h2>
		<a href="/files">查看文件</a>
	</body>
	</html>
	`
	t, _ := template.New("index").Parse(tpl)
	t.Execute(w, nil)
}

func loadMessages() {

	data, err := os.ReadFile(messageFile)

	if err != nil {

		if os.IsNotExist(err) {
			return
		}

		log.Println(err)
		return
	}

	msgMutex.Lock()
	defer msgMutex.Unlock()

	json.Unmarshal(data, &messages)
}
func editMessagePage(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.ParseInt(
		idStr,
		10,
		64,
	)

	msgMutex.RLock()
	defer msgMutex.RUnlock()

	var target *Message

	for i := range messages {

		if messages[i].ID == id {
			target = &messages[i]
			break
		}
	}

	if target == nil {
		http.NotFound(w, r)
		return
	}

	tpl := `
<html>
<head>
	<title>编辑消息</title>
</head>
<body>

<h2>编辑消息</h2>

<form action="/chat/update" method="post">

	<input
		type="hidden"
		name="id"
		value="{{.ID}}">

	<input
		type="text"
		name="content"
		value="{{.Content}}"
		style="width:400px">

	<input
		type="submit"
		value="保存">

</form>

<br>

<a href="/chat">返回聊天室</a>

</body>
</html>
`

	t := template.Must(
		template.New("edit").Parse(tpl),
	)

	t.Execute(w, target)
}

func deleteMessageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.ParseInt(
		idStr,
		10,
		64,
	)

	msgMutex.Lock()

	for i, msg := range messages {

		if msg.ID == id {

			messages = append(
				messages[:i],
				messages[i+1:]...,
			)

			break
		}
	}

	msgMutex.Unlock()

	saveMessages()

	http.Redirect(
		w,
		r,
		"/chat",
		http.StatusSeeOther,
	)
}

func updateMessageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.ParseInt(
		r.FormValue("id"),
		10,
		64,
	)

	content := r.FormValue("content")

	msgMutex.Lock()

	for i := range messages {

		if messages[i].ID == id {

			messages[i].Content = content

			messages[i].UpdatedAt =
				time.Now().Format(
					"2006-01-02 15:04",
				)

			break
		}
	}

	msgMutex.Unlock()

	saveMessages()

	http.Redirect(
		w,
		r,
		"/chat",
		http.StatusSeeOther,
	)
}

func saveMessages() error {

	msgMutex.RLock()

	data, err := json.MarshalIndent(
		messages,
		"",
		"    ",
	)

	msgMutex.RUnlock()

	if err != nil {
		return err
	}

	return os.WriteFile(
		messageFile,
		data,
		0644,
	)
}

func addMessage(content string) {

	msg := Message{
		ID:        time.Now().UnixNano(),
		Content:   content,
		CreatedAt: time.Now().Format("2006-01-02 15:04"),
	}

	messages = append(messages, msg)

	saveMessages()
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `
<html>
<head>
<title>聊天室</title>
</head>
<body>

<h2>聊天室</h2>

{{range .}}

<div style="
	border:1px solid #ccc;
	padding:10px;
	margin-bottom:10px;
">

	<div>
		<b>发送时间：</b>{{.CreatedAt}}
	</div>

	{{if .UpdatedAt}}
	<div>
		<b>修改时间：</b>{{.UpdatedAt}}
	</div>
	{{end}}

	<div style="margin-top:8px;">
		{{.Content}}
	</div>

	<div style="margin-top:8px;">

		<a href="/chat/edit?id={{.ID}}">
			编辑
		</a>

		&nbsp;

		<a href="/chat/delete?id={{.ID}}">
			删除
		</a>

	</div>

</div>

{{end}}

<hr>

<form action="/chat/send" method="post">

	<textarea
		name="content"
		rows="10"
		cols="100"
		placeholder="请输入消息"></textarea>

	<br><br>

	<input
		type="submit"
		value="发送">

</form>

<br>

<a href="/">返回首页</a>

</body>
</html>
`

	t := template.Must(
		template.New("chat").Parse(tpl),
	)

	t.Execute(w, messages)
}

func sendMessageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Redirect(
			w,
			r,
			"/chat",
			http.StatusSeeOther,
		)
		return
	}

	content := r.FormValue("content")

	if content != "" {

		msgMutex.Lock()

		messages = append(
			messages,
			Message{
				ID:        time.Now().UnixNano(),
				Content:   content,
				CreatedAt: time.Now().Format("2006-01-02 15:04"),
			},
		)

		msgMutex.Unlock()

		saveMessages()
	}

	http.Redirect(
		w,
		r,
		"/chat",
		http.StatusSeeOther,
	)
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
