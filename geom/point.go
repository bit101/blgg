package geom

import (
	"math"

	"github.com/bit101/blgg/blmath"
	"github.com/bit101/blgg/random"
)

type Point struct {
	X, Y float64
}

// NewPoint creates a new 2d point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		X: x, Y: y,
	}
}

// LerpPoint linearly interpolates between two points.
func LerpPoint(t float64, p0 *Point, p1 *Point) *Point {
	return NewPoint(
		blmath.Lerp(t, p0.X, p1.X),
		blmath.Lerp(t, p0.Y, p1.Y),
	)
}

// RandomPoint returns a point within the rectangle defined in the params x, y, w, h.
func RandomPoint(x, y, w, h float64) *Point {
	return NewPoint(
		random.FloatRange(x, x+w),
		random.FloatRange(y, y+h),
	)
}

func RandomPolarPoint(x, y, r float64) *Point {
	angle := random.FloatRange(0, math.Pi*2)
	radius := random.FloatRange(0, r)
	return NewPoint(
		x+math.Cos(angle)*radius,
		y+math.Sin(angle)*radius,
	)
}

// RandomPointInTriangle returns a randomly generated point within the triangle described by the given points.
func RandomPointInTriangle(A, B, C *Point) *Point {
	s := random.Float()
	t := random.Float()
	a := 1.0 - math.Sqrt(t)
	b := (1.0 - s) * math.Sqrt(t)
	c := s * math.Sqrt(t)
	return NewPoint(a*A.X+b*B.X+c*C.X, a*A.Y+b*B.Y+c*C.Y)
}

// FromPolar creates a new point from and angle and radius.
func FromPolar(angle float64, radius float64) *Point {
	return NewPoint(math.Cos(angle)*radius, math.Sin(angle)*radius)
}

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) Add(q *Point) *Point {
	return NewPoint(p.X+q.X, p.Y+q.Y)
}

func (p *Point) Sub(q *Point) *Vector {
	return NewVector(p.X-q.X, p.Y-q.Y)
}

func (p *Point) Displaced(v *Vector, times float64) *Point {
	scaledVec := v.Scaled(times)
	return NewPoint(p.X+scaledVec.U, p.Y+scaledVec.V)
}

func (p *Point) Equals(q *Point) bool {
	if p == q {
		return true
	}
	return AreClose(p.X, q.X) && AreClose(p.Y, q.Y)
}
