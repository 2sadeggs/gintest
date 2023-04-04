package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
*/
/*本例子简单展示：通过postform方法从请求urk中简单获取对应key的value 并打印*/
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})
	router.Run("localhost:8080")
}

//测试命令
//curl --form message=xiaoxiha --form name=nicheng  http://localhost:8080/post?id=123&page=1
//验证效果 gin后台输出
/*[GIN] 2023/04/04 - 22:33:52 | 200 |       585.3µs |       127.0.0.1 | POST     "/post?id=1234"
id: 123; page: 0; name: nicheng; message: xiaoxiha
*/
//本例子既有从表单中取值 也有从url参数列表取值
