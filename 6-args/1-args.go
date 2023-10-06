package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	// fmt.Printf("%T\n", os.Args)
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	a := os.Args[1:]
	if len(a) != 3 {
		log.Println("Hello")
		return
	}
	fmt.Println(a)
}
