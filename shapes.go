// Package blgg is the main package for this module.
package blgg

import (
	"math"

	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
)

////////////////////
// ARC
////////////////////

// FillArc draws an arc and fills it.
func (c *Context) FillArc(x, y, r, a1, a2 float64) {
	c.DrawArc(x, y, r, a1, a2)
	c.Fill()
}

// StrokeArc draws an arc and strokes it.
func (c *Context) StrokeArc(x, y, r, a1, a2 float64) {
	c.DrawArc(x, y, r, a1, a2)
	c.Stroke()
}

////////////////////
// ARROW
////////////////////

// DrawArrow draws and arrow shape.
func (c *Context) DrawArrow(x0, y0, x1, y1, pointSize float64) {
	angle := math.Atan2(y1-y0, x1-x0)
	length := math.Hypot(x1-x0, y1-y0)
	c.Push()
	c.Translate(x0, y0)
	c.Rotate(angle)
	c.MoveTo(0, 0)
	c.LineTo(length, 0)
	c.LineTo(length-pointSize, -pointSize*0.6)
	c.MoveTo(length, 0)
	c.LineTo(length-pointSize, pointSize*0.6)
	c.Pop()
}

// StrokeArrow draws and arrow and strokes it.
func (c *Context) StrokeArrow(x0, y0, x1, y1, pointSize float64) {
	c.DrawArrow(x0, y0, x1, y1, pointSize)
	c.Stroke()
}

// DrawDoubleArrow draws an arrow shape with points at both ends.
func (c *Context) DrawDoubleArrow(x0, y0, x1, y1, pointSize float64) {
	angle := math.Atan2(y1-y0, x1-x0)
	length := math.Hypot(x1-x0, y1-y0)
	c.Push()
	c.Translate(x0, y0)
	c.Rotate(angle)
	c.MoveTo(0, 0)
	c.LineTo(length, 0)
	c.LineTo(length-pointSize, -pointSize*0.6)
	c.MoveTo(length, 0)
	c.LineTo(length-pointSize, pointSize*0.6)
	c.MoveTo(pointSize, pointSize*0.6)
	c.LineTo(0, 0)
	c.LineTo(pointSize, -pointSize*0.6)
	c.Pop()
}

// StrokeDoubleArrow draws an arrow shape with points at both ends and strokes it.
func (c *Context) StrokeDoubleArrow(x0, y0, x1, y1, pointSize float64) {
	c.DrawDoubleArrow(x0, y0, x1, y1, pointSize)
	c.Stroke()
}

////////////////////
// CIRCLE
////////////////////

// FillCircle draws a circle and fills it.
func (c *Context) FillCircle(x, y, r float64) {
	c.DrawCircle(x, y, r)
	c.Fill()
}

// StrokeCircle draws a circle and strokes it.
func (c *Context) StrokeCircle(x, y, r float64) {
	c.DrawCircle(x, y, r)
	c.Stroke()
}

////////////////////
// ELLIPSE
////////////////////

// FillEllipse draws an ellipse and fills it.
func (c *Context) FillEllipse(x, y, rx, ry float64) {
	c.DrawEllipse(x, y, rx, ry)
	c.Fill()
}

// StrokeEllipse draws an ellipse and strokes it.
func (c *Context) StrokeEllipse(x, y, rx, ry float64) {
	c.DrawEllipse(x, y, rx, ry)
	c.Stroke()
}

////////////////////
// ELLIPTICAL ARC
////////////////////

// FillEllipticalArc draws an elliptical arc and fills it.
func (c *Context) FillEllipticalArc(x, y, rx, ry, a1, a2 float64) {
	c.DrawEllipticalArc(x, y, rx, ry, a1, a2)
	c.Fill()
}

// StrokeEllipticalArc draws an elliptical arc and strokes it.
func (c *Context) StrokeEllipticalArc(x, y, rx, ry, a1, a2 float64) {
	c.DrawEllipticalArc(x, y, rx, ry, a1, a2)
	c.Stroke()
}

////////////////////
// FRACTAL LINE
////////////////////

