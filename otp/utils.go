package otp

import "math/rand"

/*
	tests := []numRange{
			{0, 100},
			{100, 200},
			{3000, 10000},
			{10000000, 99999999},
		}

		for _, test := range tests {
			fmt.Printf("Num %d <= %d <= %d\n", test.low, rangeIn(test.low, test.hi), test.hi)
		}
*/
func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low) + 100
}

func genOtp(len int) int {
	return rangeIn(len, len*10-1)
}
