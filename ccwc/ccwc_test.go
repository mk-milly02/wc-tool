package ccwc_test

import (
	"testing"
	"wctool/ccwc"
)

func TestGetNumberOfBytes(t *testing.T) {
	t.Parallel()
	filepath := "../test.txt"
	want := 342190
	got := ccwc.GetNumberOfBytes(filepath)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfNewLines(t *testing.T) {
	t.Parallel()
	filepath := "../test.txt"
	want := 7145
	got := ccwc.GetNumberOfLines(filepath)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfWords(t *testing.T) {
	t.Parallel()
	filepath := "../test.txt"
	want := 58164
	got := ccwc.GetNumberOfWords(filepath)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfCharacters(t *testing.T) {
	t.Parallel()
	filepath := "../test.txt"
	want := 339292
	got := ccwc.GetNumberOfCharacters(filepath)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
