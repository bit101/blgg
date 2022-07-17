package render

import (
	"fmt"
	"math"
	"os"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/blgg"
)

const (
	None = iota
	Image
	Gif
	Video
	SpriteSheet
)

type RenderFrameFunc func(*blgg.Context, float64, float64, float64)

func RenderImage(width, height float64, path string, renderFrameFunc RenderFrameFunc, percent float64) {
	context := blgg.NewContext(int(width), int(height))
	renderFrameFunc(context, width, height, percent)
	context.SavePNG(path)
}

func RenderFrames(width, height float64, numFrames int, frames string, renderFrameFunc RenderFrameFunc) {
	os.RemoveAll(frames)
	os.MkdirAll(frames, 0755)
	context := blgg.NewContext(int(width), int(height))
	for frame := 0; frame < numFrames; frame++ {
		percent := float64(frame) / float64(numFrames)
		fmt.Printf("\r%f", percent)
		renderFrameFunc(context, width, height, percent)
		context.SavePNG(fmt.Sprintf("%s/frame_%04d.png", frames, frame))
	}
	fmt.Println("\nDone!")
}

func RenderSpriteSheet(width, height float64, bg blcolor.Color, path string, numFrames int, renderFrameFunc RenderFrameFunc) {
	x := 0.0
	y := 0.0
	nf := float64(numFrames)
	size := math.Ceil(math.Sqrt(nf))
	context := blgg.NewContext(int(width*size), int(height*size))
	context.ClearColor(bg)

	for i := 0.0; i < nf; i++ {
		context.Push()
		context.Translate(x, y)
		percent := i / float64(numFrames)
		renderFrameFunc(context, width, height, percent)
		context.Pop()

		x += width
		if x >= size*width {
			x = 0.0
			y += height
		}
	}
	context.SavePNG(path)
}
