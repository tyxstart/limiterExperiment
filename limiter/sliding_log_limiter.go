package limiter

import "time"

// SlidingLogLimiter 滑动日志限流器
type SlidingLogLimiter struct {
	windowSize time.Duration // 窗口大小
	limit      int           // 限制数量
	logs       []int64       // 请求日志
}

// NewSlidingLogLimiter 创建一个新的滑动日志限流器
func NewSlidingLogLimiter(windowSize time.Duration, limit int) *SlidingLogLimiter {
	return &SlidingLogLimiter{
		windowSize: windowSize,
		limit:      limit,
		logs:       make([]int64, 0),
	}
}

// allow 判断是否允许通过请求，返回值是bool类型
func (l *SlidingLogLimiter) Allow() bool {
	now := time.Now()
	deadline := now.Add(-l.windowSize).UnixNano()

	for i := 0; i < len(l.logs); i++ {
		if l.logs[i] > deadline {
			l.logs = l.logs[i:]
			break
		}
	}

	if len(l.logs) >= l.limit {
		return false
	}

	l.logs = append(l.logs, now.UnixNano())
	return true
}
