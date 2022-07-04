package geom

import (
	"math"

	"github.com/bit101/blgg/blmath"
	"github.com/bit101/blgg/random"
	"github.com/fogleman/gg"
)

// NewPoint creates a new 2d point
func NewPoint(x float64, y float64) *gg.Point {
	return &gg.Point{
		X: x, Y: y,
	}
}

// LerpPoint linearly interpolates between two points.
func LerpPoint(t float64, p0 *gg.Point, p1 *gg.Point) *gg.Point {
	return NewPoint(
		blmath.Lerp(t, p0.X, p1.X),
		blmath.Lerp(t, p0.Y, p1.Y),
	)
}

// RandomPoint returns a point within the rectangle defined in the params x, y, w, h.
func RandomPoint(x, y, w, h float64) *gg.Point {
	return NewPoint(
		random.FloatRange(x, x+w),
		random.FloatRange(y, y+h),
	)
}

func RandomPolarPoint(x, y, r float64) *gg.Point {
	angle := random.FloatRange(0, math.Pi*2)
	radius := random.FloatRange(0, r)
	return NewPoint(
		x+math.Cos(angle)*radius,
		y+math.Sin(angle)*radius,
	)
}

// RandomPointInTriangle returns a randomly generated point within the triangle described by the given points.
func RandomPointInTriangle(A, B, C *gg.Point) *gg.Point {
	s := random.Float()
	t := random.Float()
	a := 1.0 - math.Sqrt(t)
	b := (1.0 - s) * math.Sqrt(t)
	c := s * math.Sqrt(t)
	return NewPoint(a*A.X+b*B.X+c*C.X, a*A.Y+b*B.Y+c*C.Y)
}

// FromPolar creates a new point from and angle and radius.
func FromPolar(angle float64, radius float64) *gg.Point {
	return NewPoint(math.Cos(angle)*radius, math.Sin(angle)*radius)
}
