package sprites

import "cart/w4"

// numbers
const numbersWidth = 30
const numbersHeight = 5
const numbersFlags = 0 // BLIT_1BPP
var numbers = [19]byte{0xe7, 0xf9, 0xe7, 0xfe, 0x92, 0x6c, 0x86, 0xda, 0x7f, 0xff, 0x9f, 0xe9, 0x84, 0x9a, 0x69, 0xe7, 0xf3, 0xf9, 0xe4}

// DrawNumber draws numbers at the given position
func DrawNumber(number, x, y, max int) {
	*w4.DRAW_COLORS = 0x0030
	startX := (max-1)*5 + x

	by10 := number
	remainder10 := number

	for {
		remainder10 = by10 % 10
		by10 /= 10

		w4.BlitSub(&numbers[0], startX, y, 3, 5, uint(remainder10*3), 0, numbersWidth, numbersFlags)
		startX -= 5

		if by10 == 0 {
			break
		}
	}
}
