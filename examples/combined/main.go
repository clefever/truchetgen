package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/clefever/truchetgen"
)

const fileName = "combined.svg"

func main() {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 300
	height := 240

	canvas := svg.New(f)

	canvas.Start(width, height)
	canvas.Scale(0.833)

	canvas.Rect(0, 0, width, height, "fill:#D9D9D9")
	truchetgen.Tile2(canvas, 0, 0, 240, false)
	truchetgen.Tile1(canvas, 150, 30, 120, true)
	truchetgen.Tile2(canvas, 150, 90, 120, true)

	canvas.Gend()
	canvas.End()
}
