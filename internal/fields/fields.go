package fields

import "github.com/CrossChEp/kv-db/internal/entity"

type Field map[string]interface{}

const (
	fieldError      = "error"
	fieldQuery      = "query"
	fieldArgs       = "args"
	fieldCommand    = "command"
	fieldAddress    = "address"
	fieldPanic      = "panic"
	fieldReadCount  = "readCount"
	fieldBufferSize = "bufferSize"
)

func Join(fields ...Field) Field {
	field := make(Field)

	for _, f := range fields {
		for k, v := range f {
			field[k] = v
		}
	}

	return field
}

func WithError(err error) Field {
	if err == nil {
		return nil
	}

	return Field{
		fieldError: err.Error(),
	}
}

func WithQuery(query string) Field {
	return Field{
		fieldQuery: query,
	}
}

func WithArgs(args []string) Field {
	return Field{
		fieldArgs: args,
	}
}

func WithCommand(command entity.Command) Field {
	return Field{
		fieldCommand: command,
	}
}

func WithAddress(addr string) Field {
	return Field{
		fieldAddress: addr,
	}
}

func WithPanic(p any) Field {
	return Field{
		fieldPanic: p,
	}
}

func WithReadCount(count int) Field {
	return Field{
		fieldReadCount: count,
	}
}

func WithBufferSize(size int) Field {
	return Field{
		fieldBufferSize: size,
	}
}
