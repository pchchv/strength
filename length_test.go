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
