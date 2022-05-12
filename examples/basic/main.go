package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/clefever/truchetgen"
)

const fileName = "basic.svg"

func main() {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 480
	height := 240

	canvas := svg.New(f)

	canvas.Start(width, height)
	canvas.Scale(0.833)

	canvas.Rect(0, 0, 240, 240, "fill:#D9D9D9")
	truchetgen.Tile2(canvas, 0, 0, 240, false)

	canvas.Rect(264, 120, 120, 120, "fill:#D9D9D9")
	truchetgen.Tile2(canvas, 264, 120, 120, true)

	canvas.Rect(408, 180, 60, 60, "fill:#D9D9D9")
	truchetgen.Tile2(canvas, 408, 180, 60, false)

	canvas.Gend()
	canvas.End()
}
