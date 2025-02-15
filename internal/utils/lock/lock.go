package lock

import "sync"

func WithLock(mu *sync.Mutex, fn func()) {
	if fn == nil {
		return
	}

	mu.Lock()
	defer func() {
		if p := recover(); p != nil {
			return
		}

		mu.Unlock()
	}()

	fn()
}
