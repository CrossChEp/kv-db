package compute

import (
	"context"
	"github.com/CrossChEp/kv-db/internal/entity"
	testUtils "github.com/CrossChEp/kv-db/internal/utils/test_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tt := []struct {
		name        string
		in          string
		out         entity.Query
		expectedErr error
	}{
		{
			name: "successful_parse_of_set_command",
			in:   "SET key value",
			out: entity.Query{
				Type: entity.TypeSet,
				Key:  "key",
				Val:  "value",
			},
			expectedErr: nil,
		},
		{
			name:        "one_argument_in_set_command",
			in:          "SET key",
			out:         entity.Query{},
			expectedErr: entity.ErrInvalidAmountOfArgs,
		},
		{
			name:        "invalid_symbols_in_argument",
			in:          "SET \n ``",
			out:         entity.Query{},
			expectedErr: entity.ErrInvalidArgumentSymbols,
		},
		{
			name: "successful_parse_of_get_command",
			in:   "GET key",
			out: entity.Query{
				Type: entity.TypeGet,
				Key:  "key",
				Val:  "",
			},
			expectedErr: nil,
		},
		{
			name:        "two_arguments_in_set_command",
			in:          "GET key val",
			out:         entity.Query{},
			expectedErr: entity.ErrInvalidAmountOfArgs,
		},
		{
			name: "successful_parse_of_del_command",
			in:   "DEL key",
			out: entity.Query{
				Type: entity.TypeDel,
				Key:  "key",
				Val:  "",
			},
			expectedErr: nil,
		},
		{
			name:        "two_arguments_in_del_command",
			in:          "DEL key val",
			out:         entity.Query{},
			expectedErr: entity.ErrInvalidAmountOfArgs,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := New(testUtils.StubLogger{})

			res, err := p.Parse(context.Background(), tc.in)

			require.Equal(t, tc.out, res)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
