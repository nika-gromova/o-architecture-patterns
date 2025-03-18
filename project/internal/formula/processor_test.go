package formula

import (
	"context"
	"fmt"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/registrars"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/cache"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
	"github.com/nika-gromova/o-architecture-patterns/project/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProcessor_buildExpression(t *testing.T) {
	type fields struct {
		parser func() Parser
	}
	tests := []struct {
		name    string
		fields  fields
		want    interpreter.AbstractExpression[any]
		wantErr bool
	}{
		{
			name: "should return error if parser failed",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(nil, fmt.Errorf("error"))
					return parser
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error if failed to build expression node",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(&models.ParsingNode{
						Value:      "test",
						IsOperator: false,
					}, nil)
					return parser
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error if failed to build expression",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(&models.ParsingNode{
						Value:      "unknownOperator",
						IsOperator: true,
						Left: &models.ParsingNode{
							Value:      models.EqualOperator,
							IsOperator: true,
							Left: &models.ParsingNode{
								Value: "Test",
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
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error if failed to convert expression",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse(gomock.Any()).Return(&models.ParsingNode{
						Value:      "invalidOperator",
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
			want:    nil,
			wantErr: true,
		},
		{
			name: "should build correctly",
			fields: fields{
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
			chain := registrars.NewChain(
				&registrars.IoCFormulaOperatorsOrRegistrar{},
				&registrars.IoCFormulaOperatorsEqualRegistrar{},
				&registrars.IoCFormulaStringVariableRegistrar{VariableName: "Locale"},
			)

			ctx, err := chain.Register(context.Background())
			require.NoError(t, err)
			err = ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+"invalidOperator", func(args ...any) (any, error) {
				return nil, nil
			})
			require.NoError(t, err)

			f := &Processor{
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

func TestProcessor_toExpressionNode(t *testing.T) {
	type args struct {
		node *models.ParsingNode
	}
	tests := []struct {
		name string
		args args
		want interpreter.ExpressionNode
	}{
		{
			name: "should return nil if root node is no an operator",
			args: args{
				node: &models.ParsingNode{
					Value:      "Test",
					IsOperator: false,
				},
			},
			want: nil,
		},
		{
			name: "should return nil if left if nil for root node",
			args: args{
				node: &models.ParsingNode{
					Value:      models.OrOperator,
					IsOperator: true,
					Right: &models.ParsingNode{
						Value:      models.EqualOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value: "Locale",
						},
						Right: &models.ParsingNode{
							Value: "ru",
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "should return nil if right if nil for root node",
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
				},
			},
			want: nil,
		},
		{
			name: "should build correctly",
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
			f := &Processor{}

			chain := registrars.NewChain(
				&registrars.IoCFormulaOperatorsOrRegistrar{},
				&registrars.IoCFormulaOperatorsEqualRegistrar{},
				&registrars.IoCFormulaStringVariableRegistrar{VariableName: "Locale"},
			)

			ctx, err := chain.Register(context.Background())
			require.NoError(t, err)

			got := f.toExpressionNode(ctx, tt.args.node)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProcessor_toExpression(t *testing.T) {
	tests := []struct {
		name string
		arg  *models.ParsingNode
		want interpreter.ExpressionNode
	}{
		{
			name: "should build correctly",
			arg: &models.ParsingNode{
				Value:      models.EqualOperator,
				IsOperator: true,
				Left: &models.ParsingNode{
					Value:      "Locale",
					IsOperator: false,
				},
				Right: &models.ParsingNode{
					Value:      "ru",
					IsOperator: false,
				},
			},
			want: &interpreter.NodeOperator{
				Value: models.EqualOperator,
				Left: &interpreter.NodeLeaf{
					Value:        "Locale",
					VariableName: "Locale",
				},
				Right: &interpreter.NodeLeaf{
					Value:        "ru",
					VariableName: "Locale",
				},
			},
		},
		{
			name: "should return nil if unknown variable found",
			arg: &models.ParsingNode{
				Value:      models.EqualOperator,
				IsOperator: true,
				Left: &models.ParsingNode{
					Value:      "Test",
					IsOperator: false,
				},
				Right: &models.ParsingNode{
					Value:      "ru",
					IsOperator: false,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := registrars.NewChain(
				&registrars.IoCFormulaOperatorsOrRegistrar{},
				&registrars.IoCFormulaOperatorsEqualRegistrar{},
				&registrars.IoCFormulaStringVariableRegistrar{VariableName: "Locale"},
			)

			ctx, err := chain.Register(context.Background())
			require.NoError(t, err)

			f := &Processor{}
			got := f.toExpression(ctx, tt.arg.Value, tt.arg.Left, tt.arg.Right)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProcessor_Evaluate(t *testing.T) {
	type fields struct {
		parser func() Parser
		cache  func() Cache
	}
	type args struct {
		input *models.Formula
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "should use expression from cache",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					return parser
				},
				cache: func() Cache {
					c := cache.New()
					c.Set("test", &interpreter.NilExpression[any]{})
					return c
				},
			},
			args: args{
				input: &models.Formula{
					Expression: "test",
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "should build expression if not found in cache",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse("test").Return(&models.ParsingNode{
						Value:      models.OrOperator,
						IsOperator: true,
						Left: &models.ParsingNode{
							Value:      models.EqualOperator,
							IsOperator: true,
							Left: &models.ParsingNode{
								Value: "ru",
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
				cache: func() Cache {
					c := cache.New()
					return c
				},
			},
			args: args{
				input: &models.Formula{
					Expression: "test",
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "should return error if failed to build expression",
			fields: fields{
				parser: func() Parser {
					ctrl := gomock.NewController(t)
					parser := mocks.NewMockParser(ctrl)
					parser.EXPECT().Parse("test").Return(&models.ParsingNode{
						IsOperator: false,
					}, nil)
					return parser
				},
				cache: func() Cache {
					c := cache.New()
					return c
				},
			},
			args: args{
				input: &models.Formula{
					Expression: "test",
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cache()
			p := New(tt.fields.parser(), c)

			chain := registrars.NewChain(
				&registrars.IoCFormulaOperatorsOrRegistrar{},
				&registrars.IoCFormulaOperatorsEqualRegistrar{},
				&registrars.IoCFormulaStringVariableRegistrar{VariableName: "Locale"},
			)

			ctx, err := chain.Register(context.Background())
			require.NoError(t, err)

			got, err := p.Evaluate(ctx, tt.args.input, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate() got = %v, want %v", got, tt.want)
			}
			if !tt.wantErr {
				_, ok := c.Get(tt.args.input.Expression)
				require.True(t, ok)
			}
		})
	}
}
