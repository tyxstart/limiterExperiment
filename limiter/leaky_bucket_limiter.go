package leaky_bucket_limiter

import (
	"sync"
	"time"
)

// LeakyBucketLimiter 漏桶限流器
type LeakyBucketLimiter struct {
	mu        sync.Mutex    // 互斥锁
	rate      time.Duration // 漏水速率
	capacity  int           // 桶容量
	remaining int           // 剩余水量
	last      time.Time     // 上次漏水时间
}

// NewLeakyBucketLimiter 创建一个新的漏桶限流器
func NewLeakyBucketLimiter(rate time.Duration, capacity int) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		rate:     rate,
		capacity: capacity,
		last:     time.Now(),
	}
}

// allow 判断是否允许通过请求，返回值是bool类型
func (l *LeakyBucketLimiter) allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	l.remaining -= int(now.Sub(l.last) / l.rate)
	if l.remaining < 0 {
		l.remaining = 0
	}
	l.last = now

	if l.remaining >= l.capacity {
		return false
	}

	l.remaining++
	return true
}
