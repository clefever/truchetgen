package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	f, err := os.OpenFile("truchet.svg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 300
	height := 240
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:#D9D9D9")
	Tile2(canvas, 0, 0, 240, false)
	Tile1(canvas, 150, 30, 120, true)
	Tile2(canvas, 150, 90, 120, true)
	canvas.End()
}

func samplePallet(canvas *svg.SVG, x, y, size int) {
	Tile1(canvas, 0, 0, size, false)
	Tile2(canvas, size, 0, size, false)
	Tile3(canvas, 2*size, 0, size)
	Tile4(canvas, 3*size, 0, size)
	Tile5(canvas, 4*size, 0, size)

	Tile6(canvas, 0, size, size)
	Tile7(canvas, size, size, size)
	Tile8(canvas, 2*size, size, size)
	Tile9(canvas, 3*size, size, size)
	Tile10(canvas, 4*size, size, size)

	Tile11(canvas, 0, 2*size, size)
	Tile12(canvas, size, 2*size, size)
	Tile13(canvas, 2*size, 2*size, size)
	Tile14(canvas, 3*size, 2*size, size)
	Tile15(canvas, 4*size, 2*size, size)
}
