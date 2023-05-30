package strength

func deleteRuneAt(runes []rune, i int) []rune {
	if i >= len(runes) || i < 0 {
		return runes
	}

	copy(runes[i:], runes[i+1:])

	runes[len(runes)-1] = 0

	return runes[:len(runes)-1]
}

func removeMoreThanTwoRepeatingChars(s string) string {
	var prevPrev rune
	var prev rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == prev && r == prevPrev {
			runes = deleteRuneAt(runes, i)
			i--
		}
		prevPrev = prev
		prev = r
	}

	return string(runes)
}
