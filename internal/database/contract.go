package database

import (
	"context"

	"github.com/CrossChEp/kv-db/internal/entity"
)

type (
	Parser interface {
		Parse(ctx context.Context, query string) (entity.Query, error)
	}

	Engine interface {
		Del(key string)
		Get(key string) (string, bool)
		Set(key, val string)
	}
)
