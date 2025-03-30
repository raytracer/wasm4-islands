package sprites

import "cart/w4"

// ship_front
const ship_frontWidth = 64
const ship_frontHeight = 64
const ship_frontFlags = 1   // BLIT_2BPP
const ship_frontOffset = 32 // BLIT_2BPP

var ship_front = [1024]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x0a, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa0, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xc0, 0x00, 0x0f, 0xec, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0xf0, 0x00, 0xb7, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xbc, 0x0f, 0x57, 0x7f, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x23, 0xff, 0x3f, 0xd5, 0xdf, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc8, 0xff, 0x0d, 0xbd, 0xdd, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xf0, 0xc0, 0x0d, 0x6b, 0xdd, 0x5f, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x7f, 0xf0, 0x0d, 0x57, 0xfd, 0x7f, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3d, 0x57, 0xcf, 0x0d, 0x55, 0xf7, 0xd7, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xcd, 0x55, 0xff, 0xc3, 0x55, 0xff, 0x57, 0x55, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x0d, 0x55, 0xef, 0xf0, 0x95, 0xdd, 0x57, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x0d, 0x55, 0xd5, 0x7f, 0x7d, 0x5d, 0x57, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x0d, 0x55, 0xd5, 0x7f, 0xc3, 0xf5, 0x57, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x03, 0xc3, 0xd5, 0xd5, 0x75, 0xfc, 0x35, 0x57, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x83, 0x03, 0xfc, 0x3d, 0xd5, 0x75, 0x7f, 0x35, 0x7f, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x03, 0x03, 0x7f, 0xc3, 0xd5, 0x7d, 0x57, 0xff, 0xd5, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x0f, 0x03, 0x57, 0xf3, 0xd5, 0x5d, 0x55, 0xff, 0x09, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x3f, 0xf3, 0x55, 0x7f, 0xd5, 0x5f, 0x55, 0xff, 0xf3, 0xd7, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x0f, 0xfd, 0x55, 0x5f, 0xdf, 0xfd, 0x55, 0xff, 0xfc, 0x37, 0x00, 0x00, 0x00, 0x00, 0x03, 0x3c, 0x03, 0x57, 0xd5, 0x5f, 0xf0, 0x3d, 0x55, 0xff, 0xef, 0x0f, 0x3c, 0x00, 0x00, 0x00, 0x0c, 0x0f, 0xc3, 0x56, 0xf5, 0x5f, 0xff, 0x03, 0xd5, 0x7c, 0xe9, 0xf0, 0xdc, 0x00, 0x00, 0x00, 0x0c, 0x0d, 0xf3, 0x55, 0x7d, 0x5f, 0xff, 0xf0, 0x3d, 0x7c, 0xf7, 0xff, 0xdc, 0x00, 0x00, 0x00, 0x0c, 0x0d, 0x7f, 0xd5, 0x75, 0x5e, 0xde, 0xfc, 0x03, 0x3c, 0x3f, 0xff, 0x70, 0x00, 0x00, 0x00, 0x0c, 0x0d, 0x77, 0xd5, 0xdf, 0x5f, 0xde, 0xef, 0xc2, 0xfc, 0x3f, 0xfd, 0xf0, 0x00, 0x00, 0x00, 0x30, 0x35, 0xd7, 0x55, 0xc3, 0x5f, 0xcf, 0xba, 0xf2, 0xff, 0xfd, 0x75, 0xf0, 0x00, 0x00, 0x00, 0x30, 0x35, 0xcd, 0x55, 0xc0, 0x9f, 0xc3, 0xee, 0xbf, 0xff, 0xff, 0xdf, 0xc0, 0x00, 0x00, 0x00, 0x30, 0xf7, 0x1d, 0x55, 0xc0, 0x3f, 0xc3, 0xea, 0xaf, 0xff, 0xeb, 0x7b, 0xc0, 0x00, 0x00, 0x00, 0xc0, 0xdd, 0x1d, 0x57, 0x00, 0x03, 0xc0, 0xeb, 0xff, 0xff, 0xfd, 0xdf, 0x00, 0x00, 0x00, 0x00, 0xc3, 0x5c, 0x1f, 0xfc, 0x00, 0x03, 0xc0, 0xef, 0xff, 0xff, 0xee, 0xbf, 0x00, 0x00, 0x00, 0x00, 0xcd, 0x5c, 0x1f, 0xdf, 0xc0, 0x03, 0xc0, 0xff, 0x5f, 0xff, 0x9e, 0xff, 0x00, 0x00, 0x00, 0x03, 0x0d, 0xfc, 0x1f, 0x15, 0xfc, 0x03, 0xfb, 0xf5, 0xf7, 0xff, 0x97, 0xfe, 0x00, 0x00, 0x00, 0x03, 0x37, 0x03, 0x1f, 0x17, 0xff, 0xc3, 0xff, 0xff, 0x7f, 0xfd, 0x7f, 0xfe, 0x80, 0x00, 0x00, 0x03, 0x3c, 0x00, 0x07, 0x5f, 0xeb, 0xff, 0xff, 0xad, 0xff, 0xeb, 0xff, 0xfa, 0x80, 0x00, 0x00, 0x3f, 0xf0, 0x00, 0xff, 0xf9, 0x7b, 0xff, 0xfe, 0xa7, 0xfd, 0xff, 0xff, 0xea, 0x80, 0x00, 0x00, 0x0f, 0xff, 0xff, 0xff, 0x3f, 0xff, 0xff, 0xff, 0xfa, 0xd7, 0xff, 0xfc, 0xaa, 0x80, 0x00, 0x00, 0x03, 0xff, 0xff, 0xff, 0xc7, 0xfd, 0x7f, 0xff, 0x5e, 0xbf, 0xff, 0xc2, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x9f, 0xff, 0xff, 0xeb, 0xdf, 0xfd, 0xea, 0xff, 0xfc, 0x2a, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xad, 0x97, 0x66, 0xff, 0x7f, 0xfd, 0xbf, 0xff, 0xf2, 0xaa, 0xa8, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xab, 0xba, 0xee, 0xd7, 0xff, 0xd7, 0xff, 0xff, 0x0a, 0xaa, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x03, 0xf9, 0x55, 0x65, 0xbe, 0xb9, 0x7f, 0xff, 0xf0, 0xaa, 0xaa, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xaa, 0xaa, 0xaa, 0xeb, 0xff, 0xff, 0x2a, 0xaa, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xaa, 0xaa, 0xa8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0a, 0xaa, 0xaa, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xf2, 0xaa, 0xaa, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0xff, 0xff, 0xff, 0xfc, 0x0a, 0xaa, 0xaa, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2b, 0xff, 0xff, 0xff, 0x02, 0xaa, 0xaa, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a, 0xff, 0x00, 0x00, 0xaa, 0xaa, 0xaa, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a, 0x80, 0xaa, 0xaa, 0xaa, 0xaa, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0xaa, 0xaa, 0xaa, 0xaa, 0xa8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xaa, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func DrawShipFrontTile(x, y float64, flip bool) {
	tx, ty := IsoTransform(x, y)
	flags := uint(ship_frontFlags)
	if flip {
		flags |= w4.BLIT_FLIP_X
	}
	w4.Blit(&ship_front[0], tx+CameraX-ship_frontOffset, ty+CameraY-ship_frontHeight, ship_frontWidth, ship_frontHeight, flags)
}
