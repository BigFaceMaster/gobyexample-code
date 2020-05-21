package main

import "fmt"

type biu struct {
	biuX, biuY int
}

func main() {
	aBiu := biu{1, 2}
	fmt.Printf("%v\n", aBiu)
	fmt.Printf("%+v\n", aBiu)
	fmt.Printf("%#v\n", aBiu)
}
