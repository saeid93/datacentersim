package main

import (
	"fmt"
	"math"
	"math/rand"
)

func startserver(users int) float64 {
	return math.Pow(float64(users), 3)
}

func main() {
	var users int = 5
	fmt.Printf("datacenter started with %d servers\n", rand.Intn(100))
	fmt.Printf("and also %g users are now connnected to dat\n", startserver(users))
}
