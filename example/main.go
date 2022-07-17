package main

import (
	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/blgg"
	"github.com/bit101/blgg/render"
)

func main() {
	target := render.Gif

	switch target {
	case render.Image:
		render.RenderImage(800, 800, "out.png", renderFrame, 0.5)
		render.ViewImage("out.png")
		break

	case render.SpriteSheet:
		render.RenderSpriteSheet(40, 40, blcolor.White(), "out.png", 25, renderSpriteSheetFrame)
		render.ViewImage("out.png")
		break

	case render.Gif:
		render.RenderFrames(400, 400, 60, "frames", renderFrame)
		render.MakeGIF("ffmpeg", "frames", "out.gif", 30)
		render.ViewImage("out.gif")
		break

	case render.Video:
		render.RenderFrames(1280, 800, 60, "frames", renderFrame)
		render.ConvertToYoutube("frames", "out.mp4", 60)
		render.VLC("out.mp4", true)
		break
	}
}

func renderFrame(context *blgg.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Push()
	context.TranslateCenter()
	r := blmath.LerpSin(percent, 0, width/2)
	context.FillCircle(0, 0, r)
	context.Pop()
}

func renderSpriteSheetFrame(context *blgg.Context, width, height, percent float64) {
	context.SetBlack()
	r := blmath.LerpSin(percent, 2, width*0.45)
	context.FillCircle(width/2, height/2, r)
}
