package strength

func deleteRuneAt(runes []rune, i int) []rune {
	if i >= len(runes) || i < 0 {
		return runes
	}

	copy(runes[i:], runes[i+1:])

	runes[len(runes)-1] = 0

	return runes[:len(runes)-1]
}
