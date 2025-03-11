package sprites

import "cart/w4"

// grass
const grassWidth = 16
const grassHeight = 8
const grassFlags = 0 // BLIT_1BPP
const grassOffset = 8

var grass = [16]byte{0x03, 0xc0, 0x0f, 0xf0, 0x3f, 0xfc, 0xff, 0xff, 0xff, 0xff, 0x3f, 0xfc, 0x0f, 0xf0, 0x03, 0xc0}

// DrawGrass draws grass at the given position
func DrawGrass(x, y int) {
	*w4.DRAW_COLORS = 0x20
	w4.Blit(&grass[0], x-grassOffset, y-grassHeight, grassWidth, grassHeight, grassFlags)
}
