package ioc

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Resolve(t *testing.T) {
	var (
		scope        = NewScope(context.Background())
		newScope     = NewScope(context.Background())
		invalidScope = context.WithValue(context.Background(), ScopeDependenciesKey{}, nil)
	)
	err := Register(scope, "Test", func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, errors.New("invalid args")
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, errors.New("invalid args")
		}
		return "test" + "." + value, nil
	})
	require.NoError(t, err)
	fromParentScope := NewFromParent(scope, newScope)

	type args struct {
		ctx  context.Context
		key  string
		args []any
	}
	tests := map[string]struct {
		args    args
		want    string
		wantErr bool
	}{
		"зависимость резолвится корректно": {
			args: args{
				ctx: scope,
				key: "Test",
				args: []any{
					"value",
				},
			},
			want:    "test.value",
			wantErr: false,
		},
		"зависимость резолвится корректно, from parent scope": {
			args: args{
				ctx: fromParentScope,
				key: "Test",
				args: []any{
					"value",
				},
			},
			want:    "test.value",
			wantErr: false,
		},
		"зависимость не найдена": {
			args: args{
				ctx: newScope,
				key: "Test",
				args: []any{
					"value",
				},
			},
			wantErr: true,
		},
		"зависимость не найдена, nil": {
			args: args{
				ctx: invalidScope,
				key: "Test",
				args: []any{
					"value",
				},
			},
			wantErr: true,
		},
		"не найдены зависимости, невалидный скоуп": {
			args: args{
				ctx: context.Background(),
				key: "Test",
				args: []any{
					"value",
				},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Resolve(tt.args.ctx, tt.args.key, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			require.Equal(t, tt.want, got)
		})
	}
}
