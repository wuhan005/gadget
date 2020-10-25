package gadget

import "math"

// Abs returns the Abs of a Integer.
func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
