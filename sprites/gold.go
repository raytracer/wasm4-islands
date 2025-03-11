package sprites

import "cart/w4"

// gold
const goldWidth = 8
const goldHeight = 8
const goldFlags = 1 // BLIT_2BPP
var gold = [16]byte{0x00, 0x00, 0x3f, 0x00, 0xd5, 0xc0, 0xd9, 0xc0, 0xd5, 0xc0, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00}

// DrawGold draws gold at the given position
func DrawGold(x, y int) {
	*w4.DRAW_COLORS = 0x4321
	w4.Blit(&gold[0], x, y, goldWidth, goldHeight, goldFlags)
}
