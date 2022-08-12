package main

import (
	"datacenter/node"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const datacenterName = "South Carolina"

var (
	users             int = 4
	usersAccess       int = 2
	serverUtilization int = 87
)
var (
	valueToPoint string = "pointed value"
	pointer      *string
)

type Node struct {
	capacity int
	value    int
}

func startserver(users int) float64 {
	return math.Pow(float64(users), float64(usersAccess))
}

func main() {
	defer fmt.Printf("I defered leaving the beuilding until I'm sure server is turned off\n")
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("datacenter started with %d servers\n", rand.Intn(100))
	fmt.Printf("and also %g users are now connnected to dat\n", startserver(users))
	utilized := serverUtilization
	fmt.Printf("I infered that the server utilization is at %v\n", utilized)
	fmt.Printf("You cannot move the server from %s\n", datacenterName)
	for i := 0; i < int(startserver(users)); i++ {
		fmt.Printf("users %v is now conncted\n", i)
	}
	var on int = rand.Intn(4)
	fmt.Printf("on %v\n", on)
	if on == 1 {
		fmt.Printf("Server truned off")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("not the best time to start your server!\n")
	case t.Hour() < 17:
		fmt.Printf("right on time!\n")
	default:
		fmt.Printf("Too late for this shit!\n")
	}
	pointer = &valueToPoint
	fmt.Println(*pointer)
	e := node.New("SC", "Datacenter", 30, 20)
	e.AddedNode()
	var datacenters [10]int
	for i := 0; i < len(datacenters); i++ {
		// e := node.New("SC", "Datacenter", 30, 20)
		datacenters[i] = i
		fmt.Printf("%v", datacenters[i])
	}
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

}
