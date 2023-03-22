package limiter

import (
	"sync"
	"time"
)

// TokenBucketLimiter 令牌桶限流器
type TokenBucketLimiter struct {
	mu        sync.Mutex    // 互斥锁
	rate      time.Duration // 产生令牌速率
	capacity  int           // 桶容量
	remaining int           // 剩余令牌数量
	last      time.Time     // 上次产生令牌时间
}

// NewTokenBucketLimiter 创建一个新的令牌桶限流器
func NewTokenBucketLimiter(rate time.Duration, capacity int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		rate:     rate,
		capacity: capacity,
		last:     time.Now(),
	}
}

// allow 判断是否允许通过请求，返回值是bool类型

func (l *TokenBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	tokens := int(now.Sub(l.last) / l.rate)
	l.remaining += tokens
	if l.remaining > l.capacity {
		l.remaining = l.capacity
	}
	l.last = now

	if l.remaining <= 0 {
		return false
	}

	l.remaining--
	return true
}

/*该方法首先使用互斥锁来保证线程安全。然后，根据当前时间和上次产生令牌的时间，计算出在这段时间内应该产生的令牌数量，
并将其添加到剩余令牌数量中。如果剩余令牌数量超过了桶容量，则将其设置为桶容量。
接下来，判断剩余令牌数量是否小于等于 0。如果是，则拒绝该请求；
否则，从剩余令牌数量中减去一个令牌，并允许该请求通过。*/
