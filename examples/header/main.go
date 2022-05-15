package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/clefever/truchetgen"
)

const fileName = "header.svg"

func main() {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 540
	height := 540
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Scale(0.83)

	// Medium tiles
	truchetgen.Tile9(canvas, 60, 0, 120, false)
	truchetgen.Tile9(canvas, 180, 0, 120, false)
	truchetgen.Tile8(canvas, 300, 0, 120, false)
	truchetgen.Tile9(canvas, 420, 0, 120, false)

	truchetgen.Tile8(canvas, 0, 60, 120, false)
	truchetgen.Tile9(canvas, 120, 60, 120, false)
	truchetgen.Tile9(canvas, 240, 60, 120, false)
	truchetgen.Tile10(canvas, 360, 60, 120, false)

	truchetgen.Tile11(canvas, 60, 120, 120, false)
	truchetgen.Tile11(canvas, 180, 120, 120, false)
	truchetgen.Tile11(canvas, 300, 120, 120, false)
	truchetgen.Tile11(canvas, 420, 120, 120, false)

	truchetgen.Tile10(canvas, 0, 180, 120, false)
	truchetgen.Tile9(canvas, 120, 180, 120, false)
	truchetgen.Tile8(canvas, 240, 180, 120, false)
	truchetgen.Tile9(canvas, 360, 180, 120, false)

	// Small tiles
	truchetgen.Tile8(canvas, 15, 15, 60, true)
	truchetgen.Tile11(canvas, 45, 15, 60, true)
	truchetgen.Tile10(canvas, 15, 45, 60, true)
	truchetgen.Tile9(canvas, 45, 45, 60, true)

	truchetgen.Tile8(canvas, 135, 15, 60, true)
	truchetgen.Tile10(canvas, 165, 15, 60, true)
	truchetgen.Tile11(canvas, 135, 45, 60, true)
	truchetgen.Tile9(canvas, 165, 45, 60, true)

	truchetgen.Tile9(canvas, 255, 15, 60, true)
	truchetgen.Tile8(canvas, 285, 15, 60, true)
	truchetgen.Tile11(canvas, 255, 45, 60, true)
	truchetgen.Tile8(canvas, 285, 45, 60, true)

	truchetgen.Tile10(canvas, 375, 15, 60, true)
	truchetgen.Tile11(canvas, 405, 15, 60, true)
	truchetgen.Tile10(canvas, 375, 45, 60, true)
	truchetgen.Tile11(canvas, 405, 45, 60, true)

	// --------

	truchetgen.Tile8(canvas, 75, 75, 60, true)
	truchetgen.Tile11(canvas, 105, 75, 60, true)
	truchetgen.Tile10(canvas, 75, 105, 60, true)
	truchetgen.Tile11(canvas, 105, 105, 60, true)

	truchetgen.Tile10(canvas, 195, 75, 60, true)
	truchetgen.Tile10(canvas, 225, 75, 60, true)
	truchetgen.Tile8(canvas, 195, 105, 60, true)
	truchetgen.Tile8(canvas, 225, 105, 60, true)

	// truchetgen.Tile9(canvas, 255, 15, 60, true)
	// truchetgen.Tile8(canvas, 285, 15, 60, true)
	// truchetgen.Tile11(canvas, 255, 45, 60, true)
	// truchetgen.Tile8(canvas, 285, 45, 60, true)

	// truchetgen.Tile10(canvas, 375, 15, 60, true)
	// truchetgen.Tile11(canvas, 405, 15, 60, true)
	// truchetgen.Tile10(canvas, 375, 45, 60, true)
	// truchetgen.Tile11(canvas, 405, 45, 60, true)

	canvas.Gend()
	canvas.End()
}
