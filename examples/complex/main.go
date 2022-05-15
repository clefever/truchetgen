package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/clefever/truchetgen"
)

const fileName = "complex.svg"

func main() {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 600
	height := 600

	canvas := svg.New(f)

	canvas.Start(width, height)
	canvas.Scale(0.35)

	// Idea for describing the pattern:
	// Size: 16x16
	// (0 ,0  2)
	// (2 ,0  2)
	// (4 ,0  2)
	// (6 ,0  2)
	// (8 ,0  2)
	// (12,0  2)
	// (0, 2  2)
	// (2, 2  2)
	// (3, 2  2)
	// (2, 3  2)
	// (3, 3  2)
	// (4, 2  2)
	// ...

	// Large tiles
	truchetgen.Tile2(canvas, 240, 0, 240, false)
	truchetgen.Tile2(canvas, 360, 0, 240, false)
	truchetgen.Tile2(canvas, 240, 120, 240, false)
	truchetgen.Tile2(canvas, 360, 120, 240, false)

	truchetgen.Tile2(canvas, 0, 240, 240, false)

	truchetgen.Tile2(canvas, 120, 360, 240, false)
	truchetgen.Tile2(canvas, 240, 360, 240, false)
	truchetgen.Tile2(canvas, 360, 360, 240, false)

	// Medium tiles
	truchetgen.Tile2(canvas, 30, 30, 120, true)
	truchetgen.Tile2(canvas, 90, 30, 120, true)
	truchetgen.Tile2(canvas, 150, 30, 120, true)
	truchetgen.Tile2(canvas, 210, 30, 120, true)

	truchetgen.Tile2(canvas, 30, 90, 120, true)
	truchetgen.Tile2(canvas, 150, 90, 120, true)

	truchetgen.Tile2(canvas, 30, 150, 120, true)
	truchetgen.Tile2(canvas, 90, 150, 120, true)

	truchetgen.Tile2(canvas, 90, 210, 120, true)
	truchetgen.Tile2(canvas, 150, 210, 120, true)
	truchetgen.Tile2(canvas, 210, 210, 120, true)

	truchetgen.Tile2(canvas, 150, 270, 120, true)
	truchetgen.Tile2(canvas, 210, 270, 120, true)
	truchetgen.Tile2(canvas, 270, 270, 120, true)
	truchetgen.Tile2(canvas, 330, 270, 120, true)
	truchetgen.Tile2(canvas, 390, 270, 120, true)
	truchetgen.Tile2(canvas, 450, 270, 120, true)

	truchetgen.Tile2(canvas, 150, 330, 120, true)
	truchetgen.Tile2(canvas, 210, 330, 120, true)

	truchetgen.Tile2(canvas, 330, 330, 120, true)
	truchetgen.Tile2(canvas, 390, 330, 120, true)

	truchetgen.Tile2(canvas, 30, 390, 120, true)
	truchetgen.Tile2(canvas, 90, 390, 120, true)
	truchetgen.Tile2(canvas, 30, 450, 120, true)
	truchetgen.Tile2(canvas, 90, 450, 120, true)

	// Small tiles
	truchetgen.Tile2(canvas, 105, 105, 60, false)
	truchetgen.Tile2(canvas, 135, 105, 60, false)
	truchetgen.Tile2(canvas, 105, 135, 60, false)
	truchetgen.Tile2(canvas, 135, 135, 60, false)

	truchetgen.Tile2(canvas, 225, 105, 60, false)
	truchetgen.Tile2(canvas, 255, 105, 60, false)
	truchetgen.Tile2(canvas, 225, 135, 60, false)
	truchetgen.Tile2(canvas, 255, 135, 60, false)

	truchetgen.Tile2(canvas, 165, 165, 60, false)
	truchetgen.Tile2(canvas, 195, 165, 60, false)
	truchetgen.Tile2(canvas, 165, 195, 60, false)
	truchetgen.Tile2(canvas, 195, 195, 60, false)

	truchetgen.Tile2(canvas, 225, 165, 60, false)
	truchetgen.Tile2(canvas, 255, 165, 60, false)
	truchetgen.Tile2(canvas, 225, 195, 60, false)
	truchetgen.Tile2(canvas, 255, 195, 60, false)

	truchetgen.Tile2(canvas, 45, 225, 60, false)
	truchetgen.Tile2(canvas, 75, 225, 60, false)
	truchetgen.Tile2(canvas, 45, 255, 60, false)
	truchetgen.Tile2(canvas, 75, 255, 60, false)

	truchetgen.Tile2(canvas, 285, 345, 60, false)
	truchetgen.Tile2(canvas, 315, 345, 60, false)
	truchetgen.Tile2(canvas, 285, 375, 60, false)
	truchetgen.Tile2(canvas, 315, 375, 60, false)

	truchetgen.Tile2(canvas, 465, 345, 60, false)
	truchetgen.Tile2(canvas, 495, 345, 60, false)
	truchetgen.Tile2(canvas, 465, 375, 60, false)
	truchetgen.Tile2(canvas, 495, 375, 60, false)

	canvas.Gend()
	canvas.End()
}
