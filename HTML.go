package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*使用 LoadHTMLGlob() 或者 LoadHTMLFiles()*/
//使用上边两个方法加载html模板 美其名曰渲染
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
