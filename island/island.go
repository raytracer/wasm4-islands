package island

import (
	"cart/random"
	"cart/sprites"
	"math"

	"github.com/bits-and-blooms/bitset"
)

var islands = make([]Island, 9)

// Size of the map
const Size = 256

// Gold is the amount of gold
var Gold = 1000

var g = bitset.New(Size * Size)
var ng = bitset.New(Size * Size)

var seed = int64(12)

// Island is a island with a position and a size (maximum of 9 currently)
type Island struct {
	Width  byte
	Height byte
	x      byte
	y      byte
	Wood   byte
	Stone  byte
	Food   byte
}

// GenerateIslands generates up to 9 islands
func GenerateIslands() {
	islands[0] = Island{Width: byte(85), Height: byte(85), x: 0, y: 0, Wood: 100}
	islands[1] = Island{Width: byte(85), Height: byte(85), x: 85, y: 0, Wood: 100}
	islands[2] = Island{Width: byte(85), Height: byte(85), x: 170, y: 0, Wood: 100}
	islands[3] = Island{Width: byte(85), Height: byte(85), x: 0, y: 85, Wood: 100}
	islands[4] = Island{Width: byte(85), Height: byte(85), x: 85, y: 85, Wood: 100}
	islands[5] = Island{Width: byte(85), Height: byte(85), x: 170, y: 85, Wood: 100}
	islands[6] = Island{Width: byte(85), Height: byte(85), x: 0, y: 170, Wood: 100}
	islands[7] = Island{Width: byte(85), Height: byte(85), x: 85, y: 170, Wood: 100}
	islands[8] = Island{Width: byte(85), Height: byte(85), x: 170, y: 170, Wood: 100}

	for _, island := range islands {
		generateIsland(int(island.x), int(island.y), int(island.Width), int(island.Height))
	}
}

// FindIsland finds the island at the given position
func FindIsland(x, y int) *Island {
	for idx := range islands {
		if x >= int(islands[idx].x) && x < int(islands[idx].x)+int(islands[idx].Width) && y >= int(islands[idx].y) && y < int(islands[idx].y)+int(islands[idx].Height) {
			return &islands[idx]
		}
	}
	return nil
}

// CheckLand checks if the land is available on the island
func CheckLand(x, y, width, height int) bool {
	for xi := x - width + 1; xi <= x; xi++ {
		for yi := y - height + 1; yi <= y; yi++ {
			if !g.Test(uint(yi*Size + xi)) {
				return false
			}
		}
	}

	return true
}

// good enough for now
func generateIsland(startX, startY, w, h int) {
	cx, cy := float64(w)/2, float64(h)/2
	mr := math.Min(cx, cy)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if math.Hypot(float64(x)-cx, float64(y)-cy)+random.R.Float64()*mr < mr {
				g.Set(uint((y+startY)*Size + (x + startX)))
			}
		}
	}
	for s := 0; s < 2; s++ {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				c := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						yy, xx := y+dy, x+dx
						if yy >= 0 && yy < h && xx >= 0 && xx < w && g.Test(uint((yy+startY)*Size+(xx+startX))) {
							c++
						}
					}
				}

				if c > 4 {
					ng.Set(uint((y+startY)*Size + (x + startX)))
				} else {
					ng.Clear(uint((y+startY)*Size + (x + startX)))
				}
			}
		}
		g = ng
	}
}

// CheckIsLand checks if the given position is land
func CheckIsLand(x, y int) bool {
	return g.Test(uint(y*Size + x))
}

// DrawIsland draws the island
func DrawIsland() {
	minX, minY := sprites.ScreenToIsoTransform(-sprites.CameraX, -sprites.CameraY)
	max, maxY := sprites.ScreenToIsoTransform(-sprites.CameraX, -sprites.CameraY+160)
	for x := 0; x < Size; x++ {
		for y := 0; y < Size; y++ {
			if (x+y) >= (minX+minY) && (x+y) <= (max+maxY+4) {
				if g.Test(uint(y*Size + x)) {
					sx, sy := sprites.IsoTransform(x, y)
					sprites.DrawGrass(sx+sprites.CameraX, sy+sprites.CameraY)
				}
			}
		}
	}
}
