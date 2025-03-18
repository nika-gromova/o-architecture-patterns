package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack_Push(t *testing.T) {
	type args[T any] struct {
		data T
	}
	type testCase[T any] struct {
		name string
		s    Stack[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "should add element correctly to empty stack",
			s:    Stack[int]{},
			args: args[int]{
				data: 100,
			},
			want: []int{100},
		},
		{
			name: "should add element correctly to nil stack",
			s: Stack[int]{
				items: nil,
			},
			args: args[int]{
				data: 100,
			},
			want: []int{100},
		},
		{
			name: "should add element correctly to non empty stack",
			s: Stack[int]{
				items: []int{100},
			},
			args: args[int]{
				data: 200,
			},
			want: []int{100, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.data)

			require.Equal(t, tt.want, tt.s.items)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase[T any] struct {
		name        string
		s           Stack[T]
		want        T
		wantInStack []T
	}
	tests := []testCase[int]{
		{
			name:        "should return zero value for empty stack",
			s:           Stack[int]{},
			want:        0,
			wantInStack: nil,
		},
		{
			name: "should return last element from stack",
			s: Stack[int]{
				items: []int{100},
			},
			want:        100,
			wantInStack: []int{},
		},
		{
			name: "should return element from top of stack",
			s: Stack[int]{
				items: []int{100, 200, 300},
			},
			want:        300,
			wantInStack: []int{100, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Pop()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantInStack, tt.s.items)
		})
	}
}

func TestStack_Top(t *testing.T) {
	type testCase[T any] struct {
		name        string
		s           Stack[T]
		want        T
		wantInStack []T
		wantErr     bool
	}
	tests := []testCase[int]{
		{
			name:    "should return error if stack is empty",
			s:       Stack[int]{},
			wantErr: true,
		},
		{
			name: "should return last element from stack",
			s: Stack[int]{
				items: []int{100},
			},
			want:        100,
			wantInStack: []int{100},
			wantErr:     false,
		},
		{
			name: "should return top element from stack",
			s: Stack[int]{
				items: []int{100, 200, 300},
			},
			want:        300,
			wantInStack: []int{100, 200, 300},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Top()
			if (err != nil) != tt.wantErr {
				t.Errorf("Top() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantInStack, tt.s.ToSlice())
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    Stack[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "should return true if stack is empty",
			s:    Stack[int]{},
			want: true,
		},
		{
			name: "should return false if stack has elements",
			s: Stack[int]{
				items: []int{100, 200, 300},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
