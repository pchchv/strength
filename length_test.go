package strength

import "testing"

func TestRemoveMoreThanTwoFromSequence(t *testing.T) {
	actual := removeMoreThanTwoFromSequence("12345678", "0123456789")
	expected := "12"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoFromSequence("abcqwertyabc", "qwertyuiop")
	expected = "abcqwabc"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoFromSequence("", "")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoFromSequence("", "12345")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestRemoveRepeatingChars(t *testing.T) {
	actual := removeMoreThanTwoRepeatingChars("aaaa")
	expected := "aa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoRepeatingChars("bbbbbbbaaaaaaaaa")
	expected = "bbaa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoRepeatingChars("ab")
	expected = "ab"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = removeMoreThanTwoRepeatingChars("")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestGetReversedString(t *testing.T) {
	actual := getReversedString("abcd")
	expected := "dcba"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = getReversedString("1234")
	expected = "4321"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
