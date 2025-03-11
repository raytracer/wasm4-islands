package sprites

// CameraX is the x position of the camera
var CameraX = 0

// CameraY is the y position of the camera
var CameraY = 0

// TileWidth in Pixel
const TileWidth = 16

// TileHeight in Pixel
const TileHeight = 8

// IsoTransform converts isometric coordinates to screen coordinates
func IsoTransform[T int | float64](x, y T) (int, int) {
	screenX := ((x - y) * (TileWidth / 2))
	screenY := ((x + y) * (TileHeight / 2))
	return int(screenX), int(screenY)
}

// ScreenToIsoTransform converts screen coordinates to isometric coordinates
func ScreenToIsoTransform(sx, sy int) (int, int) {
	x := (float64(sx)/float64(TileWidth) + float64(sy)/float64(TileHeight))
	y := (float64(sy)/float64(TileHeight) - float64(sx)/float64(TileWidth))
	return int(x), int(y)
}
