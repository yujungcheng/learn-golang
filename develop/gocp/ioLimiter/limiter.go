package ioLimiter

import (
	"time"
	"context"
	"golang.org/x/time/rate"
)


type Limiter struct {
	limiter *rate.Limiter
	ctx context.Context
}

func (l *Limiter) SetLimiter(bytesPerSecond float64) {
	var burstLimit int = 1073741824  // 1GB
	l.ctx = context.Background()
	l.limiter = rate.NewLimiter(rate.Limit(bytesPerSecond), burstLimit)
	l.limiter.AllowN(time.Now(), burstLimit)
}

func (l *Limiter) WaitN(n int) error {
	if err := l.limiter.WaitN(l.ctx, n); err != nil {
		return err
	}
	return nil
}