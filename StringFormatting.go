package main

import "fmt"

type point struct {
	x, y int
}

func main(){
	p := point{1, 2}

	fmt.Println(p)
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 101)

	fmt.Printf("%c\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
}