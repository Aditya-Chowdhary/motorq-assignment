package server

import (
	"motorq-assignment/internal/merrors"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func (s *Server) rateLimit() gin.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := realip.FromRequest(c.Request)

		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{

				limiter: rate.NewLimiter(rate.Limit(s.limiter.rate), s.limiter.burst),
			}
		}
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			merrors.TooManyRequests(c, "rate limit")
			return
		}
		mu.Unlock()
		c.Next()
	}
}
