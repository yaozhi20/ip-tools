package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	id "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

// 登录函数
func loginEmail(Eserver, UserName, Password string) (*client.Client, error) {
	c, err := client.DialTLS(Eserver, nil)
	if err != nil {
		return nil, err
	}
	//登陆
	if err = c.Login(UserName, Password); err != nil {
		return nil, err
	}
	return c, nil
}

// 邮件接收

func emailList(Eserver, UserName, Password string) string {
	c, err := loginEmail(Eserver, UserName, Password)
	if err != nil {
		fmt.Println("login err")
		fmt.Println(err)
		return ""
	}
	idClient := id.NewClient(c)
	idClient.ID(
		id.ID{
			id.FieldName:    "IMAPClient",
			id.FieldVersion: "3.1.0",
		},
	)

	defer c.Close()
	// 选择收件箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {

		fmt.Println("select inbox err: ", err)
		return ""
	}
	if mbox.Messages == 0 {
		return ""
	}
	// 选择收取邮件的时间段
	criteria := imap.NewSearchCriteria()
	// // 收取7天之内的邮件
	// t1, err := time.Parse("2006-01-02 15:04:05", "2020-03-02 15:04:05")
	// criteria.Since = t1
	// // 按条件查询邮件
	// ids, err := c.Search(criteria)
	// 收取7天之内的邮件
	criteria.Since = time.Now().Add(-3 * time.Hour * 24)
	// 按条件查询邮件
	// ids, err := c.UidSearch(criteria)
	ids, err := c.Search(criteria)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("ids: ", ids)
	// fmt.Println("mbox.Messages: ", mbox.Messages)
	if len(ids) == 0 {
		return ""
	}
	seqset := new(imap.SeqSet)
	seqset.AddNum(ids...)
	sect := &imap.BodySectionName{}
	messages := make(chan *imap.Message, 100)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{sect.FetchItem()}, messages)
	}()

	for msg := range messages {
		// for i := 0; i < len(messages); i++ {
		// msg := messages[i]
		r := msg.GetBody(sect)
		// m, err := message.Read(r)
		// if err != nil {
		// 	fmt.Println(err)
		// 	// return err
		// }

		m, err := mail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
		}

		header := m.Header

		var subject string
		if date, err := header.Date(); err == nil {
			log.Println("Date:", date)
		}
		// if from, err := header.AddressList("From"); err == nil {
		// log.Println("From:", from)
		// }
		// if to, err := header.AddressList("To"); err == nil {
		// log.Println("To:", to)
		// }
		if subject, err = header.Subject(); err == nil {
			log.Println("Subject:", subject)
		}

		if strings.HasPrefix(subject, "ip:") {
			log.Println("The string starts with 'Hello'.")
			now := today()
			fmt.Println("today:", now)
			slice := strings.Split(subject, "/")
			if now == slice[1] {
				ip := get_dynamic_local_ip(subject)
				return ip
			} else {
				continue
			}

		}

		// // 处理邮件正文
		// _, fileName := parseEmail(m)
		// for k, _ := range fileName {
		// 	fmt.Println("收取到附件:", k)
		// }

		// 处理邮件正文
		// for {
		// 	p, err := m.NextPart()
		// 	if err == io.EOF {
		// 		break
		// 	} else if err != nil {
		// 		log.Fatal("NextPart:err ", err)
		// 	}

		// 	switch h := p.Header.(type) {
		// 	case *mail.InlineHeader:
		// 		// 正文消息文本
		// 		b, _ := ioutil.ReadAll(p.Body)
		// 		mailFile := fmt.Sprintf("INBOX/%s.eml", subject)
		// 		f, _ := os.OpenFile(mailFile, os.O_RDWR|os.O_CREATE, 0766)
		// 		f.Write(b)
		// 		f.Close()
		// 	case *mail.AttachmentHeader:
		// 		// 正文内附件
		// 		filename, _ := h.Filename()
		// 		log.Printf("attachment: %v\n", filename)
		// 	}
		// }

		// emailDate, _ := net_mail.ParseDate(header.Get("Date"))
		// subject := tools.GetSubject(header)
		// from := tools.GetFrom(header)
		// 读取邮件内容
		// body, _ := tools.ParseBody(m.Body)
		// fmt.Printf("%s 在时间为:%v 发送了主题为:%s \n", from, emailDate, subject)
		// fmt.Printf("%s 在时间为:%v 发送了主题为: \n", header, emailDate)

	}
	return ""
}

func parseEmail(mr *mail.Reader) (body []byte, fileMap map[string][]byte) {
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		if p != nil {
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				body, err = ioutil.ReadAll(p.Body)
				if err != nil {
					fmt.Println("read body err:", err.Error())
				}
				// log.Println("body:", body)
				mailFile := fmt.Sprintf("1.eml")
				f, _ := os.OpenFile(mailFile, os.O_RDWR|os.O_CREATE, 0766)
				f.Write(body)
				f.Close()

			case *mail.AttachmentHeader:
				fileName, _ := h.Filename()
				fileContent, _ := ioutil.ReadAll(p.Body)
				fileMap[fileName] = fileContent
			}
		}
	}
	return
}

func today() string {
	now := time.Now() // 默认是local
	nowday := fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())

	return nowday
}

func get_dynamic_local_ip(subject string) string {
	slice := strings.Split(subject, "/")
	slice2 := strings.Split(slice[0], ":")

	return slice2[1]
}

func latest_ip() string {
	ip := emailList("imap.163.com:993", os.Getenv("EMAIL_PRO"), os.Getenv("EMAIL_PASS"))
	if ip != "" {
		log.Println("ip:", ip)
		return ip
	} else {
		return ""
	}
}

func update_hosts() {
	ip := latest_ip()
	hostsItems(ip)
}

func hostMaster() {
	update_hosts()
	hosts_backup()
	hosts_delete()
	hosts_copy()

}

func main() {
	latest_ip()
}
