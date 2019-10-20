package piscine

func UltimateDivMod(a *int, b *int) {
	a2 := *a
	*a = *a / *b
	*b = a2 % *b
}
