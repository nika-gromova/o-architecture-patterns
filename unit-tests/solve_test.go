package unit_tests

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "нет корней",
			args: args{
				a: 2,
				b: 1,
				c: 1,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "2 корня",
			args: args{
				a: 2,
				b: 5,
				c: 3,
			},
			want:    []float64{-1, -1.5},
			wantErr: false,
		},
		{
			name: "1 корень",
			args: args{
				a: 2,
				b: 4,
				c: 2,
			},
			want:    []float64{-1, -1},
			wantErr: false,
		},
		{
			name: "1 корень, d<eps",
			args: args{
				a: 0.003,
				b: 0.004,
				c: 0.001,
			},
			want:    []float64{-0.6666, -0.6666},
			wantErr: false,
		},
		{
			name: "a равен 0",
			args: args{
				a: 0,
				b: 4,
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "a is NaN",
			args: args{
				a: math.Log(-1.0),
				b: 4,
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "b is NaN",
			args: args{
				a: 1,
				b: math.Log(-1.0),
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "c is NaN",
			args: args{
				a: 1,
				b: 4,
				c: math.Log(-1.0),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "a is +inf",
			args: args{
				a: math.Inf(1),
				b: 4,
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "a is -inf",
			args: args{
				a: math.Inf(-1),
				b: 4,
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "b is +inf",
			args: args{
				a: 1,
				b: math.Inf(1),
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "b is -inf",
			args: args{
				a: 1,
				b: math.Inf(-1),
				c: 2,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "c is +inf",
			args: args{
				a: 1,
				b: 4,
				c: math.Inf(1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "c is -inf",
			args: args{
				a: 1,
				b: 4,
				c: math.Inf(-1),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.args.a, tt.args.b, tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Float64s(got)
			sort.Float64s(tt.want)
			require.Equal(t, len(tt.want), len(got))
			for i := range tt.want {
				diff := math.Abs(got[i] - tt.want[i])
				require.LessOrEqual(t, diff, eps)
			}
		})
	}
}
