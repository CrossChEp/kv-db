package engine

import "sync"

type Engine struct {
	mu      sync.Mutex
	storage map[string]string
}

func New() *Engine {
	return &Engine{
		mu:      sync.Mutex{},
		storage: make(map[string]string),
	}
}
