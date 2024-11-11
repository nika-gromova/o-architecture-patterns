package move

import (
	"errors"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMoveCommand_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		obj func(ctrl *gomock.Controller) MovingObject
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "корректное перемещения объекта из одной точки в другую",
			fields: fields{
				obj: func(ctrl *gomock.Controller) MovingObject {
					mock := mocks.NewMockMovingObject(ctrl)
					mock.EXPECT().GetLocation().Return(base.Vector{Coordinates: []int{12, 5}}, true)
					mock.EXPECT().GetVelocity().Return(base.Vector{Coordinates: []int{-7, 3}}, true)
					mock.EXPECT().SetLocation(base.Vector{Coordinates: []int{5, 8}}).Return(true)
					return mock
				},
			},
			err: nil,
		},
		{
			name: "у объекта невозможно получить положение в пространстве",
			fields: fields{
				obj: func(ctrl *gomock.Controller) MovingObject {
					mock := mocks.NewMockMovingObject(ctrl)
					mock.EXPECT().GetLocation().Return(base.Vector{}, false)
					return mock
				},
			},
			err: base.ErrGetProperty,
		},
		{
			name: "у объекта невозможно получить скорость",
			fields: fields{
				obj: func(ctrl *gomock.Controller) MovingObject {
					mock := mocks.NewMockMovingObject(ctrl)
					mock.EXPECT().GetLocation().Return(base.Vector{Coordinates: []int{12, 5}}, true)
					mock.EXPECT().GetVelocity().Return(base.Vector{}, false)
					return mock
				},
			},
			err: base.ErrGetProperty,
		},
		{
			name: "объекту невозможно задать положение в пространстве",
			fields: fields{
				obj: func(ctrl *gomock.Controller) MovingObject {
					mock := mocks.NewMockMovingObject(ctrl)
					mock.EXPECT().GetLocation().Return(base.Vector{Coordinates: []int{12, 5}}, true)
					mock.EXPECT().GetVelocity().Return(base.Vector{Coordinates: []int{-7, 3}}, true)
					mock.EXPECT().SetLocation(base.Vector{Coordinates: []int{5, 8}}).Return(false)
					return mock
				},
			},
			err: base.ErrSetProperty,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			obj := tt.fields.obj(ctrl)
			c := &MoveCommand{
				obj: obj,
			}
			err := c.Execute()
			require.True(t, errors.Is(err, tt.err))
		})
	}
}
