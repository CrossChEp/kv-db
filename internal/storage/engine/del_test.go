package engine

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEngine_Del(t *testing.T) {
	tt := []struct {
		name            string
		in              string
		resStorage      map[string]string
		originalStorage map[string]string
	}{
		{
			name:       "successful_deleting",
			in:         "key",
			resStorage: map[string]string{},
			originalStorage: map[string]string{
				"key": "123",
			},
		},
		{
			name: "deleting_with_non_existing_key",
			in:   "key1",
			resStorage: map[string]string{
				"key": "123",
			},
			originalStorage: map[string]string{
				"key": "123",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			e.storage = tc.originalStorage

			e.Del(tc.in)

			require.Equal(t, tc.resStorage, e.storage)
		})
	}
}
