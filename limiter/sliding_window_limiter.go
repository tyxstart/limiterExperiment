package sliding_window_limiter

import (
	"sync"
	"time"
)

// SlidingWindowLimiter 滑动窗口限流器
type SlidingWindowLimiter struct {
	mu         sync.Mutex    // 互斥锁
	windowSize time.Duration // 窗口大小
	limit      int           // 限制数量
	slots      []int         // 时间槽
	cursor     int           // 当前时间槽位置
}

// NewSlidingWindowLimiter 创建一个新的滑动窗口限流器
func NewSlidingWindowLimiter(windowSize time.Duration, limit int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		windowSize: windowSize,
		limit:      limit,
		slots:      make([]int, limit),
		cursor:     0,
	}
}

// allow 判断是否允许通过请求，返回值是bool类型
func (l *SlidingWindowLimiter) allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	l.cursor = int(now.UnixNano()/int64(l.windowSize)) % l.limit

	if l.slots[l.cursor] > 0 && now.Sub(time.Unix(0, int64(l.slots[l.cursor]))) < l.windowSize {
		return false
	}

	l.slots[l.cursor] = int(now.UnixNano())
	return true
}
