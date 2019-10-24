package piscine

func Index(s string, toFind string) int {
	if StrLen(s) > 0 && StrLen(toFind) > 0 {
		index := -1
		if StrLen(s) > StrLen(toFind) {
			for i := 0; i < StrLen(s); i++ {
				count := 0
				startIndex := 0
				for j := 0; j < StrLen(toFind); j++ {
					if s[i] == toFind[j] && j == 0 {
						startIndex = i
					}
					if s[i] == toFind[j] {
						count++
						i++
					}
				}
				if count == StrLen(toFind) {
					index = startIndex
					break
				}
			}
		}
		return index
	} else {
		return 0
	}
}
