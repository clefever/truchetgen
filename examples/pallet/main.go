package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/clefever/truchetgen"
)

const fileName = "pallet.svg"

func main() {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tileSize := 120
	tileGap := 16

	width := 5*tileSize + 4*tileGap
	height := 3*tileSize + 2*tileGap
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Scale(0.642)

	// First row
	canvas.Rect(0, 0, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile1(canvas, 0, 0, tileSize, false)

	canvas.Rect(tileSize+tileGap, 0, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile2(canvas, tileSize+tileGap, 0, tileSize, false)

	canvas.Rect(2*tileSize+2*tileGap, 0, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile3(canvas, 2*tileSize+2*tileGap, 0, tileSize, false)

	canvas.Rect(3*tileSize+3*tileGap, 0, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile4(canvas, 3*tileSize+3*tileGap, 0, tileSize, false)

	canvas.Rect(4*tileSize+4*tileGap, 0, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile5(canvas, 4*tileSize+4*tileGap, 0, tileSize, false)

	// Second row
	canvas.Rect(0, tileSize+tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile6(canvas, 0, tileSize+tileGap, tileSize, false)

	canvas.Rect(tileSize+tileGap, tileSize+tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile7(canvas, tileSize+tileGap, tileSize+tileGap, tileSize, false)

	canvas.Rect(2*tileSize+2*tileGap, tileSize+tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile8(canvas, 2*tileSize+2*tileGap, tileSize+tileGap, tileSize, false)

	canvas.Rect(3*tileSize+3*tileGap, tileSize+tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile9(canvas, 3*tileSize+3*tileGap, tileSize+tileGap, tileSize, false)

	canvas.Rect(4*tileSize+4*tileGap, tileSize+tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile10(canvas, 4*tileSize+4*tileGap, tileSize+tileGap, tileSize, false)

	// Third row
	canvas.Rect(0, 2*tileSize+2*tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile11(canvas, 0, 2*tileSize+2*tileGap, tileSize, false)

	canvas.Rect(tileSize+tileGap, 2*tileSize+2*tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile12(canvas, tileSize+tileGap, 2*tileSize+2*tileGap, tileSize, false)

	canvas.Rect(2*tileSize+2*tileGap, 2*tileSize+2*tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile13(canvas, 2*tileSize+2*tileGap, 2*tileSize+2*tileGap, tileSize, false)

	canvas.Rect(3*tileSize+3*tileGap, 2*tileSize+2*tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile14(canvas, 3*tileSize+3*tileGap, 2*tileSize+2*tileGap, tileSize, false)

	canvas.Rect(4*tileSize+4*tileGap, 2*tileSize+2*tileGap, tileSize, tileSize, "fill:#D9D9D9")
	truchetgen.Tile15(canvas, 4*tileSize+4*tileGap, 2*tileSize+2*tileGap, tileSize, false)

	canvas.Gend()
	canvas.End()
}
