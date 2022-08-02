package cookbook

import (
	"fmt"
	"math/rand"
	"time"
)

func SimpleRand() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func SeedRand() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
