## s-mail
![](https://img.shields.io/badge/s-mail-brightgreen)
![](https://img.shields.io/badge/golang-9cf)
![](https://img.shields.io/badge/gomail.v2-2.0.0-blue)

基于`gomail.v2`开发的简易小程序，支持给多人发邮件，支持附件

使用前设置

```
const (
	//SenderAddr sender addr
	SenderAddr string = "xx@qq.com"

	//SederPassword sender password
	SederPassword string = "xxxxxx"

	//SMTPServer smtp server
	SMTPServer string = "smtp.qq.com"
	//SMTPPort smtp port
	SMTPPort int = 465
)
```

支持的参数

```
-c: 收邮件地址
-s：邮件主题
-m: 邮件内容
-f: 附件文件名
-t: 发邮件方式，群发或1to1
```

示例：

```
s-mail.exe -c serialt@qq.com,serialt@126.com -s qq -m first -f s-mail.exe                           
```

