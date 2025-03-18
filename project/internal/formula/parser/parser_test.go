package parser

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/parser/shunting_yard"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/stretchr/testify/require"
)

func TestParser_parseToTokens(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "should parse correctly",
			input: "'test' = 12 AND 'another' = '123'",
			want:  []string{"test", "=", "12", "AND", "another", "=", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New()
			got := p.parseToTokens(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParser_Parse(t *testing.T) {
	type fields struct {
		operatorsPriorities map[string]int
		parseStrategy       Strategy
		tokenSeparator      string
		tokenFrame          string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.ParsingNode
		wantErr bool
	}{
		{
			name: "should parse correctly",
			fields: fields{
				operatorsPriorities: defaultPriorities(),
				parseStrategy:       shunting_yard.New(),
				tokenSeparator:      " ",
				tokenFrame:          "`",
			},
			args: args{
				input: "`test` = 10",
			},
			want: &models.ParsingNode{
				Value:      models.EqualOperator,
				IsOperator: true,
				Left: &models.ParsingNode{
					Value: "test",
				},
				Right: &models.ParsingNode{
					Value: "10",
				},
			},
		},
		{
			name: "should return error if failed to parse",
			fields: fields{
				operatorsPriorities: defaultPriorities(),
				parseStrategy:       shunting_yard.New(),
				tokenSeparator:      " ",
				tokenFrame:          "`",
			},
			args: args{
				input: "`test`)",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(
				WithParseStrategy(tt.fields.parseStrategy),
				WithTokenSeparator(tt.fields.tokenSeparator),
				WithTokenFrame(tt.fields.tokenFrame),
				WithPriorities(tt.fields.operatorsPriorities),
			)

			got, err := p.Parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
