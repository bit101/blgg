package blgg

import (
	"math"

	"github.com/bit101/bitlib/geom"
)

//////////////////
// POINT
//////////////////
func (c *Context) DrawGeomPoint(p *geom.Point, r float64) {
	c.DrawCircle(p.X, p.Y, r)
}

func (c *Context) StrokeGeomPoint(p *geom.Point, r float64) {
	c.StrokeCircle(p.X, p.Y, r)
}

func (c *Context) FillGeomPoint(p *geom.Point, r float64) {
	c.FillCircle(p.X, p.Y, r)
}

//////////////////
// VECTOR
//////////////////
func (c *Context) DrawGeomVectorAt(v *geom.Vector, p *geom.Point, arrowSize float64) {
	c.DrawArrow(p.X, p.Y, p.X+v.U, p.Y+v.V, arrowSize)
}

func (c *Context) StrokeGeomVectorAt(v *geom.Vector, p *geom.Point, arrowSize float64) {
	c.StrokeArrow(p.X, p.Y, p.X+v.U, p.Y+v.V, arrowSize)
}

//////////////////
// SEGMENT
//////////////////
func (c *Context) DrawGeomSegment(s *geom.Segment) {
	c.MoveTo(s.Start.X, s.Start.Y)
	c.LineTo(s.End.X, s.End.Y)
}
func (c *Context) StrokeGeomSegment(s *geom.Segment) {
	c.DrawGeomSegment(s)
	c.Stroke()
}

//////////////////
// LINE
//////////////////
func (c *Context) DrawGeomLine(l *geom.Line, length float64) {
	c.Push()
	c.Translate(l.Base.X, l.Base.Y)
	c.Rotate(math.Atan2(l.Direction.V, l.Direction.U))
	c.MoveTo(-length, 0)
	c.LineTo(length, 0)
	c.Pop()
}

func (c *Context) StrokeGeomLine(l *geom.Line, length float64) {
	c.DrawGeomLine(l, length)
	c.Stroke()
}

//////////////////
// POLYGON
//////////////////
func (c *Context) DrawGeomPolygon(p *geom.Polygon) {
	c.MoveTo(p.Vertices[0].X, p.Vertices[0].Y)
	for i := 1; i < len(p.Vertices); i++ {
		c.LineTo(p.Vertices[i].X, p.Vertices[i].Y)
	}
	c.LineTo(p.Vertices[0].X, p.Vertices[0].Y)
}

func (c *Context) StrokeGeomPolygon(p *geom.Polygon) {
	c.DrawGeomPolygon(p)
	c.Stroke()
}
func (c *Context) FillGeomPolygon(p *geom.Polygon) {
	c.DrawGeomPolygon(p)
	c.Fill()
}

//////////////////
// RECT
//////////////////
func (c *Context) DrawGeomRect(r *geom.Rect) {
	c.DrawRectangle(r.Origin.X, r.Origin.Y, r.Size.Width, r.Size.Height)
}

func (c *Context) StrokeGeomRect(r *geom.Rect) {
	c.DrawGeomRect(r)
	c.Stroke()
}

func (c *Context) FillGeomRect(r *geom.Rect) {
	c.DrawGeomRect(r)
	c.Fill()
}

//////////////////
// Circle
//////////////////
func (c *Context) DrawGeomCircle(circle *geom.Circle) {
	c.DrawCircle(circle.Center.X, circle.Center.Y, circle.Radius)
}

func (c *Context) StrokeGeomCircle(circle *geom.Circle) {
	c.DrawGeomCircle(circle)
	c.Stroke()
}

func (c *Context) FillGeomCircle(circle *geom.Circle) {
	c.DrawGeomCircle(circle)
	c.Fill()
}
