
from utils import get_subject
import os
import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart



 
# 邮件发送者和接收者
receiver_email = sender_email = os.getenv("EMAIL_PRO", "")

password = os.getenv("EMAIL_PASS", "")  # 注意这里不是你的邮箱密码，而是开启SMTP服务后得到的密码
 
# 创建邮件对象和设置邮件内容
message = MIMEMultipart("alternative")
s = get_subject()
message["Subject"] = s
message["From"] = sender_email
message["To"] = receiver_email
 
# 创建邮件正文
text = """\
This is an example email body.
It can be in HTML or plain text.
"""

# 添加文本和HTML的部分
part1 = MIMEText(text, "plain")

 
# 添加正文到邮件对象中
message.attach(part1)

 
# 发送邮件
try:
    # 创建SMTP服务器连接
    server = smtplib.SMTP_SSL("smtp.163.com", 465)
    server.login(sender_email, password)
    server.sendmail(sender_email, receiver_email, message.as_string())
    server.close()
    print("Email sent successfully!")
except Exception as e:
    print("Something went wrong...", e)



