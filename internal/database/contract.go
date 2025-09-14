package database

import (
	"context"

	"github.com/CrossChEp/kv-db/internal/entity"
)

//go:generate mockgen -source $GOFILE -destination contract_mocks.go -package $GOPACKAGE
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
