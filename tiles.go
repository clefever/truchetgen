package truchetgen

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

// TODO: Refactor this whole file

type TileSize int

const (
	Small  TileSize = 60
	Medium TileSize = 120
	Large  TileSize = 240
)

type Tile struct {
	X    int
	Y    int
	Size TileSize
}

func DrawBase(canvas *svg.SVG, x int, y int, size int, invert bool) {
	var color string
	if invert {
		color = "black"
	} else {
		color = "white"
	}
	style := fmt.Sprintf("fill:%s", color)
	canvas.Circle(x+size/4, y+size/4, size/6, style)
	canvas.Circle(x+size-size/4, y+size/4, size/6, style)
	canvas.Circle(x+size/4, y+size-size/4, size/6, style)
	canvas.Circle(x+size-size/4, y+size-size/4, size/6, style)
	canvas.Circle(x+size/2, y+size/2, size/4, style)
}

func DrawDots(canvas *svg.SVG, x int, y int, size int, invert bool) {
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s", color)
	canvas.Circle(x+size/2, y+size/4, size/12, style)
	canvas.Circle(x+size/4, y+size/2, size/12, style)
	canvas.Circle(x+size-size/4, y+size/2, size/12, style)
	canvas.Circle(x+size/2, y+size-size/4, size/12, style)
}

func Tile1(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, invert)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, false, x+size-size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, false, x+size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, invert)
}

func Tile2(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, invert)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, true, x+size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, true, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, invert)
}

func Tile3(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Line(x+size/4, y+size/2, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile4(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Line(x+size/2, y+size/4, x+size/2, y+size-size/4, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile5(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	DrawDots(canvas, x, y, size, false)
}

func Tile6(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, false, x+size-size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, false, x+size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, true, x+size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, true, x+size-size/4, y+size/2, style)
	canvas.Circle(x+size/2, y+size/2, size/12, "fill:black")
	DrawDots(canvas, x, y, size, false)
}

func Tile7(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Line(x+size/4, y+size/2, x+size-size/4, y+size/2, style)
	canvas.Line(x+size/2, y+size/4, x+size/2, y+size-size/4, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile8(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, false, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile9(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, false, x+size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile10(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, true, x+size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile11(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, true, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile12(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, false, x+size-size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, true, x+size/4, y+size/2, style)
	canvas.Line(x+size/4, y+size/2, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile13(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, false, x+size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, true, x+size-size/4, y+size/2, style)
	canvas.Line(x+size/4, y+size/2, x+size-size/4, y+size/2, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile14(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, false, x+size-size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, true, x+size-size/4, y+size/2, style)
	canvas.Line(x+size/2, y+size/4, x+size/2, y+size-size/4, style)
	DrawDots(canvas, x, y, size, false)
}

func Tile15(canvas *svg.SVG, x int, y int, size int, invert bool) {
	DrawBase(canvas, x, y, size, false)
	var color string
	if invert {
		color = "white"
	} else {
		color = "black"
	}
	style := fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%d", "none", color, size/6)
	canvas.Arc(x+size/2, y+size/4, size/4, size/4, 0, false, true, x+size/4, y+size/2, style)
	canvas.Arc(x+size/2, y+size-size/4, size/4, size/4, 0, false, false, x+size/4, y+size/2, style)
	canvas.Line(x+size/2, y+size/4, x+size/2, y+size-size/4, style)
	DrawDots(canvas, x, y, size, false)
}
