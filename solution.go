package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type countable int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum countable) float64 {
	switch sidesNum {
	case SidesTriangle:
		return float64((math.Sqrt(3) * math.Pow(2, sideLen)) / 4)
	case SidesSquare:
		return float64(math.Pow(2, sideLen))
	case SidesCircle:
		return float64(math.Pi * math.Pow(2, sideLen))
	default:
		return float64(0)
	}
}
