package piscine

import "fmt"

func Swap(a *int, b *int) {
	*a = *b
	*b = *a
}
