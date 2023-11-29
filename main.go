package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"html/template"
	"net/http"
)

func main() {
	// 设置模板
	tmpl, err := template.New("index").Parse(`
	<html>
		<body>
			<form action="/submit" method="post">
				<label>验证码: </label>
				<img src="/captcha" alt="captcha"/>
				<input type="text" name="captcha"/>
				<input type="submit" value="提交"/>
			</form>
		</body>
	</html>
	`)
	if err != nil {
		panic(err)
	}

	// 设置验证码服务
	http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		image := captcha.New()
		fmt.Println("image:", image)
		captcha.WriteImage(w, image, 200, 50)
	})

	// 处理表单提交
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userCaptcha := r.FormValue("captcha")

		// 验证用户输入的验证码
		if captcha.VerifyString(userCaptcha, userCaptcha) {
			fmt.Fprintf(w, "验证码正确！")
		} else {
			fmt.Fprintf(w, "验证码错误！")
		}
	})

	// 设置静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 设置首页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// 启动服务器
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
