package v1

import (
	"github.com/gin-gonic/gin"
	"limiterExperiment/config"
	"limiterExperiment/limiter"
	"limiterExperiment/middleware"
	"net/http"
	"strconv"
	"time"
)

func SetLimiter(c *gin.Context) {
	//设置限流算法   创建对应限流对象
	codeStr := c.DefaultPostForm("code", "0")
	config.LimiterState, _ = strconv.Atoi(codeStr)
	sizeStr := c.DefaultPostForm("size", "10") //限制数量
	config.SizeInt, _ = strconv.Atoi(sizeStr)
	limitStr := c.DefaultPostForm("limit", "1") //限制大小
	config.LimitInt, _ = strconv.Atoi(limitStr)
	if config.LimiterState == 0 {
		return
	}
	switch config.LimiterState {
	case 0:
		config.LimiterState = 0 //0不开启限流；1固定窗口；2滑动窗口；3滑动日志；4漏桶；5令牌桶

		config.FWLimit = 10
		config.FWInterval = 1 * time.Second

		config.SWLimit = 10
		config.SWWindowSize = 1 * time.Second

		config.SLLimit = 10
		config.SLWindowSize = 1 * time.Second

		config.LBCapacity = 10
		config.LBRate = 1 * time.Second

		config.TBCapacity = 10
		config.TBRate = 1 * time.Second
	case 1:
		config.FWLimit = config.SizeInt
		config.FWInterval = time.Duration(config.LimitInt) * time.Second
		middleware.FWLimiter = limiter.NewFixedWindowRateLimiter(config.FWLimit, config.FWInterval)
	case 2:
		config.SWLimit = config.SizeInt
		config.SWWindowSize = time.Duration(config.LimitInt) * time.Second
		middleware.SWLimiter = limiter.NewSlidingWindowLimiter(config.SWWindowSize, config.SWLimit)
	case 3:
		config.SLLimit = config.SizeInt
		config.SLWindowSize = time.Duration(config.LimitInt) * time.Second
		middleware.SLLimiter = limiter.NewSlidingLogLimiter(config.SWWindowSize, config.SWLimit)
	case 4:
		config.LBCapacity = config.SizeInt
		config.LBRate = time.Duration(config.LimitInt) * time.Second
		middleware.LBLimiter = limiter.NewLeakyBucketLimiter(config.SWWindowSize, config.SWLimit)
	case 5:
		config.TBCapacity = config.SizeInt
		config.TBRate = time.Duration(config.LimitInt) * time.Second
		middleware.TBLimiter = limiter.NewTokenBucketLimiter(config.SWWindowSize, config.SWLimit)

	default:
		c.JSON(http.StatusForbidden, gin.H{
			"code":    config.LimiterState,
			"size":    config.SizeInt,
			"limit":   config.LimitInt,
			"message": "设置失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    config.LimiterState,
		"size":    config.SizeInt,
		"limit":   config.LimitInt,
		"message": "设置成功",
	})
}
