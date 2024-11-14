package rotate

import (
	"errors"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRotateCommand_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		obj func(ctrl *gomock.Controller) RotatingObject
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "корректное вращение объекта",
			fields: fields{
				obj: func(ctrl *gomock.Controller) RotatingObject {
					mock := mocks.NewMockRotatingObject(ctrl)
					mock.EXPECT().GetAngle().Return(base.Angle{Direction: 45, TotalCount: 360}, true)
					mock.EXPECT().GetAngularVelocity().Return(base.Angle{Direction: 20, TotalCount: 360}, true)
					mock.EXPECT().SetAngle(base.Angle{Direction: 65, TotalCount: 360}).Return(true)
					return mock
				},
			},
			err: nil,
		},
		{
			name: "у объекта невозможно получить угол наклона",
			fields: fields{
				obj: func(ctrl *gomock.Controller) RotatingObject {
					mock := mocks.NewMockRotatingObject(ctrl)
					mock.EXPECT().GetAngle().Return(base.Angle{}, false)
					return mock
				},
			},
			err: base.ErrGetProperty,
		},
		{
			name: "у объекта невозможно получить угловую скорость",
			fields: fields{
				obj: func(ctrl *gomock.Controller) RotatingObject {
					mock := mocks.NewMockRotatingObject(ctrl)
					mock.EXPECT().GetAngle().Return(base.Angle{Direction: 45, TotalCount: 360}, true)
					mock.EXPECT().GetAngularVelocity().Return(base.Angle{}, false)
					return mock
				},
			},
			err: base.ErrGetProperty,
		},
		{
			name: "объекту невозможно задать угол наклона",
			fields: fields{
				obj: func(ctrl *gomock.Controller) RotatingObject {
					mock := mocks.NewMockRotatingObject(ctrl)
					mock.EXPECT().GetAngle().Return(base.Angle{Direction: 45, TotalCount: 360}, true)
					mock.EXPECT().GetAngularVelocity().Return(base.Angle{Direction: 20, TotalCount: 360}, true)
					mock.EXPECT().SetAngle(base.Angle{Direction: 65, TotalCount: 360}).Return(false)
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
			c := &RotateCommand{
				Obj: obj,
			}
			err := c.Execute()
			require.True(t, errors.Is(err, tt.err))
		})
	}
}
