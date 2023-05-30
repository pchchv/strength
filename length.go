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

func removeMoreThanTwoFromSequence(s, seq string) string {
	seqRunes := []rune(seq)
	runes := []rune(s)
	matches := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(seqRunes); j++ {
			if i >= len(runes) {
				break
			}

			r := runes[i]
			r2 := seqRunes[j]
			if r != r2 {
				matches = 0
				continue
			}

			// found a match, advance the counter
			matches++
			if matches > 2 {
				runes = deleteRuneAt(runes, i)
			} else {
				i++
			}
		}
	}

	return string(runes)
}

func getReversedString(s string) string {
	n := 0
	rune := make([]rune, len(s))

	for _, r := range s {
		rune[n] = r
		n++
	}

	rune = rune[0:n]

	// reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	// convert back to UTF-8
	return string(rune)
}
