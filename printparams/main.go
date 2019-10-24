package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	lenght := 0
	for range arg {
		lenght++
	}
	for i := 1; i < lenght; i++ {
		arg2 := os.Args[i]
		for _, a := range arg2 {
			z01.PrintRune(a)
		}
		z01.PrintRune(10)
	}
}
