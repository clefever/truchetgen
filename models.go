package truchetgen

import (
	"errors"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"

	svg "github.com/ajstarks/svgo"
)

const tileMultiplier = 60

// TODO: Add invert, foreground, background
type TileLayout struct {
	Width          int
	Height         int
	DefaultType    *int
	ShowBackground bool
	Scaling        *float64
	Tiles          []Tile
}

type Tile struct {
	X    int
	Y    int
	Type *int
	Size int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (l *TileLayout) Draw(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := range l.Tiles {
		if l.Tiles[i].Type == nil {
			l.Tiles[i].Type = l.DefaultType
		}
	}

	var tiles []Tile
	tiles = append(tiles, l.Tiles...)
	if len(tiles) == 0 {
		return errors.New("no tiles to draw")
	}

	// Bigger tiles need to be drawn first
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].Size > tiles[j].Size
	})

	maxSize := tiles[0].Size
	currSize := maxSize

	width := l.Width*tileMultiplier/2 + maxSize*tileMultiplier/2
	height := l.Height*tileMultiplier/2 + maxSize*tileMultiplier/2

	canvas := svg.New(f)

	scaleFactor := 1.0
	if l.Scaling != nil {
		scaleFactor = *l.Scaling
	}

	canvas.Start(int(float64(width)*scaleFactor), int(float64(height)*scaleFactor))
	canvas.Scale(scaleFactor)

	if l.ShowBackground {
		canvas.Rect(0, 0, width, height, "fill:#D9D9D9")
	}

	invert := false

	for _, tile := range tiles {
		if tile.Size != currSize {
			invert = !invert
			currSize = tile.Size
		}

		tile.draw(canvas, maxSize)
	}

	canvas.Gend()
	canvas.End()

	return nil
}

func (t *Tile) offset(maxSize int) int {
	diff := int(math.Log2(float64(maxSize)) - math.Log2(float64(t.Size)))
	if diff == 0 {
		return 0
	}

	offset := maxSize * tileMultiplier / 8
	adder := offset

	for i := 0; i < diff-1; i++ {
		adder /= 2
		offset += adder
	}

	return offset
}

func (t *Tile) draw(canvas *svg.SVG, maxSize int) {
	var tileType int

	if t.Type == nil || *t.Type == 0 {
		tileType = rand.Intn(15) + 1
	} else {
		tileType = *t.Type
	}

	invert := int(math.Log2(float64(maxSize))-math.Log2(float64(t.Size)))%2 != 0

	x := int(float64(t.X*tileMultiplier)*0.5) + t.offset(maxSize)
	y := int(float64(t.Y*tileMultiplier)*0.5) + t.offset(maxSize)
	size := t.Size * tileMultiplier

	switch tileType {
	case 1:
		Tile1(canvas, x, y, size, invert)
	case 2:
		Tile2(canvas, x, y, size, invert)
	case 3:
		Tile3(canvas, x, y, size, invert)
	case 4:
		Tile4(canvas, x, y, size, invert)
	case 5:
		Tile5(canvas, x, y, size, invert)
	case 6:
		Tile6(canvas, x, y, size, invert)
	case 7:
		Tile7(canvas, x, y, size, invert)
	case 8:
		Tile8(canvas, x, y, size, invert)
	case 9:
		Tile9(canvas, x, y, size, invert)
	case 10:
		Tile10(canvas, x, y, size, invert)
	case 11:
		Tile11(canvas, x, y, size, invert)
	case 12:
		Tile12(canvas, x, y, size, invert)
	case 13:
		Tile13(canvas, x, y, size, invert)
	case 14:
		Tile14(canvas, x, y, size, invert)
	case 15:
		Tile15(canvas, x, y, size, invert)
	}
}
