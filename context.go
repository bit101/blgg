package blgg

import (
	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/random"
	"github.com/fogleman/gg"
)

type Context struct {
	gg.Context
}

func NewContext(w, h int) *Context {
	context := &Context{*gg.NewContext(w, h)}
	return context
}

////////////////////
// CLEAR AND SET
////////////////////
func (c *Context) BlackOnWhite() {
	c.ClearWhite()
	c.SetBlack()
}

func (c *Context) WhiteOnBlack() {
	c.ClearBlack()
	c.SetWhite()
}

////////////////////
// CLEAR
////////////////////
func (c *Context) ClearBlack() {
	c.ClearRGB(0, 0, 0)
}

func (c *Context) ClearColor(color blcolor.Color) {
	c.ClearRGB(color.R, color.G, color.B)
}

func (c *Context) ClearGray(g float64) {
	c.ClearRGB(g, g, g)
}

func (c *Context) ClearHSV(h, s, v float64) {
	c.ClearColor(blcolor.HSV(h, s, v))
}

func (c *Context) ClearHSVA(h, s, v, a float64) {
	c.ClearColor(blcolor.HSVA(h, s, v, a))
}

func (c *Context) ClearRandomGray() {
	c.ClearGray(random.Float())
}

func (c *Context) ClearRandomRGB() {
	c.ClearRGB(random.Float(), random.Float(), random.Float())
}

func (c *Context) ClearRGB(r, g, b float64) {
	c.ClearRGBA(r, g, b, 1.0)
}

func (c *Context) ClearRGBA(r, g, b, a float64) {
	c.Push()
	c.SetRGBA(r, g, b, a)
	c.Clear()
	c.Pop()
}

func (c *Context) ClearWhite() {
	c.ClearRGB(1, 1, 1)
}

////////////////////
// SET
////////////////////
func (c *Context) SetBlack() {
	c.SetRGB(0, 0, 0)
}

func (c *Context) SetColor(color blcolor.Color) {
	c.SetRGB(color.R, color.G, color.B)
}

func (c *Context) SetGray(g float64) {
	c.SetRGB(g, g, g)
}

func (c *Context) SetHSV(h, s, v float64) {
	c.SetColor(blcolor.HSV(h, s, v))
}

func (c *Context) SetHSVA(h, s, v, a float64) {
	c.SetColor(blcolor.HSVA(h, s, v, a))
}

func (c *Context) SetRandomGray() {
	c.SetGray(random.Float())
}

func (c *Context) SetRandomRGB() {
	c.SetRGB(random.Float(), random.Float(), random.Float())
}

func (c *Context) SetWhite() {
	c.SetRGB(1, 1, 1)
}
