package macro_command

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/stubs"
	"go.uber.org/mock/gomock"
)

func TestMacroCommand_Execute(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmdMock := mocks.NewMockCommand(ctrl)
	cmdMock.EXPECT().Execute().Return(nil).Times(1)

	type fields struct {
		Commands []base.Command
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "остановка выполнения команд при получении ошибки выполнения очередной команды",
			fields: fields{
				Commands: []base.Command{
					&stubs.NoErrorCommand{},
					&stubs.ErrorCommand{},
					cmdMock,
				},
			},
			wantErr: true,
		},
		{
			name: "все команды выполнились успешно",
			fields: fields{
				Commands: []base.Command{
					&stubs.NoErrorCommand{},
					cmdMock,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &MacroCommand{
				Commands: tt.fields.Commands,
			}
			if err := c.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
