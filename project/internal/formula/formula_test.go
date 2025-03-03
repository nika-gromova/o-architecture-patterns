package formula

import (
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter/types"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/ioc"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/stretchr/testify/require"
)

func TestFormula_buildExpression(t *testing.T) {
	type fields struct {
		knownVariableTokens map[string]struct{}
	}
	type args struct {
		node *models.ParsingNode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interpreter.AbstractExpression[any]
		wantErr bool
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
			ctx, err := ioc.InitForFormula()
			require.NoError(t, err)

			f := &Formula{
				knownVariableTokens: tt.fields.knownVariableTokens,
			}
			got, err := f.buildExpression(ctx, tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
