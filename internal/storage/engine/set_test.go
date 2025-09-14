package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngine_Set(t *testing.T) {
	tt := []struct {
		name            string
		key             string
		val             string
		resStorage      map[string]string
		originalStorage map[string]string
	}{
		{
			name:            "successful_setting",
			key:             "key",
			val:             "val",
			originalStorage: map[string]string{},
			resStorage: map[string]string{
				"key": "val",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			e.storage = tc.originalStorage

			e.Set(tc.key, tc.val)

			require.Equal(t, tc.resStorage, e.storage)
		})
	}
}
