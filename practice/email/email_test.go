package email

import "testing"

func TestSendQQMail(t *testing.T) {

	client := NewClient(
		"smtp.qq.com",
		587,
		"1556403682@qq.com",
		"dlogwcjjuhfzhdce",
		"价格监控机器人",
	)

	err := client.Send(
		[]string{
			"1556403682@qq.com",
		},
		"Go单元测试邮件",
		"测试成功",
	)

	if err != nil {
		t.Fatal(err)
	}
}
