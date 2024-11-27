package move

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"go.uber.org/mock/gomock"
)

func TestMoveWithFuelCommand_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		Obj  func(ctrl *gomock.Controller) MovingWithFuelObj
		Fuel base.FuelInfo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "не хватает топлива, движение не выполняется",
			fields: fields{
				Obj: func(ctrl *gomock.Controller) MovingWithFuelObj {
					obj := mocks.NewMockMovingWithFuelObj(ctrl)
					obj.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 10,
					}, true)
					return obj
				},
				Fuel: base.FuelInfo{
					Value: 20,
				},
			},
			wantErr: true,
		},
		{
			name: "хватает топлива, движение выполнено, топливо потрачено",
			fields: fields{
				Obj: func(ctrl *gomock.Controller) MovingWithFuelObj {
					obj := mocks.NewMockMovingWithFuelObj(ctrl)
					obj.EXPECT().GetFuel().Return(base.FuelInfo{Value: 10}, true)
					obj.EXPECT().GetLocation().Return(base.Vector{Coordinates: []int{1, 2}}, true)
					obj.EXPECT().GetVelocity().Return(base.Vector{Coordinates: []int{3, 4}}, true)
					obj.EXPECT().SetLocation(base.Vector{Coordinates: []int{4, 6}}).Return(true)
					obj.EXPECT().GetFuel().Return(base.FuelInfo{Value: 10}, true)
					obj.EXPECT().SetFuel(base.FuelInfo{Value: 0}).Return(true)
					return obj
				},
				Fuel: base.FuelInfo{
					Value: 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := &MoveWithFuelCommand{
				Obj:  tt.fields.Obj(ctrl),
				Fuel: tt.fields.Fuel,
			}
			if err := c.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
