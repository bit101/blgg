package main

import (
	"github.com/bit101/blgg/blcolor"
	"github.com/bit101/blgg/blgg"
	"github.com/bit101/blgg/blmath"
	"github.com/bit101/blgg/render"
	"github.com/bit101/blgg/util"
)

func main() {
	target := render.SpriteSheet

	switch target {
	case render.Image:
		render.RenderImage(800, 800, "out.png", renderFrame, 0.5)
		util.ViewImage("out.png")
		break

	case render.SpriteSheet:
		render.RenderSpriteSheet(40, 40, blcolor.White(), "out.png", 25, renderSpriteSheetFrame)
		util.ViewImage("out.png")
		break

	case render.Gif:
		render.RenderFrames(400, 400, 60, "frames", renderFrame)
		util.MakeGIF("ffmpeg", "frames", "out.gif", 30)
		util.ViewImage("out.gif")
		break

	case render.Video:
		render.RenderFrames(1280, 800, 60, "frames", renderFrame)
		util.ConvertToYoutube("frames", "out.mp4", 60)
		util.VLC("out.mp4", true)
		break
	}
}

func renderFrame(context *blgg.Context, width, height, percent float64) {
	context.BlackOnWhite()
	r := blmath.LerpSin(percent, 0, width/2)
	context.FillCircle(width/2, height/2, r)
}

func renderSpriteSheetFrame(context *blgg.Context, width, height, percent float64) {
	context.SetBlack()
	r := blmath.LerpSin(percent, 2, width*0.45)
	context.FillCircle(width/2, height/2, r)
}
