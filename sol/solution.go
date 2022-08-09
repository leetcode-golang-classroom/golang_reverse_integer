package sol

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		// positive out of range
		if res > math.MaxInt32/10 || (res == math.MaxInt32 && digit > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32 && digit < -8) {
			return 0
		}
		res = res*10 + digit
	}
	return res
}
