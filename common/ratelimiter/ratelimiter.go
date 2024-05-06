// Package ratelimiter implementes a helper to perform an action only some of the time.
package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter ensures that actions are only performed with a minimum duration between them.
type RateLimiter struct {
	mu   sync.Mutex
	last time.Time

	// Count of how often the function would have been called.
	Counter uint64

	// Minimum duration between calls.
	Duration time.Duration
}

// NewRateLimiter returns a RateLimiter object that ensures a minimum duration d between actions.
func NewRateLimiter(d time.Duration) *RateLimiter {
	return &RateLimiter{Counter: 0, Duration: d}
}

// Do rate limit calls of function f.
func (e *RateLimiter) Do(f func()) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if time.Since(e.last) >= e.Duration {
		e.last = time.Now()
		f()
		e.Counter = 0
	} else {
		e.Counter++
	}
}
