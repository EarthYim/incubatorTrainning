package fizzbuzz

import "testing"

func TestFizzBuzzT(t *testing.T) {

	result := Say(4)
	want := 6

	if result != want {
		t.Error("Testing FizzBuzz")
	}
}
