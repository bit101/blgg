// Package render renders a single image or a number of frames
package render

import (
	"fmt"
	"math"
	"os"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/blgg"
)

const (
	// None will render nothing.
	None = iota
	// ImageTarget will render an image.
	ImageTarget
	// GifTarget will render an animated gif.
	GifTarget
	// VideoTarget will render a video.
	VideoTarget
	// SpriteSheetTarget will render a sprite sheet.
	SpriteSheetTarget
)

// FrameFunc is the interface for a function that renders a single frame.
type FrameFunc func(*blgg.Context, float64, float64, float64)

// Image sets up the rendering of a single image.
func Image(width, height float64, path string, frameFunc FrameFunc, percent float64) {
	context := blgg.NewContext(int(width), int(height))
	frameFunc(context, width, height, percent)
	context.SavePNG(path)
}

// Frames sets up the renderin of a series of frames.
func Frames(width, height float64, numFrames int, frames string, frameFunc FrameFunc) {
	os.RemoveAll(frames)
	os.MkdirAll(frames, 0755)
	context := blgg.NewContext(int(width), int(height))
	for frame := 0; frame < numFrames; frame++ {
		percent := float64(frame) / float64(numFrames)
		fmt.Printf("\r%f", percent)
		frameFunc(context, width, height, percent)
		context.SavePNG(fmt.Sprintf("%s/frame_%04d.png", frames, frame))
	}
	fmt.Println("\nDone!")
}

// SpriteSheet sets up the rendering of a sprite sheet.
func SpriteSheet(width, height float64, bg blcolor.Color, path string, numFrames int, frameFunc FrameFunc) {
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
		frameFunc(context, width, height, percent)
		context.Pop()

		x += width
		if x >= size*width {
			x = 0.0
			y += height
		}
	}
	context.SavePNG(path)
}
