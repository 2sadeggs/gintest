package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run("localhost:8080")
}

//正面例子 也就是正常登录成功的例子
//$ curl foo:bar@localhost:8080/admin/secrets -v
//*   Trying 127.0.0.1:8080...
//% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
//Dload  Upload   Total   Spent    Left  Speed
//0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (127.0.0.1) port 8080 (#0)
//* Server auth using Basic with user 'foo'
//> GET /admin/secrets HTTP/1.1
//> Host: localhost:8080
//> Authorization: Basic Zm9vOmJhcg==
//> User-Agent: curl/7.79.1
//> Accept: */*
//>
//* Mark bundle as not supporting multiuse
//< HTTP/1.1 200 OK    //这里相应的200 也这名验证通过了
//< Content-Type: application/json; charset=utf-8
//< Date: Tue, 04 Apr 2023 14:57:13 GMT
//< Content-Length: 64
//<
//{ [64 bytes data]
//100    64  100    64    0     0   9588      0 --:--:-- --:--:-- --:--:-- 64000{
//换行重点说明
//"secret":{"email":"foo@bar.com","phone":"123433"},"user":"foo"}
//* Connection #0 to host localhost left intact

//反面验证不通过的
//$ curl admin:123456@localhost:8080/admin/secrets -v
//*   Trying 127.0.0.1:8080...
//% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
//Dload  Upload   Total   Spent    Left  Speed
//0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (127.0.0.1) port 8080 (#0)
//* Server auth using Basic with user 'admin'
//> GET /admin/secrets HTTP/1.1
//> Host: localhost:8080
//> Authorization: Basic YWRtaW46MTIzNDU2
//> User-Agent: curl/7.79.1
//> Accept: */*
//>
//* Mark bundle as not supporting multiuse
//< HTTP/1.1 401 Unauthorized       //注意这里是401 未认证
//* Authentication problem. Ignoring this.
//< Www-Authenticate: Basic realm="Authorization Required"
//< Date: Tue, 04 Apr 2023 14:56:52 GMT
//< Content-Length: 0
//<
//  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
//* Connection #0 to host localhost left intact
//最终也没有获取到用户信息
