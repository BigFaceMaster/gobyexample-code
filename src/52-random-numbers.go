package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p()

	p(rand.Float64())
	p()
	p((rand.Float64() * 5) + 5, ',')
	p((rand.Float64() * 5) + 5)
	p()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