// FractalLine draws a rough, fractal line between two points.
func (c *Context) FractalLine(x1, y1, x2, y2, roughness float64, iterations int) {
	dx := x2 - x1
	dy := y2 - y1
	offset := math.Sqrt(dx*dx+dy*dy) * 0.15

	var path []*geom.Point
	path = append(path, geom.NewPoint(x1, y1))
	path = append(path, geom.NewPoint(x2, y2))

	for i := 0; i < iterations; i++ {
		var newPath []*geom.Point
		for j, point := range path {
			newPath = append(newPath, geom.NewPoint(point.X, point.Y))
			if j < len(path)-1 {
				mid := geom.MidPoint(point, path[j+1])
				mid.X += random.FloatRange(-offset, offset)
				mid.Y += random.FloatRange(-offset, offset)
				newPath = append(newPath, mid)
			}
		}
		offset *= roughness
		path = newPath
	}
	c.Path(path)
}

// StrokeFractalLine draws a fractal line between two points and strokes it.
func (c *Context) StrokeFractalLine(x1, y1, x2, y2, roughness float64, iterations int) {
	c.FractalLine(x1, y1, x2, y2, roughness, iterations)
	c.Stroke()
}

////////////////////
// GRID
////////////////////

// Grid draws a grid on the context (stroked).
func (c *Context) Grid(x, y, w, h, xres, yres float64) {
	xx := x
	yy := y
	for xx <= x+w {
		c.MoveTo(xx, y)
		c.LineTo(xx, y+h)
		xx += xres
	}
	for yy <= y+h {
		c.MoveTo(x, yy)
		c.LineTo(x+w, yy)
		yy += yres
	}
	c.Stroke()
}

////////////////////
// GUIDES
////////////////////

// DrawAxes draws a vertical and horizontal line for each axis at the origin.
func (c *Context) DrawAxes() {
	w, h := c.Size()
	c.MoveTo(0, -h)
	c.LineTo(0, h)
	c.MoveTo(-w, 0)
	c.LineTo(w, 0)
	c.Stroke()
}

// DrawOrigin draws a "plus" mark at the origin, with the given size.
func (c *Context) DrawOrigin(size float64) {
	c.MoveTo(0, -size/2)
	c.LineTo(0, size/2)
	c.MoveTo(-size/2, 0)
	c.LineTo(size/2, 0)
	c.Stroke()
}

////////////////////
// HEART
////////////////////

// Heart draws a heart shape.
func (c *Context) Heart(x, y, w, h, r float64) {
	c.Push()
	c.Translate(x, y)
	c.Rotate(r)
	var path []*geom.Point
	res := math.Sqrt(w * h)
	for i := 0.0; i < res; i++ {
		a := math.Pi * 2 * i / res
		x := w * math.Pow(math.Sin(a), 3.0)
		y := h * (0.8125*math.Cos(a) - 0.3125*math.Cos(2.0*a) - 0.125*math.Cos(3.0*a) - 0.0625*math.Cos(4.0*a))
		path = append(path, geom.NewPoint(x, -y))
	}
	c.Path(path)
	c.Pop()
}

// FillHeart draws a heart shape and fills it.
func (c *Context) FillHeart(x, y, w, h, r float64) {
	c.Heart(x, y, w, h, r)
	c.Fill()
}

// StrokeHeart draws a heart shape and strokes it.
func (c *Context) StrokeHeart(x, y, w, h, r float64) {
	c.Heart(x, y, w, h, r)
	c.Stroke()
}

////////////////////
// HEX GRID
////////////////////

// HexGrid draws a hexagonal grid.
func (c *Context) HexGrid(x, y, w, h, res0, res1 float64) {
	sin60r := math.Sin(math.Pi/3.0) * res0
	xInc := 2.0 * sin60r
	yInc := res0 * 1.5
	offset := 0.0

	for yy := y; yy < y+h+yInc; yy += yInc {
		for xx := x; xx < x+w+xInc; xx += xInc {
			c.DrawRegularPolygon(6, xx+offset, yy, res1, math.Pi/2)
		}
		if offset == 0 {
			offset = sin60r
		} else {
			offset = 0
		}
	}
}

