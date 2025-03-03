package ioc

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	move2 "github.com/nika-gromova/o-architecture-patterns/game/macro_command/move"
	"github.com/nika-gromova/o-architecture-patterns/game/move"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	moveCommandKey          = "IoC.MoveCommand"
	moveCommandConstructor1 = Constructor(func(args ...any) (any, error) {
		return &move.MoveCommand{
			Obj: args[0].(move.MovingObject),
		}, nil
	})
	moveCommandConstructor2 = Constructor(func(args ...any) (any, error) {
		return &move2.MoveWithFuelCommand{
			Obj: args[0].(move2.MovingWithFuelObj),
		}, nil
	})
)

func TestContainer_Register(t *testing.T) {
	var (
		c     = New()
		scope = c.NewScope(context.Background())
	)
	type args struct {
		ctx  context.Context
		key  string
		args []any
	}
	tests := map[string]struct {
		args    args
		want    *RegisterCommand
		wantErr bool
	}{
		"зависимость регистрируется корректно": {
			args: args{
				ctx: scope,
				key: "IoC.Register",
				args: []any{
					scope,
					moveCommandKey,
					moveCommandConstructor1,
				},
			},
			want: &RegisterCommand{
				scope:       scope,
				key:         moveCommandKey,
				constructor: moveCommandConstructor1,
			},
			wantErr: false,
		},
		"ошибка количества параметров": {
			args: args{
				ctx: scope,
				key: "IoC.Register",
				args: []any{
					scope,
					moveCommandKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		"ошибка типа (1й параметр)": {
			args: args{
				ctx: scope,
				key: "IoC.Register",
				args: []any{
					"test",
					moveCommandKey,
					moveCommandConstructor1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		"ошибка типа (2й параметр)": {
			args: args{
				ctx: scope,
				key: "IoC.Register",
				args: []any{
					scope,
					123,
					moveCommandConstructor1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		"ошибка типа (3й параметр)": {
			args: args{
				ctx: scope,
				key: "IoC.Register",
				args: []any{
					scope,
					moveCommandKey,
					func() error { return nil },
				},
			},
			want:    nil,
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

			require.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(got))
			gotTyped := got.(*RegisterCommand)
			require.Equal(t, tt.want.key, gotTyped.key)
			require.Equal(t, tt.want.scope, gotTyped.scope)

			err = gotTyped.Execute(context.Background())
			require.NoError(t, err)

			dependencies, ok := scope.Value(ScopeDependenciesKey{}).(*Dependencies)
			require.True(t, ok)
			require.NotNil(t, (*dependencies)[moveCommandKey])
		})
	}
}

func TestContainer_Resolve(t *testing.T) {
	var (
		ctrl       = gomock.NewController(t)
		movingMock = mocks.NewMockMovingObject(ctrl)
		c          = New()
		scope      = c.NewScope(context.Background())
		newScope   = c.NewScope(context.Background())
	)
	cmd, err := Resolve(scope, "IoC.Register", scope, moveCommandKey, moveCommandConstructor1)
	require.NoError(t, err)
	cmdTyped, ok := cmd.(*RegisterCommand)
	require.True(t, ok)
	require.NoError(t, cmdTyped.Execute(context.Background()))

	type args struct {
		ctx  context.Context
		key  string
		args []any
	}
	tests := map[string]struct {
		args    args
		want    *move.MoveCommand
		wantErr bool
	}{
		"зависимость резолвится корректно": {
			args: args{
				ctx: scope,
				key: moveCommandKey,
				args: []any{
					movingMock,
				},
			},
			want: &move.MoveCommand{
				Obj: movingMock,
			},
			wantErr: false,
		},
		"зависимость не найдена": {
			args: args{
				ctx: newScope,
				key: moveCommandKey,
				args: []any{
					movingMock,
				},
			},
			wantErr: true,
		},
		"не найдены зависимости, невалидный скоуп": {
			args: args{
				ctx: context.Background(),
				key: moveCommandKey,
				args: []any{
					movingMock,
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

func TestContainer_Goroutines(t *testing.T) {
	var (
		ctrl               = gomock.NewController(t)
		movingMock         = mocks.NewMockMovingObject(ctrl)
		movingWithFuelMock = mocks.NewMockMovingWithFuelObj(ctrl)
		c                  = New()
		game1Scope         = c.NewScope(context.Background())
		game2Scope         = c.NewScope(context.Background())
	)

	t.Run("зависимости регистрируются и резолвятся корректно", func(t *testing.T) {
		movingMock.EXPECT().GetLocation().Return(base.Vector{}, true)
		movingMock.EXPECT().GetVelocity().Return(base.Vector{}, true)
		movingMock.EXPECT().SetLocation(gomock.Any()).Return(true)

		movingWithFuelMock.EXPECT().GetFuel().Return(base.FuelInfo{}, true).Times(2)
		movingWithFuelMock.EXPECT().GetLocation().Return(base.Vector{}, true)
		movingWithFuelMock.EXPECT().GetVelocity().Return(base.Vector{}, true)
		movingWithFuelMock.EXPECT().SetLocation(gomock.Any()).Return(true)
		movingWithFuelMock.EXPECT().SetFuel(gomock.Any()).Return(true)

		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()

			res, err := testMoveObject(c, game1Scope, moveCommandConstructor1, movingMock)

			require.NoError(t, err)
			moveCommand, ok := res.(*move.MoveCommand)
			require.True(t, ok)
			err = moveCommand.Execute()
			require.NoError(t, err)
		}()

		go func() {
			defer wg.Done()

			res, err := testMoveObject(c, game2Scope, moveCommandConstructor2, movingWithFuelMock)

			require.NoError(t, err)
			moveCommand, ok := res.(*move2.MoveWithFuelCommand)
			require.True(t, ok)
			err = moveCommand.Execute()
			require.NoError(t, err)
		}()

		wg.Wait()
	})
}

func testMoveObject(c *Container, scope context.Context, f Constructor, obj any) (any, error) {
	register, err := Resolve(scope, "IoC.Register", scope, moveCommandKey, f)
	if err != nil {
		return nil, err
	}
	registerTyped, ok := register.(*RegisterCommand)
	if !ok {
		return nil, errors.New("register is not a RegisterCommand")
	}
	if err = registerTyped.Execute(context.Background()); err != nil {
		return nil, err
	}

	return Resolve(scope, moveCommandKey, obj)
}
