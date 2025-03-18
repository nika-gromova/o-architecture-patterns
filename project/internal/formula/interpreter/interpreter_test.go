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
		{
			name: "should return error if failed to interpret left expression for equal",
			expression: &EqualExpression[testType]{
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
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret right expression for equal",
			expression: &EqualExpression[testType]{
				Left: Variable[testType]{
					Name: "test-A",
				},
				Right: Variable[testType]{
					Name: "test-ABC",
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret left expression for less",
			expression: &LessExpression[testType]{
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
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret right expression for less",
			expression: &LessExpression[testType]{
				Left: Variable[testType]{
					Name: "test-A",
				},
				Right: Variable[testType]{
					Name: "test-ABC",
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret left expression for OR",
			expression: &OrExpression[testType]{
				Left: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-A",
					},
					Right: Variable[testType]{
						Name: "test-ABC",
					},
				},
				Right: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-A",
					},
					Right: Variable[testType]{
						Name: "test-ABC",
					},
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret right expression for OR",
			expression: &OrExpression[testType]{
				Left: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-A",
					},
					Right: Variable[testType]{
						Name: "test-B",
					},
				},
				Right: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-D",
					},
					Right: &Const[testType]{
						Value: testType(3),
					},
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
					"test-B": 2,
					"test-C": 1,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret left expression for AND",
			expression: &AndExpression[testType]{
				Left: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-A",
					},
					Right: Variable[testType]{
						Name: "test-ABC",
					},
				},
				Right: &NilExpression[testType]{},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return error if failed to interpret right expression for AND",
			expression: &AndExpression[testType]{
				Left: &NilExpression[testType]{},
				Right: &GraterExpression[testType]{
					Left: Variable[testType]{
						Name: "test-D",
					},
					Right: &Const[testType]{
						Value: testType(3),
					},
				},
			},
			context: &testContext{
				values: map[string]testType{
					"test-A": 1,
					"test-B": 2,
					"test-C": 1,
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

type testContextInvalid struct {
	values map[string]int
}

func (tc *testContextInvalid) GetValue(name string) (any, error) {
	v, ok := tc.values[name]
	if !ok {
		return v, fmt.Errorf("not found %s", name)
	}
	return v, nil
}

func TestVariable_InterpretComparable(t *testing.T) {
	t.Run("should return error if failed to convert value", func(t *testing.T) {
		context := &testContextInvalid{
			values: map[string]int{
				"test-A": 1,
			},
		}
		variable := &Variable[any]{
			Name: "test-A",
		}
		_, err := variable.InterpretComparable(context)

		require.Error(t, err)
	})
	t.Run("should return no error and valid value", func(t *testing.T) {
		context := &testContext{
			values: map[string]testType{
				"test-A": 1,
			},
		}
		variable := &Variable[testType]{
			Name: "test-A",
		}
		value, err := variable.InterpretComparable(context)

		require.NoError(t, err)
		require.Equal(t, testType(1), value)
	})
}
