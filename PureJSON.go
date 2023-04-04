package main

/*通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。如果要按字面对这些字符进行编码，则可以使用 PureJSON。*/
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run("localhost:8080")
}

/*测试验证效果*/
//普通json
//{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
//真的是把尖括号转成Unicode了
//pureJSON
//{"html":"<b>Hello, world!</b>"}
