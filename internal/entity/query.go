package entity

import "errors"

type QueryType int

const (
	TypeSet QueryType = iota
	TypeGet
	TypeDel

	ArgPattern = `^[\w\/\*]+$`
)

var (
	ErrEmptyQuery = errors.New("empty query provided")
)

type Query struct {
	Type QueryType
	Key  string
	Val  string
}
