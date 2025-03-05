package interpreter

import (
	"fmt"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/stretchr/testify/require"
)

type testType int

func (t testType) Equals(value types.Comparable) bool {
	r, ok := value.(testType)
	if !ok {
		return false
	}
	return t == r
}

func (t testType) GreaterThan(value types.Comparable) bool {
	r, ok := value.(testType)
	if !ok {
		return false
	}
	return t > r
}

func (t testType) LessThan(value types.Comparable) bool {
	r, ok := value.(testType)
	if !ok {
		return false
	}
	return t < r
}

type testContext struct {
	values map[string]testType
}

func (tc *testContext) GetValue(name string) (testType, error) {
	v, ok := tc.values[name]
	if !ok {
		return v, fmt.Errorf("not found %s", name)
	}
	return v, nil
}

func Test_Interpret(t *testing.T) {
	type testCase[T testType] struct {
		name       string
		expression AbstractExpression[T]
		context    models.Data[T]
		want       bool
		wantErr    bool
	}
	tests := []testCase[testType]{
		{
			name: "should interpret expressions correctly, return true",
			expression: &AndExpression[testType]{
				Left: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-B",
					},
					Right: Variable[testType]{
						Name: "test-A",
					},
				},
				Right: &OrExpression[testType]{
					Left: &LessExpression[testType]{
						Left: Variable[testType]{
							Name: "test-B",
						},
						Right: Variable[testType]{
							Name: "test-A",
						},
					},
					Right: &EqualExpression[testType]{
						Left: Variable[testType]{
							Name: "test-C",
						},
						Right: &Const[testType]{
							Value: testType(3),
						},
					},
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
					"test-B": 2,
					"test-C": 3,
				},
			},
			want: true,
		},
		{
			name: "should interpret expressions correctly, return false",
			expression: &AndExpression[testType]{
				Left: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-B",
					},
					Right: Variable[testType]{
						Name: "test-A",
					},
				},
				Right: &OrExpression[testType]{
					Left: &LessExpression[testType]{
						Left: Variable[testType]{
							Name: "test-B",
						},
						Right: Variable[testType]{
							Name: "test-A",
						},
					},
					Right: &EqualExpression[testType]{
						Left: Variable[testType]{
							Name: "test-C",
						},
						Right: Variable[testType]{
							Name: "test-B",
						},
					},
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
					"test-B": 2,
					"test-C": 3,
					"test-D": 3,
				},
			},
			want: false,
		},
		{
			name: "should return error if failed to get value from context",
			expression: &GraterExpression[testType]{
				Left: Variable[testType]{
					Name: "test-ABC",
				},
				Right: Variable[testType]{
					Name: "test-A",
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
					"test-B": 2,
					"test-C": 3,
					"test-D": 3,
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.expression.Interpret(tt.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interpret() = %v, want %v", got, tt.want)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
