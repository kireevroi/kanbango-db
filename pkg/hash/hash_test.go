package hash

import "testing"

func TestHash(t *testing.T) {
	str := "lolol"
	hash, err := HashPassword(str)
	if err != nil {
		t.Errorf("Got error: %v, want: other things.", err)
	}
	if !CheckPassword(hash, str) {
		t.Errorf("Result was incorrect, got: %v, want: other things.", hash)
	}
}