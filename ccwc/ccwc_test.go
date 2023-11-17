package ccwc

import "testing"

func GetNumberOfCharactersTest(t *testing.T) {
	t.Parallel()
	filepath := "test.txt"
	want := 342190
	got := GetNumberOfCharacters(filepath)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
