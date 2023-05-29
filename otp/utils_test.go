package otp

import "testing"

func Test_rangeIn(t *testing.T) {
	type args struct {
		low int
		hi  int
	}
	tests := []struct {
		name   string
		args   args
		wantLo int
		wantHi int
	}{
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 1e6,
				hi:  1e7 - 1,
			},
			1e6,
			1e7 - 1,
		},
		{
			"happy case",
			args{
				low: 10,
				hi:  11,
			},
			0,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeIn(tt.args.low, tt.args.hi); got < tt.wantLo || got > tt.wantHi {
				t.Errorf("rangeIn() = %v, %v <= want <=%v", got, tt.wantLo, tt.wantHi)
			}
		})
	}
}
