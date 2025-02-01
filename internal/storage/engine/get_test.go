package engine

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEngine_Get(t *testing.T) {
	type resp struct {
		res string
		ok  bool
	}

	tt := []struct {
		name    string
		in      string
		out     resp
		storage map[string]string
	}{
		{
			name: "successful_getting",
			in:   "key",
			out: resp{
				res: "123",
				ok:  true,
			},
			storage: map[string]string{
				"key": "123",
			},
		},
		{
			name: "get_non_existing_key",
			in:   "key1",
			out: resp{
				res: "",
				ok:  false,
			},
			storage: map[string]string{
				"key": "123",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			e.storage = tc.storage

			res, ok := e.Get(tc.in)

			require.Equal(t, tc.out.res, res)
			require.Equal(t, tc.out.ok, ok)
		})
	}
}
