package gosketch

import "math"

// Constants for math operations
const (
	PI = math.Pi
	HALF_PI = math.Pi / 2
	QUARTER_PI = math.Pi / 4
	TWO_PI = math.Pi * 2
)

// Sin calculates the sine of an angle (in degrees)
func Sin(angle float64) float64 {
	return math.Sin(angle * math.Pi / 180.0)
}

// Cos calculates the cosine of an angle (in degrees)
func Cos(angle float64) float64 {
	return math.Cos(angle * math.Pi / 180.0)
}

// Tan calculates the tangent of an angle (in degrees)
func Tan(angle float64) float64 {
	return math.Tan(angle * math.Pi / 180.0)
}

// Degrees converts an angle from radians to degrees
func Degrees(radians float64) float64 {
	return radians * 180.0 / math.Pi
}

// Radians converts an angle from degrees to radians
func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// Map remaps a number from one range to another
func Map(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}

// Constrain constrains a value between a minimum and maximum value
func Constrain(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Lerp calculates a number between two numbers at a specific increment
func Lerp(start, stop, amt float64) float64 {
	return start + (stop-start)*amt
}

// Dist calculates the distance between two points
func Dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

// Random returns a random number between min and max
func Random(min, max float64) float64 {
	return min + math.Mod(math.Abs(math.Sin(rand)), max-min)
}

// Private variable for quick randomness
var rand float64 = 9.7 