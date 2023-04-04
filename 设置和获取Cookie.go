package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

//PS D:\GolandProjects\gintest> curl http://localhost:8080/cookie
//
//
//StatusCode        : 200
//StatusDescription : OK
//Content           : {}
//RawContent        : HTTP/1.1 200 OK
//Content-Length: 0
//Date: Tue, 04 Apr 2023 15:23:09 GMT
//Set-Cookie: gin_cookie=test; Path=/; Domain=localhost; Max-Age=3600; HttpOnly
//
//
//Headers           : {[Content-Length, 0], [Date, Tue, 04 Apr 2023 15:23:09 GMT], [Set-Cookie, gin_cookie=test; Path=/; Domain=localhost; Max-Age=3600;
//HttpOnly]}
//RawContentLength  : 0