// FillHexGrid draws a hex grid and fills it.
func (c *Context) FillHexGrid(x, y, w, h, res0, res1 float64) {
	c.Push()
	c.DrawRectangle(x, y, w, h)
	c.Clip()
	c.HexGrid(x, y, w, h, res0, res1)
	c.Fill()
	c.ResetClip()
	c.Pop()
}

// StrokeHexGrid draws a hex grid and strokes it.
func (c *Context) StrokeHexGrid(x, y, w, h, res0, res1 float64) {
	c.Push()
	c.DrawRectangle(x, y, w, h)
	c.Clip()
	c.HexGrid(x, y, w, h, res0, res1)
	c.Stroke()
	c.ResetClip()
	c.Pop()
}

////////////////////
// LINE
////////////////////

// StrokeLine strokes a line between two points.
func (c *Context) StrokeLine(x0, y0, x1, y1 float64) {
	c.DrawLine(x0, y0, x1, y1)
	c.Stroke()
}

// LineThrough strokes a line that extends a certain distance beyond two points.
func (c *Context) LineThrough(x0, y0, x1, y1, overlap float64) {
	c.Push()
	c.Translate(x0, y0)
	c.Rotate(math.Atan2(y1-y0, x1-x0))
	p2 := math.Hypot(x0-x1, y0-y1)

	c.MoveTo(-overlap, 0)
	c.LineTo(p2+overlap, 0)
	c.Stroke()
	c.Pop()
}

////////////////////
// MULTI CURVE
////////////////////

// MultiCurve draws a piecewise bezier curve through a series of points.
func (c *Context) MultiCurve(points []*geom.Point) {
	c.MoveTo(points[0].X, points[0].Y)
	mid := geom.MidPoint(points[0], points[1])
	c.LineTo(mid.X, mid.Y)
	i := 1
	for i < len(points)-1 {
		p0 := points[i]
		p1 := points[i+1]
		mid := geom.MidPoint(p0, p1)
		c.QuadraticTo(p0.X, p0.Y, mid.X, mid.Y)
		i++

	}
	p := points[len(points)-1]
	c.LineTo(p.X, p.Y)
}

// StrokeMultiCurve draws a multi curve and strokes it.
func (c *Context) StrokeMultiCurve(points []*geom.Point) {
	c.MultiCurve(points)
	c.Stroke()
}

////////////////////
// MULTI LOOP
////////////////////

// MultiLoop draws a closed piecewise bezier curve through a series of points.
func (c *Context) MultiLoop(points []*geom.Point) {
	pA := points[0]
	pZ := points[len(points)-1]
	mid1 := geom.MidPoint(pZ, pA)
	c.MoveTo(mid1.X, mid1.Y)
	for i := 0; i < len(points)-1; i++ {
		p0 := points[i]
		p1 := points[i+1]
		mid := geom.MidPoint(p0, p1)
		c.QuadraticTo(p0.X, p0.Y, mid.X, mid.Y)
	}
	c.QuadraticTo(pZ.X, pZ.Y, mid1.X, mid1.Y)
}

// FillMultiLoop draws a filled, smooth, closed curve between a set of points.
func (c *Context) FillMultiLoop(points []*geom.Point) {
	c.MultiLoop(points)
	c.Fill()
}

// StrokeMultiLoop draws a stroked, smooth, closed curve between a set of points.
func (c *Context) StrokeMultiLoop(points []*geom.Point) {
	c.MultiLoop(points)
	c.Stroke()
}

////////////////////
// PATH
////////////////////

// Path draws a series of lines through a set of points.
func (c *Context) Path(points []*geom.Point) {
	for _, point := range points {
		c.LineTo(point.X, point.Y)
	}
}

// FillPath draws a path and fills it.
func (c *Context) FillPath(points []*geom.Point) {
	c.Path(points)
	c.Fill()
}

// StrokePath draws a path and strokes it.
func (c *Context) StrokePath(points []*geom.Point, close bool) {
	c.Path(points)
	if close {
		c.ClosePath()
	}
	c.Stroke()
}

////////////////////
// POINT
////////////////////

// StrokePoint draws a circle at a point and strokes it.
func (c *Context) StrokePoint(x, y, r float64) {
	c.DrawPoint(x, y, r)
	c.Stroke()
}

