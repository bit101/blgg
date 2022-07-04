package render

import (
	"fmt"
	"os"

	"github.com/bit101/blgg/blgg"
)

const (
	Image = iota
	Gif
	Video
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
