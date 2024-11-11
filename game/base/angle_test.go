package base

import (
	"reflect"
	"testing"
)

func TestAngle_Plus(t *testing.T) {
	type fields struct {
		Direction  int
		TotalCount int
	}
	type args struct {
		other Angle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Angle
	}{
		{
			name: "одинаковое кол-во точек",
			fields: fields{
				Direction:  10,
				TotalCount: 100,
			},
			args: args{
				other: Angle{
					Direction:  50,
					TotalCount: 100,
				},
			},
			want: Angle{
				Direction:  60,
				TotalCount: 100,
			},
		},
		{
			name: "одинаковое кол-во точек, полный круг",
			fields: fields{
				Direction:  20,
				TotalCount: 100,
			},
			args: args{
				other: Angle{
					Direction:  90,
					TotalCount: 100,
				},
			},
			want: Angle{
				Direction:  10,
				TotalCount: 100,
			},
		},
		{
			name: "разное кол-во точек",
			fields: fields{
				Direction:  10,
				TotalCount: 100,
			},
			args: args{
				other: Angle{
					Direction:  10,
					TotalCount: 360,
				},
			},
			want: Angle{
				Direction:  13,
				TotalCount: 100,
			},
		},
		{
			name: "разное кол-во точек, полный круг",
			fields: fields{
				Direction:  270,
				TotalCount: 360,
			},
			args: args{
				other: Angle{
					Direction:  40,
					TotalCount: 100,
				},
			},
			want: Angle{
				Direction:  54,
				TotalCount: 360,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Angle{
				Direction:  tt.fields.Direction,
				TotalCount: tt.fields.TotalCount,
			}
			if got := a.Plus(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}
