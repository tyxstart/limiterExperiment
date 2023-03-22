package router

import (
	"github.com/gin-gonic/gin"
	"limiterExperiment/config"
	"net/http"
)

func InitRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		})
	}

	//// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	//// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	//router.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.Run(config.HttpPort)
}
