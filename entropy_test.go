package strength

import (
	"math"
	"testing"
)

func TestLogX(t *testing.T) {
	expected := math.Round(math.Log2(7) / math.Log2(2))
	actual := math.Round(logX(2, 7))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(10) / math.Log2(2))
	actual = math.Round(logX(2, 10))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(11) / math.Log2(2))
	actual = math.Round(logX(2, 11))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLogPow(t *testing.T) {
	expected := math.Round(math.Log2(math.Pow(7, 8)))
	actual := math.Round(logPow(7, 8, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(10, 11)))
	actual = math.Round(logPow(10, 11, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(11, 17)))
	actual = math.Round(logPow(11, 17, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log10(math.Pow(13, 21)))
	actual = math.Round(logPow(13, 21, 10))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
