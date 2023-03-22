package v1

import (
	"github.com/gin-gonic/gin"
	"limiterExperiment/middleware"
)

func Test(c *gin.Context) {
	//执行限流组件，判断是否开启限流，开启限流后时候达到阀值
	middleware.LimitHandler()

	//
}
