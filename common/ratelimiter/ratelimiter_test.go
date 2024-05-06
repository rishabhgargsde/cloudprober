package ratelimiter

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	sM := NewRateLimiter(30 * 24 * time.Hour)
	called := false
	sM.Do(func() { called = true })
	if !called {
		t.Errorf("RateLimiter.Do() didn't call function. Expected first call to always happen.")
	}
	if sM.Counter != 0 {
		t.Errorf("RateLimiter.Counter != 0, got %d", sM.Counter)
	}
	called = false
	sM.Do(func() { called = true })
	if called {
		t.Errorf("RateLimiter.Do() didn't wait a month. Expected second call to not happen (1 month duration).")
	}
	if sM.Counter != 1 {
		t.Errorf("RateLimiter.Counter != 1, got %d", sM.Counter)
	}

	sN := NewRateLimiter(1 * time.Nanosecond)
	called = false
	sN.Do(func() { called = true })
	if !called {
		t.Errorf("RateLimiter.Do() didn't call function. Expected first call to always happen.")
	}
	if sN.Counter != 0 {
		t.Errorf("RateLimiter.Counter != 0, got %d", sM.Counter)
	}
	time.Sleep(1 * time.Millisecond)
	called = false
	sN.Do(func() { called = true })
	if !called {
		t.Errorf("RateLimiter.Do() didn't call function. Expected second call to happen after 1ns.")
	}
	if sN.Counter != 0 {
		t.Errorf("RateLimiter.Counter != 0, got %d", sM.Counter)
	}
}
