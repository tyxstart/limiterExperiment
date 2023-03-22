package fixed_window_limiter

import (
	"sync"
	"time"
)

// FixedWindowLimiter 结构
type FixedWindowLimiter struct {
	mu       sync.Mutex    //锁
	limit    int           //限流阈值
	interval time.Duration //时间窗口长度
	count    int           //当前时间窗口内已经通过的请求数量
	lastTime time.Time     //上一次请求通过时所处的时间窗口的起始时间
}

// NewFixedWindowRateLimiter 创建新的限流器
func NewFixedWindowRateLimiter(limit int, interval time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:    limit,
		interval: interval,
	}
}

func (r *FixedWindowLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.lastTime) >= r.interval {
		r.count = 0
		r.lastTime = now
	}

	if r.count < r.limit {
		r.count++
		return true
	}
	return false
}
