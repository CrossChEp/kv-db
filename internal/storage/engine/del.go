package engine

import "github.com/CrossChEp/kv-db/internal/utils/lock"

func (e *Engine) Del(key string) {
	lock.WithLock(&e.mu, func() {
		delete(e.storage, key)
	})
}
