package fields

type Field map[string]interface{}

const (
	fieldError = "error"
	fieldQuery = "query"
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
