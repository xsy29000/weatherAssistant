package inform

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"log"
)

type Mail struct {
}

func (mail Mail)Inform(data string,receiver string) bool{
	fmt.Println("dd ",data,"ret ",receiver)

	m := gomail.NewMessage()

	m.SetAddressHeader("From", "337612001@qq.com" /*"发件人地址"*/, "天气之子") // 发件人

	m.SetHeader("To", m.FormatAddress(receiver, "User")) // 收件人

	m.SetHeader("Subject", "生活指数")     // 主题

	m.SetBody("text/html", data)


	d := gomail.NewDialer("smtp.qq.com", 465, "337612001@qq.com", "") // 发送邮件服务器、端口、发件人账号、发件人密码

	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return false
	}

	log.Println("done.发送成功")


	return true
}
