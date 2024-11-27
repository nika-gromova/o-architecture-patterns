package change_velocity

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"go.uber.org/mock/gomock"
)

func TestChangeVelocityCommand_Execute(t *testing.T) {
	type fields struct {
		Obj   func(ctrl *gomock.Controller) RotatingAndMovingObject
		Angle base.Angle
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "у объекта нет вектора скорости",
			fields: fields{
				Obj: func(ctrl *gomock.Controller) RotatingAndMovingObject {
					mock := mocks.NewMockRotatingAndMovingObject(ctrl)
					mock.EXPECT().GetVelocityVector().Return(base.Vector{}, false)
					return mock
				},
			},
			wantErr: false,
		},
		{
			name: "объекту нельзя задать вектор скорости",
			fields: fields{
				Obj: func(ctrl *gomock.Controller) RotatingAndMovingObject {
					mock := mocks.NewMockRotatingAndMovingObject(ctrl)
					mock.EXPECT().GetVelocityVector().Return(base.Vector{
						Coordinates: []int{1, 2},
					}, true)
					mock.EXPECT().SetVelocityVector(gomock.Any()).Return(false)
					return mock
				},
				Angle: base.Angle{
					Direction:  20,
					TotalCount: 360,
				},
			},
			wantErr: true,
		},
		{
			name: "вектор скорости изменен",
			fields: fields{
				Obj: func(ctrl *gomock.Controller) RotatingAndMovingObject {
					mock := mocks.NewMockRotatingAndMovingObject(ctrl)
					mock.EXPECT().GetVelocityVector().Return(base.Vector{
						Coordinates: []int{1, 2},
					}, true)
					mock.EXPECT().SetVelocityVector(base.Vector{
						Coordinates: []int{2, -1},
					}).Return(true)
					return mock
				},
				Angle: base.Angle{
					Direction:  90,
					TotalCount: 360,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			obj := tt.fields.Obj(ctrl)
			c := ChangeVelocityCommand{
				Obj:   obj,
				Angle: tt.fields.Angle,
			}
			if err := c.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
