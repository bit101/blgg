// Package blgg is the main package for this module.
package blgg

import (
	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	"github.com/fogleman/gg"
)

// Context represents the drawing context.
type Context struct {
	gg.Context
	ClampColors bool
}

// NewContext creates a new blgg context with the given width and height.
func NewContext(w, h int) *Context {
	context := &Context{
		*gg.NewContext(w, h),
		true,
	}
	return context
}

// NewContextF creates a new blgg context with the given float64 width and height.
func NewContextF(w, h float64) *Context {
	return NewContext(int(w), int(h))
}

// TranslateCenter translates the context to the center of the image.
func (c *Context) TranslateCenter() {
	center := c.Center()
	c.Translate(center.X, center.Y)
}

// Center returns a point representing the center of the image.
func (c *Context) Center() *geom.Point {
	w, h := c.Size()
	return geom.NewPoint(w/2, h/2)
}

// Size returns the width and height of the image as float64s.
func (c *Context) Size() (float64, float64) {
	return float64(c.Width()), float64(c.Height())
}

// //////////////////
// CLEAR AND SET
// //////////////////

// BlackOnWhite clears the image to white and sets the drawing color to black.
func (c *Context) BlackOnWhite() {
	c.ClearWhite()
	c.SetBlack()
}

// WhiteOnBlack clears the image to black and sets the drawing color to white.
func (c *Context) WhiteOnBlack() {
	c.ClearBlack()
	c.SetWhite()
}

// //////////////////
// CLEAR
// //////////////////

// ClearBlack clears the image to black.
func (c *Context) ClearBlack() {
	c.ClearRGB(0, 0, 0)
}

// ClearColor clears the image to the given blcolor.
func (c *Context) ClearColor(color blcolor.Color) {
	c.ClearRGB(color.R, color.G, color.B)
}

// ClearGray clears the image to the given shade of gray.
func (c *Context) ClearGray(g float64) {
	if c.ClampColors {
		g = blmath.Clamp(g, 0, 1)
	}
	c.ClearRGB(g, g, g)
}

// ClearHSV clears the image to the given hsv value.
func (c *Context) ClearHSV(h, s, v float64) {
	c.ClearColor(blcolor.HSV(h, s, v))
}

// ClearHSVA clears the image to the given hsva value.
func (c *Context) ClearHSVA(h, s, v, a float64) {
	c.ClearColor(blcolor.HSVA(h, s, v, a))
}

// ClearRandomGray clears the image to a random shade of gray.
func (c *Context) ClearRandomGray() {
	c.ClearGray(random.Float())
}

// ClearRandomRGB clears the image to a random rgb value.
func (c *Context) ClearRandomRGB() {
	c.ClearRGB(random.Float(), random.Float(), random.Float())
}

// ClearRGB clears the image to the given rgb value.
func (c *Context) ClearRGB(r, g, b float64) {
	if c.ClampColors {
		r = blmath.Clamp(r, 0, 1)
		g = blmath.Clamp(g, 0, 1)
		b = blmath.Clamp(b, 0, 1)
	}
	c.ClearRGBA(r, g, b, 1.0)
}

// ClearRGBA clears the image to the given rgba value.
func (c *Context) ClearRGBA(r, g, b, a float64) {
	c.Push()
	c.SetRGBA(r, g, b, a)
	c.Clear()
	c.Pop()
}

// ClearWhite clears the image to white.
func (c *Context) ClearWhite() {
	c.ClearRGB(1, 1, 1)
}

// //////////////////
// SET
// //////////////////

// SetBlack sets the drawing color to black.
func (c *Context) SetBlack() {
	c.SetRGB(0, 0, 0)
}

// SetColor sets the drawing color to the given blcolor.
func (c *Context) SetColor(color blcolor.Color) {
	c.SetRGB(color.R, color.G, color.B)
}

// SetGray sets the drawing color to the given shade of gray.
func (c *Context) SetGray(g float64) {
	if c.ClampColors {
		g = blmath.Clamp(g, 0, 1)
	}
	c.SetRGB(g, g, g)
}

// SetHSV sets the drawing color to the given hsv value.
func (c *Context) SetHSV(h, s, v float64) {
	c.SetColor(blcolor.HSV(h, s, v))
}

// SetHSVA sets the drawing color to the given hsva value.
func (c *Context) SetHSVA(h, s, v, a float64) {
	c.SetColor(blcolor.HSVA(h, s, v, a))
}

// SetRandomGray sets the drawing color to a random gray shade.
func (c *Context) SetRandomGray() {
	c.SetGray(random.Float())
}

// SetRandomRGB sets the drawing color to a random rgb value.
func (c *Context) SetRandomRGB() {
	c.SetRGB(random.Float(), random.Float(), random.Float())
}

// SetRGB clears the image to the given rgb value.
func (c *Context) SetRGB(r, g, b float64) {
	if c.ClampColors {
		r = blmath.Clamp(r, 0, 1)
		g = blmath.Clamp(g, 0, 1)
		b = blmath.Clamp(b, 0, 1)
	}
	c.Context.SetRGBA(r, g, b, 1.0)
}

// SetWhite sets the drawing color to white.
func (c *Context) SetWhite() {
	c.SetRGB(1, 1, 1)
}

// SetPixelF sets the given pixel to the active drawing color, using float64 coords.
func (c *Context) SetPixelF(x, y float64) {
	c.SetPixel(int(x), int(y))
}

// ProcessPixels runs a function for every pixel in the context.
func (c *Context) ProcessPixels(pixelFunc func(context *Context, x, y float64)) {
	w, h := c.Size()
	for x := 0.0; x < w; x++ {
		for y := 0.0; y < h; y++ {
			pixelFunc(c, x, y)
		}
	}
}
