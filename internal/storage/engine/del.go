package engine

func (e *Engine) Del(key string) {
	delete(e.storage, key)
}
