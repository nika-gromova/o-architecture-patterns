package shunting_yard

import (
	"reflect"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/stretchr/testify/require"
)

func TestToNode(t *testing.T) {
	tests := []struct {
		name    string
		postfix []token
		want    *models.ParsingNode
		wantErr bool
	}{
		{
			name: "ok",
			postfix: []token{
				{
					value: "2",
				},
				{
					value: "3",
				},
				{
					value: "5",
				},
				{
					value:      "+",
					isOperator: true,
				},
				{
					value:      "*",
					isOperator: true,
				},
			},
			wantErr: false,
			want: &models.ParsingNode{
				Value:      "*",
				IsOperator: true,
				Right: &models.ParsingNode{
					Value:      "+",
					IsOperator: true,
					Left: &models.ParsingNode{
						Value: "3",
					}, Right: &models.ParsingNode{
						Value: "5",
					},
				},
				Left: &models.ParsingNode{
					Value: "2",
				},
			},
		},
		{
			name: "should return error if postfix is invalid",
			postfix: []token{
				{
					value:      "+",
					isOperator: true,
				},
				{
					value: "2",
				},
				{
					value: "3",
				},
				{
					value: "5",
				},
				{
					value:      "+",
					isOperator: true,
				},
				{
					value:      "*",
					isOperator: true,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error if postfix is invalid",
			postfix: []token{
				{
					value: "2",
				},
				{
					value:      "+",
					isOperator: true,
				},
				{
					value: "3",
				},
				{
					value: "5",
				},
				{
					value:      "+",
					isOperator: true,
				},
				{
					value:      "*",
					isOperator: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := toNode(tt.postfix)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestShuntingYardStrategy_Parse(t *testing.T) {
	type args struct {
		data *models.ParsingData
	}
	tests := []struct {
		name    string
		args    args
		want    *models.ParsingNode
		wantErr bool
	}{
		{
			name: "should return error if failed to parse",
			args: args{
				data: &models.ParsingData{
					OperandsPriorities: map[string]int{
						models.EqualOperator: 100,
						"(":                  200,
						")":                  200,
					},
					Tokens: []string{"test", "(", "test"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error if failed to convert to node",
			args: args{
				data: &models.ParsingData{
					OperandsPriorities: map[string]int{},
					Tokens:             []string{},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return no error",
			args: args{
				data: &models.ParsingData{
					OperandsPriorities: map[string]int{
						models.EqualOperator: 100,
						"(":                  200,
						")":                  200,
					},
					Tokens: []string{"Locale", "=", "ru"},
				},
			},
			want: &models.ParsingNode{
				Value:      models.EqualOperator,
				IsOperator: true,
				Left: &models.ParsingNode{
					Value: "Locale",
				},
				Right: &models.ParsingNode{
					Value: "ru",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			got, err := s.Parse(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		priorities map[string]int
		tokens     []string
	}
	tests := []struct {
		name    string
		args    args
		want    []token
		wantErr bool
	}{
		{
			name: "should return error if ( not found",
			args: args{
				priorities: map[string]int{
					models.EqualOperator: 100,
					models.OrOperator:    50,
					models.ANDOperator:   70,
					"(":                  200,
					")":                  200,
				},
				tokens: []string{"test1", "=", "test2", "AND", ")", "token1"},
			},
			wantErr: true,
		},
		{
			name: "should parse correctly",
			args: args{
				priorities: map[string]int{
					models.EqualOperator: 100,
					models.OrOperator:    50,
					models.ANDOperator:   70,
					"(":                  200,
					")":                  200,
				},
				tokens: []string{"test1", "=", "test2", "AND", "(", "token1", "=", "token2", "OR", "tmp1", "=", "tmp2", ")"},
			},
			wantErr: false,
			want: []token{
				{
					value: "test1",
				},
				{
					value: "test2",
				},
				{
					value:      "=",
					isOperator: true,
				},
				{
					value: "token1",
				},
				{
					value: "token2",
				},
				{
					value:      "=",
					isOperator: true,
				},
				{
					value: "tmp1",
				},
				{
					value: "tmp2",
				},
				{
					value:      "=",
					isOperator: true,
				},
				{
					value:      "OR",
					isOperator: true,
				},
				{
					value:      "AND",
					isOperator: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.priorities, tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
