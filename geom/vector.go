package geom

import "math"

type Vector struct {
	U float64
	V float64
}

func IsCloseToOne(val float64) bool {
	return AreClose(val, 1.0)
}

func IsCloseToZero(val float64) bool {
	return AreClose(val, 0)
}

func AreClose(val0, val1 float64) bool {
	return math.Abs(val0-val1) < 1e-5
}

func NewVector(u, v float64) *Vector {
	return &Vector{u, v}
}

func NewVectorBetwen(p0, p1 *Point) *Vector {
	return p0.Sub(p1)
}

// NewVersor return a vector of unit length
func NewVersor(u, v float64) *Vector {
	return NewVector(u, v).Normalized()
}

// NewVersorBetween return a vector of unit length
func NewVersorBetween(p0, p1 *Point) *Vector {
	return NewVersorBetween(p0, p1).Normalized()
}

func (v *Vector) Add(w *Vector) *Vector {
	return NewVector(v.U+w.U, v.V+w.V)
}

func (v *Vector) Sub(w *Vector) *Vector {
	return NewVector(v.U-w.U, v.V-w.V)
}

func (v *Vector) Scaled(factor float64) *Vector {
	return NewVector(v.U*factor, v.V*factor)
}

func (v *Vector) Norm() float64 {
	return math.Sqrt(v.U*v.U + v.V*v.V)
}

func (v *Vector) IsNorm() bool {
	return IsCloseToOne(v.Norm())
}

func (v *Vector) Normalized() *Vector {
	return v.Scaled(1.0 / v.Norm())
}

func (v *Vector) Dot(w *Vector) float64 {
	return v.U*w.U + v.V*w.V
}

func (v *Vector) Project(w *Vector) float64 {
	return v.Dot(w.Normalized())
}

func (v *Vector) Cross(w *Vector) float64 {
	return v.U*w.V - v.V*w.U
}

func (v *Vector) IsParallel(w *Vector) bool {
	return IsCloseToZero(v.Cross(w))
}

func (v *Vector) IsPerpendicular(w *Vector) bool {
	return IsCloseToZero(v.Dot(w))
}

// AngleValueTo returns the absolute value between the two vectors.
func (v *Vector) AngleValueTo(w *Vector) float64 {
	dotProduct := v.Dot(w)
	normProduct := v.Norm() * w.Norm()
	return math.Acos(dotProduct / normProduct)
}

func (v *Vector) AngleTo(w *Vector) float64 {
	value := v.AngleValueTo(w)
	crossProduct := v.Cross(w)
	return math.Copysign(value, crossProduct)
}

func (v *Vector) Rotated(radians float64) *Vector {
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	return NewVector(v.U*cos-v.V*sin, v.U*sin+v.V*cos)
}

func (v *Vector) Sine() float64 {
	return v.V / v.Norm()
}

func (v *Vector) Cosine() float64 {
	return v.U / v.Norm()
}

func (v *Vector) Equals(w *Vector) bool {
	if v == w {
		return true
	}
	return AreClose(v.U, w.U) && AreClose(v.V, w.V)
}
