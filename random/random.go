package random

import "math/rand"

// R is the random number generator
var R *rand.Rand

// InitRand initializes the random number generator
func InitRand(seed int64) {
	R = rand.New(rand.NewSource(seed))
}
