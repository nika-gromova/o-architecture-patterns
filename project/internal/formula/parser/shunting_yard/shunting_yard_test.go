package shunting_yard

import (
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
