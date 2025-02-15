package common

import (
	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseBufferSize(t *testing.T) {
	tt := []struct {
		name    string
		in      string
		out     int
		wantErr error
	}{
		{
			name:    "success_no_format",
			in:      "10",
			out:     10,
			wantErr: nil,
		},
		{
			name:    "success_byte",
			in:      "6B",
			out:     6,
			wantErr: nil,
		},
		{
			name:    "success_kilobyte",
			in:      "4KB",
			out:     4096,
			wantErr: nil,
		},
		{
			name:    "success_megabyte",
			in:      "5mb",
			out:     5242880,
			wantErr: nil,
		},
		{
			name:    "success_gigabyte",
			in:      "2GB",
			out:     2147483648,
			wantErr: nil,
		},
		{
			name:    "empty_req",
			in:      "",
			out:     0,
			wantErr: entity.ErrInvalidBufferSizeFormat,
		},
		{
			name:    "invalid_req",
			in:      "123-1=23q",
			out:     0,
			wantErr: entity.ErrInvalidBufferSizeFormat,
		},
		{
			name:    "invalid_suffix",
			in:      "123llhg",
			out:     0,
			wantErr: entity.ErrInvalidBufferSizeSuffix,
		},
	}

	t.Parallel()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ParseBufferSize(tc.in)

			require.Equal(t, tc.out, res)
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}
