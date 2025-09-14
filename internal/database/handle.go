package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/CrossChEp/kv-db/internal/fields"
)

func (db *Database) Handle(ctx context.Context, q string) string {
	q = strings.ReplaceAll(q, "\n", "")

	ctx = db.log.WithFields(ctx, fields.WithQuery(q))

	query, err := db.parser.Parse(ctx, q)
	if err != nil {
		return fmt.Sprintf("[err] %v", err)
	}

	switch query.Type {
	case entity.TypeGet:
		return db.get(query.Key)
	case entity.TypeSet:
		return db.set(query.Key, query.Val)
	case entity.TypeDel:
		return db.del(query.Key)
	default:
		return "[err] parsing error"
	}
}

func (db *Database) get(key string) string {
	val, ok := db.engine.Get(key)
	if !ok {
		return fmt.Sprintf("[err] no such key")
	}

	return fmt.Sprintf("[ok] %s", val)
}

func (db *Database) set(key, val string) string {
	db.engine.Set(key, val)

	return fmt.Sprint("[ok] Value was set")
}

func (db *Database) del(key string) string {
	db.engine.Del(key)

	return fmt.Sprint("[ok] Value was deleted")
}
