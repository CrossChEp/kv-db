package engine

import "github.com/CrossChEp/kv-db/internal/utils/lock"

func (e *Engine) Get(key string) (val string, ok bool) {
	lock.WithLock(&e.mu, func() {
		val, ok = e.storage[key]
	})

	return
}
