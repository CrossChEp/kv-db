package entity

import "errors"

type Command string

const (
	CommandSet Command = "SET"
	CommandGet Command = "GET"
	CommandDel Command = "DEL"
)

var (
	Commands = []Command{CommandGet, CommandSet, CommandDel}

	CommandToType = map[Command]QueryType{
		CommandGet: TypeGet,
		CommandSet: TypeSet,
		CommandDel: TypeDel,
	}
	CommandToArgQuantity = map[Command]int{
		CommandSet: 2,
		CommandDel: 1,
		CommandGet: 1,
	}
)

var (
	ErrInvalidCommand         = errors.New("provided command is invalid")
	ErrInvalidAmountOfArgs    = errors.New("invalid amount of args provided")
	ErrInvalidArgumentSymbols = errors.New(
		"invalid argument symbols, you can enter only letters, numbers, '/', '*' and '_'",
	)
)
