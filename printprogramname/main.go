package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args[0]
	for _, a := range arg {
		z01.PrintRune(a)
	}
	z01.PrintRune(10)
}