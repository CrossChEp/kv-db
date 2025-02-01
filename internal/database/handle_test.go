package database

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CrossChEp/kv-db/internal/entity"
	testUtils "github.com/CrossChEp/kv-db/internal/utils/test_utils"
)

func TestDatabase_Handle(t *testing.T) {
	type fields struct {
		parser Parser
		engine Engine
	}

	tt := []struct {
		name       string
		in         string
		out        string
		setupMocks func(ctrl *gomock.Controller) fields
	}{
		{
			name: "successful_set",
			in:   "SET key value\n",
			out:  "[ok] Value was set",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "SET key value").
					Return(entity.Query{
						Type: entity.TypeSet,
						Key:  "key",
						Val:  "value",
					}, nil).Times(1)

				engine := NewMockEngine(ctrl)
				engine.EXPECT().Set("key", "value").Times(1)

				return fields{
					parser: parser,
					engine: engine,
				}
			},
		},
		{
			name: "successful_get",
			in:   "GET key\n",
			out:  "[ok] value",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "GET key").
					Return(entity.Query{
						Type: entity.TypeGet,
						Key:  "key",
						Val:  "",
					}, nil).Times(1)

				engine := NewMockEngine(ctrl)
				engine.EXPECT().Get("key").Return("value", true).Times(1)

				return fields{
					parser: parser,
					engine: engine,
				}
			},
		},
		{
			name: "get_by_non_existing_key",
			in:   "GET key\n",
			out:  "[err] no such key",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "GET key").
					Return(entity.Query{
						Type: entity.TypeGet,
						Key:  "key",
						Val:  "",
					}, nil).Times(1)

				engine := NewMockEngine(ctrl)
				engine.EXPECT().Get("key").Return("", false).Times(1)

				return fields{
					parser: parser,
					engine: engine,
				}
			},
		},
		{
			name: "successful_delete",
			in:   "DEL key\n",
			out:  "[ok] Value was deleted",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "DEL key").
					Return(entity.Query{
						Type: entity.TypeDel,
						Key:  "key",
						Val:  "",
					}, nil).Times(1)

				engine := NewMockEngine(ctrl)
				engine.EXPECT().Del("key").Times(1)

				return fields{
					parser: parser,
					engine: engine,
				}
			},
		},
		{
			name: "err_from_parser",
			in:   "DEL key\n",
			out:  "[err] some err",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "DEL key").
					Return(entity.Query{}, errors.New("some err")).Times(1)

				return fields{
					parser: parser,
				}
			},
		},
		{
			name: "parser_returned_invalid_type",
			in:   "DEL key\n",
			out:  "[err] parsing error",
			setupMocks: func(ctrl *gomock.Controller) fields {
				parser := NewMockParser(ctrl)
				parser.EXPECT().Parse(gomock.Any(), "DEL key").
					Return(entity.Query{
						Type: 133,
						Key:  "key",
						Val:  "",
					}, nil).Times(1)

				return fields{
					parser: parser,
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			f := tc.setupMocks(ctrl)

			db := New(testUtils.StubLogger{}, f.parser, f.engine)

			res := db.Handle(context.Background(), tc.in)

			require.Equal(t, tc.out, res)
		})
	}
}
