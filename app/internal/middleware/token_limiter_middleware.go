package middleware

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	burst   = 5
	rate    = 5
	seconds = 2
)

type TokenLimiterMiddleware struct {
	redisConf redis.RedisConf
}

func NewTokenLimiterMiddleware(redisConf redis.RedisConf) *TokenLimiterMiddleware {
	return &TokenLimiterMiddleware{
		redisConf: redisConf,
	}
}

func (m *TokenLimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("tokenlimiter middleware")
		// Passthrough to next handler if need
		fmt.Println("=====", m.redisConf.Host)
		store, err := redis.NewRedis(m.redisConf)
		if err != nil {
			fmt.Println("err:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// New tokenLimiter
		var allowed, denied int32
		limiter := limit.NewTokenLimiter(rate, burst, store, "rate-token-limiter")
		if limiter.Allow() {
			atomic.AddInt32(&allowed, 1)
		} else {
			atomic.AddInt32(&denied, 1)
			http.Error(w, "token limiter deny", http.StatusBadRequest)
			return
		}

		fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)

		next(w, r)
	}
}
