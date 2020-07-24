package main

import "github.com/gin-gonic/gin"

/*

  ___ ___ .____       _________________    ________
 /   |   \|    |     /   _____/\_____  \  /  _____/
/    ~    \    |     \_____  \  /   |   \/   __  \
\    Y    /    |___  /        \/    |    \  |__\  \
 \___|_  /|_______ \/_______  /\_______  /\_____  /
       \/         \/        \/         \/       \/

*/

/*
	conf：用于存储配置文件
	middleware：应用中间件
	models：应用数据库模型
	pkg：第三方包
	routers 路由逻辑处理
	runtime：应用运行时数据
*/

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
