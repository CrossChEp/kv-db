package engine

type Engine struct {
	storage map[string]string
}

func New() *Engine {
	return &Engine{
		storage: make(map[string]string),
	}
}
