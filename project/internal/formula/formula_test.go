package formula

import (
	"context"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type testStorage struct {
	knownVariableTokens map[string]struct{}
}

func (ts *testStorage) IsKnownVariableToken(value string) bool {
	_, found := ts.knownVariableTokens[value]
	return found
}

func TestFormula_buildExpression(t *testing.T) {
	type fields struct {
		knownVariableTokens map[string]struct{}
		parser              func() Parser
	}
	tests := []struct {
		name    string
		fields  fields
		want    interpreter.AbstractExpression[any]
		wantErr bool
	}{
		{
			name: "should build correctly",
			fields: fields{
				knownVariableTokens: map[string]struct{}{
					"Locale": {},
				},
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(&models.ParsingNode{
						Value:      models.OrOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value:      models.EqualOperator,
							IsOperator: true,
							Left: &models.ParsingNode{
								Value: "Locale",
							},
							Right: &models.ParsingNode{
								Value: "ru",
							},
						},
						Right: &models.ParsingNode{
							Value:      models.EqualOperator,
							IsOperator: true,
							Left: &models.ParsingNode{
								Value: "es",
							},
							Right: &models.ParsingNode{
								Value: "Locale",
							},
						},
					}, nil)
					return parser
				},
			},
			want: &interpreter.OrExpression[any]{
				Left: &interpreter.EqualExpression[any]{
					Left: &interpreter.Variable[any]{
						Name: "Locale",
					},
					Right: &interpreter.Const[any]{
						Value: &types.StringType{
							Value: "ru",
						},
					},
				},
				Right: &interpreter.EqualExpression[any]{
					Left: &interpreter.Const[any]{
						Value: &types.StringType{
							Value: "es",
						},
					},
					Right: &interpreter.Variable[any]{
						Name: "Locale",
					},
				},
			},
		},
		{
			name: "should build correctly, right is empty",
			fields: fields{
				knownVariableTokens: map[string]struct{}{
					"Locale": {},
				},
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(&models.ParsingNode{
						Value:      models.OrOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value:      models.EqualOperator,
							IsOperator: true,
							Left: &models.ParsingNode{
								Value: "Locale",
							},
							Right: &models.ParsingNode{
								Value: "ru",
							},
						},
						Right: &models.ParsingNode{
							Value: "test",
						},
					}, nil)
					return parser
				},
			},
			want: &interpreter.OrExpression[any]{
				Left: &interpreter.EqualExpression[any]{
					Left: &interpreter.Variable[any]{
						Name: "Locale",
					},
					Right: &interpreter.Const[any]{
						Value: &types.StringType{
							Value: "ru",
						},
					},
				},
				Right: &interpreter.NilExpression[any]{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registrar := &IoCFormulaOperatorsRegistrar{
				next: &IoCFormulaStringVariableRegistrar{
					variableName: "Locale",
				},
			}

			ctx, err := registrar.Register(context.Background())
			require.NoError(t, err)

			f := &Formula{
				storage: &testStorage{
					knownVariableTokens: tt.fields.knownVariableTokens,
				},
				parser: tt.fields.parser(),
			}
			got, err := f.buildExpression(ctx, "test")
			if (err != nil) != tt.wantErr {
				t.Errorf("buildExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFormula_toExpressionNode(t *testing.T) {
	type fields struct {
		knownVariableTokens map[string]struct{}
	}
	type args struct {
		node *models.ParsingNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interpreter.ExpressionNode
	}{
		{
			name: "should build correctly",
			fields: fields{
				knownVariableTokens: map[string]struct{}{
					"Locale": {},
				},
			},
			args: args{
				node: &models.ParsingNode{
					Value:      models.OrOperator,
					IsOperator: true,
					Left: &models.ParsingNode{
						Value:      models.EqualOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value: "Locale",
						},
						Right: &models.ParsingNode{
							Value: "ru",
						},
					},
					Right: &models.ParsingNode{
						Value:      models.EqualOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value: "es",
						},
						Right: &models.ParsingNode{
							Value: "Locale",
						},
					},
				},
			},
			want: &interpreter.NodeOperator{
				Left: &interpreter.NodeOperator{
					Left: &interpreter.NodeLeaf{
						Value:        "Locale",
						VariableName: "Locale",
					},
					Right: &interpreter.NodeLeaf{
						Value:        "ru",
						VariableName: "Locale",
					},
					Value: models.EqualOperator,
				},
				Right: &interpreter.NodeOperator{
					Left: &interpreter.NodeLeaf{
						Value:        "es",
						VariableName: "Locale",
					},
					Right: &interpreter.NodeLeaf{
						Value:        "Locale",
						VariableName: "Locale",
					},
					Value: models.EqualOperator,
				},
				Value: models.OrOperator,
			},
		},
		{
			name: "should build correctly, right is empty",
			fields: fields{
				knownVariableTokens: map[string]struct{}{
					"Locale": {},
				},
			},
			args: args{
				node: &models.ParsingNode{
					Value:      models.OrOperator,
					IsOperator: true,
					Left: &models.ParsingNode{
						Value:      models.EqualOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value: "Locale",
						},
						Right: &models.ParsingNode{
							Value: "ru",
						},
					},
					Right: &models.ParsingNode{
						Value: "test",
					},
				},
			},
			want: &interpreter.NodeOperator{
				Left: &interpreter.NodeOperator{
					Left: &interpreter.NodeLeaf{
						Value:        "Locale",
						VariableName: "Locale",
					},
					Right: &interpreter.NodeLeaf{
						Value:        "ru",
						VariableName: "Locale",
					},
					Value: models.EqualOperator,
				},
				Right: nil,
				Value: models.OrOperator,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formula{
				storage: &testStorage{
					knownVariableTokens: tt.fields.knownVariableTokens,
				},
			}
			got := f.toExpressionNode(tt.args.node)
			require.Equal(t, tt.want, got)
		})
	}
}
