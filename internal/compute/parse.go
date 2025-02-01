package compute

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/CrossChEp/kv-db/internal/fields"
)

func (p *Parser) Parse(ctx context.Context, query string) (entity.Query, error) {
	var err error

	defer func() {
		if err != nil {
			p.log.Error(
				p.log.WithFields(ctx, fields.WithError(err)),
				"Couldn't parse query",
			)
		}
	}()

	if len(strings.TrimSpace(query)) == 0 {
		return entity.Query{}, entity.ErrEmptyQuery
	}

	parts := strings.Split(query, " ")
	command := entity.Command(parts[0])
	args := parts[1:]

	ctx = p.log.WithFields(ctx, fields.Join())

	if err = validateQuery(command, args); err != nil {
		return entity.Query{}, err
	}

	var (
		val       string
		queryType entity.QueryType = entity.CommandToType[command]
	)

	if queryType == entity.TypeSet {
		val = args[1]
	}

	return entity.Query{
		Type: queryType,
		Key:  args[0],
		Val:  val,
	}, nil
}

func validateQuery(command entity.Command, args []string) error {
	if !slices.Contains(entity.Commands, command) {
		return fmt.Errorf("%w, allowed commands: %+v", entity.ErrInvalidCommand, entity.Commands)
	}

	if argQuantity := entity.CommandToArgQuantity[command]; len(args) != argQuantity {
		return fmt.Errorf(
			"%w:  command %s requires %d amount of arguments",
			entity.ErrInvalidAmountOfArgs,
			command,
			argQuantity,
		)
	}

	for i, arg := range args {
		match, _ := regexp.Match(entity.ArgPattern, []byte(arg))
		if !match {
			return fmt.Errorf(
				"arg %d conrains invalid symbols, you can enter only letters, numbers, '/', '*' and '_'",
				i+1,
			)
		}
	}

	return nil
}
