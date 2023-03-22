package middleware

import (
	"limiterExperiment/config"
	"limiterExperiment/limiter"
)

var FWLimiter *limiter.FixedWindowLimiter
var SWLimiter *limiter.SlidingWindowLimiter
var SLLimiter *limiter.SlidingLogLimiter
var LBLimiter *limiter.LeakyBucketLimiter
var TBLimiter *limiter.TokenBucketLimiter

func LimitHandler() bool {
	if 0 != config.LimiterState { //开启限流
		switch config.LimiterState {
		case 1:
			if FWLimiter.Allow() == false {
				return false
			}
		case 2:
			if SWLimiter.Allow() == false {
				return false
			}
		case 3:
			if SLLimiter.Allow() == false {
				return false
			}
		case 4:
			if LBLimiter.Allow() == false {
				return false
			}
		case 5:
			if TBLimiter.Allow() == false {
				return false
			}
		default:
			return true
		}
	} else { //未开限流
		return true
	}
	return true
}
