package v1

import (
	"github.com/gin-gonic/gin"
	"limiterExperiment/config"
	"limiterExperiment/middleware"
	"net/http"
)

func Test(c *gin.Context) {
	//执行限流组件，判断是否开启限流，开启限流后时候达到阀值
	if middleware.LimitHandler() == false {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "你已经被限流",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    config.LimiterState,
			"size":    config.SizeInt,
			"limit":   config.LimitInt,
			"message": "访问成功",
		})
	}

}
