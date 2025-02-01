package engine

func (e *Engine) Set(key, val string) {
	e.storage[key] = val
}
