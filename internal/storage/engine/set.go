package engine

import "github.com/CrossChEp/kv-db/internal/utils/lock"

func (e *Engine) Set(key, val string) {
	lock.WithLock(&e.mu, func() {
		e.storage[key] = val
	})
}
