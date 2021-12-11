module github.com/snowlyg/iris-admin-rbac

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/gookit/color v1.5.0
	github.com/kataras/iris/v12 v12.2.0-alpha4.0.20211013142751-e2f40ca06e5e
	github.com/mojocn/base64Captcha v1.3.5
	github.com/snowlyg/helper v0.0.0-20211211140133-76c34b3ebebd
	github.com/snowlyg/iris-admin v0.0.0-20211210085502-a51dbf1f9a88
	github.com/snowlyg/multi v0.0.0-20211210024630-97a71119f4b8
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
	gorm.io/gorm v1.22.4
)

require gorm.io/driver/postgres v1.2.3 // indirect
