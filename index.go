package piscine

func Index(s string, toFind string) int {
	sToRune := []rune(s)
	toFindToRune := []rune(toFind)
	length := 0
	for range toFindToRune {
		length++
	}
	for index, str := range sToRune {
		if length > 0 && str == toFindToRune[0] {
			return index
		}
	}
	return -1
}
