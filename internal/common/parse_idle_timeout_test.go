package common

import (
	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseIdleTimeout(t *testing.T) {
	tt := []struct {
		name    string
		in      string
		out     time.Duration
		wantErr error
	}{
		{
			name:    "success_millisecond",
			in:      "6ms",
			out:     6 * time.Millisecond,
			wantErr: nil,
		},
		{
			name:    "success_second",
			in:      "10S",
			out:     10 * time.Second,
			wantErr: nil,
		},
		{
			name:    "success_minute",
			in:      "5m",
			out:     5 * time.Minute,
			wantErr: nil,
		},
		{
			name:    "success_hour",
			in:      "2h",
			out:     2 * time.Hour,
			wantErr: nil,
		},
		{
			name:    "empty_req",
			in:      "",
			out:     0,
			wantErr: entity.ErrInvalidTimeoutFormat,
		},
		{
			name:    "invalid_req",
			in:      "123-1=23q",
			out:     0,
			wantErr: entity.ErrInvalidTimeoutFormat,
		},
		{
			name:    "invalid_suffix",
			in:      "123gg",
			out:     0,
			wantErr: entity.ErrInvalidTimeSuffixFormat,
		},
	}

	t.Parallel()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ParseIdleTimeout(tc.in)

			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.out, res)
		})
	}
}
