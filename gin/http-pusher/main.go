package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

// 1、生成RSA密钥的方法

// openssl genrsa -des3 -out privkey.pem 2048
// 这个命令会生成一个2048位的密钥，同时有一个des3方法加密的密码，如果你不想要每次都输入密码，可以改成：
// openssl genrsa -out privkey.pem 2048
// 建议用2048位密钥，少于此可能会不安全或很快将不安全。

// 2、生成一个证书请求
// openssl req -new -key privkey.pem -out cert.csr
// 这个命令将会生成一个证书请求，当然，用到了前面生成的密钥privkey.pem文件
// 这里将生成一个新的文件cert.csr，即一个证书请求文件，你可以拿着这个文件去数字证书颁发机构（即CA）申请一个数字证书。CA会给你一个新的文件cacert.pem，那才是你的数字证书。

// 如果是自己做测试，那么证书的申请机构和颁发机构都是自己。就可以用下面这个命令来生成证书：
// openssl req -new -x509 -key privkey.pem -out cacert.pem -days 1095
// 这个命令将用上面生成的密钥privkey.pem生成一个数字证书cacert.pem

// 3、使用数字证书和密钥
// 有了privkey.pem和cacert.pem文件后就可以在自己的程序中使用了，比如做一个加密通讯的服务器

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
