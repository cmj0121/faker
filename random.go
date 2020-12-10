package faker

import (
	"math/rand"
	"time"
)

// the interface for the random generator
type Random interface {
	Seed(int64)
	Int63() int64
	Float64() float64
}

var generator Random

func init() {
	// set the default random generator
	SetGenerator(rand.New(rand.NewSource(0)))
}

func SetGenerator(random Random) {
	generator = random
	now := time.Now().UnixNano()
	generator.Seed(now)
	generator.Seed(generator.Int63())
}

func Seed(seed int64) {
	generator.Seed(seed)
	return
}