// FillPoint  draws a cricle at a point and fills it.
func (c *Context) FillPoint(x, y, r float64) {
	c.DrawPoint(x, y, r)
	c.Fill()
}

// Points draws and fills a circle at each point in a list of points.
func (c *Context) Points(points []*geom.Point, radius float64) {
	for _, point := range points {
		c.FillPoint(point.X, point.Y, radius)
	}
}

////////////////////
// RAY
////////////////////

// Ray draws a line segment extending from a point at an angle.
func (c *Context) Ray(x, y, angle, offset, length float64) {
	c.Push()
	c.Translate(x, y)
	c.Rotate(angle)
	c.MoveTo(offset, 0)
	c.LineTo(offset+length, 0)
	c.Stroke()
	c.Pop()
}

////////////////////
// RECTANGLE
////////////////////

// FillRectangle draws a rectangle and fills it.
func (c *Context) FillRectangle(x, y, w, h float64) {
	c.DrawRectangle(x, y, w, h)
	c.Fill()
}

// StrokeRectangle draws a rectangle and strokes it.
func (c *Context) StrokeRectangle(x, y, w, h float64) {
	c.DrawRectangle(x, y, w, h)
	c.Stroke()
}

////////////////////
// REGULAR POLYGON
////////////////////

// FillRegularPolygon draws a regular polygon and fills it.
func (c *Context) FillRegularPolygon(n int, x, y, r, rot float64) {
	c.DrawRegularPolygon(n, x, y, r, rot)
	c.Fill()
}

// StrokeRegularPolygon draws a regular polygon and strokes it.
func (c *Context) StrokeRegularPolygon(n int, x, y, r, rot float64) {
	c.DrawRegularPolygon(n, x, y, r, rot)
	c.Stroke()
}

////////////////////
// RIGHT TRIANGLE
////////////////////

// DrawRightTriangle draws a right triangle.
func (c *Context) DrawRightTriangle(x, y, w, h, r float64) {
	c.Push()
	c.Translate(x, y)
	c.Rotate(r)
	c.MoveTo(0, 0)
	c.LineTo(w, 0)
	c.LineTo(0, h)
	c.LineTo(0, 0)
	c.Pop()
}

// StrokeRightTriangle draws a right triangle and strokes it.
func (c *Context) StrokeRightTriangle(x, y, w, h, r float64) {
	c.DrawRightTriangle(x, y, w, h, r)
	c.Stroke()
}

// FillRightTriangle draws a right triangle and fills it.
func (c *Context) FillRightTriangle(x, y, w, h, r float64) {
	c.DrawRightTriangle(x, y, w, h, r)
	c.Fill()
}

////////////////////
// ROUNDED RECTANGLE
////////////////////

// FillRoundedRectangle draws a rounded rectangle and fills it.
func (c *Context) FillRoundedRectangle(x, y, w, h, r float64) {
	c.DrawRoundedRectangle(x, y, w, h, r)
	c.Fill()
}

// StrokeRoundedRectangle draws a rounded rectangle and strokes it.
func (c *Context) StrokeRoundedRectangle(x, y, w, h, r float64) {
	c.DrawRoundedRectangle(x, y, w, h, r)
	c.Stroke()
}

////////////////////
// STAR
////////////////////

// Star draws a star shape.
func (c *Context) Star(x, y, r0, r1 float64, points int, rotation float64) {
	c.Push()
	c.Translate(x, y)
	c.Rotate(rotation)
	for i := 0; i < points*2; i++ {
		r := r1
		if i%2 == 1 {
			r = r0
		}
		angle := math.Pi / float64(points) * float64(i)
		c.LineTo(math.Cos(angle)*r, math.Sin(angle)*r)
	}
	c.ClosePath()
	c.Pop()
}

// StrokeStar draws a star and strokes it.
func (c *Context) StrokeStar(x, y, r0, r1 float64, points int, rotation float64) {
	c.Star(x, y, r0, r1, points, rotation)
	c.Stroke()
}

// FillStar draws a star and fills it.
func (c *Context) FillStar(x, y, r0, r1 float64, points int, rotation float64) {
	c.Star(x, y, r0, r1, points, rotation)
	c.Fill()
}
