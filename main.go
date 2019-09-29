package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Gomoku",
		Bounds: pixel.R(0, 0, 1280, 900),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	DrawCheckerboard(imd)

	for !win.Closed() {
		win.Clear(colornames.Khaki)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

// DrawCheckerboard draws the checkerboard
func DrawCheckerboard(imd *imdraw.IMDraw) {
	imd.Color = colornames.Black
	imd.EndShape = imdraw.SharpEndShape
	// imd.Push(pixel.V(400, 150), pixel.V(800, 550))

	var width float64 = 40.00
	for i := 0; i <= 15; i++ {
		imd.Push(pixel.V(400, 150+width*float64(i)), pixel.V(1000, 150+width*float64(i)))
		imd.Line(2)
	}
	for i := 0; i <= 15; i++ {
		imd.Push(pixel.V(400+width*float64(i), 150), pixel.V(400+width*float64(i), 750))
		imd.Line(2)
	}

}
