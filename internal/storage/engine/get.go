package engine

func (e *Engine) Get(key string) (string, bool) {
	val, ok := e.storage[key]

	return val, ok
}
