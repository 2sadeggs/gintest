package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		message2 := c.PostForm("hhahahah23333") //自定义个别的key 看看效果
		nick := c.DefaultPostForm("nick", "anonymous")
		nick2 := c.DefaultPostForm("nick2", "jiaoshahaone") //反正都是自定义 不是来自实际表单 随便写啦

		/*c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})*/
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message2,
			"nick":    nick2,
		})
	})
	router.Run("localhost:8080")
}

//curl -v --form user=user --form password=password http://localhost:8080/form_post
//6361{"message":"","nick":"anonymous","status":"posted"}
/*简单的表单post请求 看不出太多信息*/
//json 返回三列信息：状态status 固定的posted、消息message 这个是从表单中获得的（可修改测试效果）、昵称nick 也是从表单中获得的

/*shibinbin@shibinbin00 MINGW64 /d/GolandProjects/gintest (main)
$ curl --form message=xiaoxiha --form nick=nicheng  http://localhost:8080/form_post
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--100   307  100    57  100   250   9319  40876 --:--:-- --:--:-- --:--:--  299k
//换行说明
{"message":"xiaoxiha","nick":"nicheng","status":"posted"}
*/