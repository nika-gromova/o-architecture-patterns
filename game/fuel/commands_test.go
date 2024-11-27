package fuel

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"go.uber.org/mock/gomock"
)

func TestCheckFuelCommand_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockFunc   func(ctrl *gomock.Controller) UsingFuelObject
		NeededFuel base.FuelInfo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "у объекта невозможно получить данные о топливе",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{}, false)
					return mock
				},
				NeededFuel: base.FuelInfo{},
			},
			wantErr: true,
		},
		{
			name: "не достаточно топлива",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 10,
					}, true)
					return mock
				},
				NeededFuel: base.FuelInfo{
					Value: 20,
				},
			},
			wantErr: true,
		},
		{
			name: "достаточно топлива, ровно столько, сколько нужно",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 10,
					}, true)
					return mock
				},
				NeededFuel: base.FuelInfo{
					Value: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "достаточно топлива, есть запас",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 20,
					}, true)
					return mock
				},
				NeededFuel: base.FuelInfo{
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
			obj := tt.fields.mockFunc(ctrl)
			c := &CheckFuelCommand{
				Obj:        obj,
				NeededFuel: tt.fields.NeededFuel,
			}
			if err := c.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBurnFuelCommand_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockFunc func(ctrl *gomock.Controller) UsingFuelObject
		fuel     base.FuelInfo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "у объекта невозможно получить данные о топливе",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{}, false)
					return mock
				},
				fuel: base.FuelInfo{},
			},
			wantErr: true,
		},
		{
			name: "объекту невозможно задать новое значение топлива",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 10,
					}, true)
					mock.EXPECT().SetFuel(base.FuelInfo{
						Value: 5,
					}).Return(false)
					return mock
				},
				fuel: base.FuelInfo{
					Value: 5,
				},
			},
			wantErr: true,
		},
		{
			name: "корректное уменьшение топлива у объекта",
			fields: fields{
				mockFunc: func(ctrl *gomock.Controller) UsingFuelObject {
					mock := mocks.NewMockUsingFuelObject(ctrl)
					mock.EXPECT().GetFuel().Return(base.FuelInfo{
						Value: 10,
					}, true)
					mock.EXPECT().SetFuel(base.FuelInfo{
						Value: 5,
					}).Return(true)
					return mock
				},
				fuel: base.FuelInfo{
					Value: 5,
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
			obj := tt.fields.mockFunc(ctrl)
			c := &BurnFuelCommand{
				Obj:    obj,
				ToBurn: tt.fields.fuel,
			}
			if err := c.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
