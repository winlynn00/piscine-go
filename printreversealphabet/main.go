package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
		if i == 'a' {
			z01.PrintRune('\n')
		}
	}
}
