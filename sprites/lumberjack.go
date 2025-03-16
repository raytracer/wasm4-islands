package sprites

import "cart/w4"

// lumberjack
const lumberjackWidth = 64
const lumberjackHeight = 48
const lumberjackFlags = 1 // BLIT_2BPP
const lumberjackOffset = 32

var lumberjack = [768]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x35, 0xbc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xda, 0x5b, 0xc3, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x55, 0x96, 0xfe, 0xac, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0xa5, 0x55, 0x6f, 0xac, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0xf5, 0x5a, 0x5a, 0x96, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x5f, 0x03, 0x5a, 0x55, 0xa5, 0x65, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3d, 0xad, 0xcd, 0x55, 0x95, 0x59, 0x5a, 0x96, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xd6, 0xf5, 0xf5, 0x55, 0x59, 0x56, 0x95, 0x69, 0x5b, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x6a, 0xd7, 0x3f, 0x55, 0x66, 0x95, 0x69, 0x56, 0x95, 0xbc, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf5, 0xaf, 0xbe, 0x0f, 0xf5, 0x95, 0x69, 0x56, 0x95, 0x65, 0xab, 0x00, 0x00, 0x00, 0x00, 0x03, 0x57, 0xb5, 0xf0, 0x00, 0xff, 0x55, 0x56, 0x95, 0x65, 0x57, 0xd5, 0xf0, 0x00, 0x00, 0x00, 0x03, 0x5b, 0xd6, 0xf0, 0x0a, 0xff, 0xd5, 0x55, 0x65, 0x5a, 0x5f, 0xfe, 0x5c, 0x00, 0x00, 0x00, 0x03, 0xef, 0xbf, 0xb0, 0xaa, 0xef, 0xfd, 0x55, 0x5a, 0x55, 0x7b, 0xbf, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x3e, 0xeb, 0xbf, 0xaa, 0xee, 0xff, 0xd5, 0x55, 0xa5, 0xfb, 0xff, 0x00, 0x00, 0x00, 0x00, 0x03, 0xfe, 0xeb, 0xb5, 0xeb, 0xee, 0xbf, 0xfd, 0x55, 0x57, 0xbb, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x3d, 0x5e, 0xed, 0xaf, 0xe9, 0xef, 0xd7, 0xff, 0xd5, 0x5f, 0xbb, 0xfb, 0x00, 0x00, 0x00, 0x00, 0xd4, 0x0e, 0xf5, 0xfa, 0x96, 0xee, 0xfd, 0x7f, 0xfd, 0x7f, 0xaf, 0xbb, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xd7, 0xfa, 0x5a, 0xee, 0xef, 0xef, 0xbf, 0xfe, 0xfa, 0xbb, 0xf0, 0x00, 0x00, 0x00, 0x03, 0x95, 0xbf, 0xf6, 0xab, 0xef, 0x65, 0xff, 0xaf, 0xff, 0xbd, 0x7b, 0xac, 0x00, 0x00, 0x00, 0x03, 0xda, 0xfa, 0xee, 0xab, 0xef, 0x7e, 0xbf, 0xee, 0xef, 0xd7, 0x7b, 0xfb, 0x00, 0x00, 0x00, 0x0e, 0xbf, 0xfa, 0xee, 0xab, 0xed, 0xf5, 0xff, 0xbe, 0xed, 0x7f, 0x7b, 0xac, 0x00, 0x00, 0x00, 0xfa, 0xae, 0xea, 0xee, 0xaa, 0xb6, 0xbe, 0x7f, 0xae, 0xed, 0xff, 0x7b, 0xac, 0x00, 0x00, 0x03, 0xbf, 0xae, 0xea, 0xee, 0xaa, 0xb7, 0x95, 0xbf, 0xee, 0xed, 0xfb, 0x7b, 0xec, 0x00, 0x00, 0x03, 0xbd, 0xba, 0xeb, 0xbb, 0x95, 0xfd, 0x69, 0xff, 0xfe, 0xed, 0xeb, 0x7b, 0xf0, 0x00, 0x00, 0x57, 0xee, 0xbb, 0xef, 0xff, 0xab, 0xeb, 0xd7, 0xff, 0xfe, 0xed, 0xeb, 0x6e, 0xaf, 0x00, 0x05, 0x6b, 0xbf, 0x7b, 0xea, 0xaa, 0xbd, 0x56, 0xb7, 0xff, 0xfe, 0xed, 0xfb, 0xaa, 0xab, 0xc0, 0x55, 0xab, 0xbf, 0xbe, 0xfa, 0x97, 0xd5, 0x59, 0x5f, 0xff, 0xfe, 0xed, 0xea, 0xa5, 0x6a, 0xa9, 0x55, 0x56, 0xfa, 0xef, 0xf9, 0x7d, 0x65, 0xa5, 0x7f, 0xfb, 0xfe, 0xed, 0xea, 0xa9, 0x6a, 0x95, 0x05, 0x55, 0xaa, 0xfe, 0xab, 0xd5, 0x5a, 0x95, 0xba, 0xaa, 0xbe, 0xae, 0xa5, 0x6a, 0xa5, 0x50, 0x00, 0x59, 0x5a, 0xa9, 0x7d, 0x55, 0xa5, 0x5a, 0xaa, 0xaa, 0x6b, 0xbe, 0xaa, 0xaa, 0x55, 0x00, 0x00, 0x05, 0x66, 0x67, 0xd6, 0x5a, 0x95, 0xfe, 0xaa, 0xaa, 0x95, 0xaa, 0xaa, 0xa5, 0x50, 0x00, 0x00, 0x00, 0x55, 0x5d, 0x55, 0xa5, 0x5f, 0xaa, 0x6a, 0xfe, 0xa9, 0x5a, 0xa9, 0x55, 0x00, 0x00, 0x00, 0x00, 0x05, 0x75, 0x5a, 0x55, 0x7e, 0x9a, 0xab, 0xa7, 0xaa, 0xaa, 0xa5, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x55, 0x5b, 0xea, 0xa9, 0xab, 0xab, 0xaa, 0xaa, 0x95, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x55, 0xfe, 0x95, 0x6a, 0xaf, 0xaf, 0xa5, 0x69, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x57, 0xaa, 0xa6, 0xaa, 0xeb, 0xee, 0xaa, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x6a, 0x6a, 0xab, 0xbe, 0xaa, 0xa9, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x6a, 0xab, 0xbb, 0xaa, 0xa5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x56, 0xab, 0xbe, 0xaa, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x56, 0xaa, 0xaa, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x96, 0xa5, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5a, 0x95, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// DrawLumberjack at x, y
func DrawLumberjack(x int, y int) {
	*w4.DRAW_COLORS = 0x4320
	w4.Blit(&lumberjack[0], x-lumberjackOffset, y-lumberjackHeight, lumberjackWidth, lumberjackHeight, lumberjackFlags)
}

// DrawLumberjackTile at x, y
func DrawLumberjackTile(x int, y int, srcX byte, srcY byte) {
	*w4.DRAW_COLORS = 0x4320

	drawTile(x, y, srcX, srcY, lumberjack[:], lumberjackWidth, lumberjackHeight, lumberjackOffset)
}
