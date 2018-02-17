package main

import (
	"fmt"
)

// Talker : define interface
type Talker interface {
	Talk()
}

// Greeter : define struct
type Greeter struct {
	name string
}

// Talk : struct Greeter has Talk method
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
