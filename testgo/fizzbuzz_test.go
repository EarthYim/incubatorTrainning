package fizzbuzz

import "testing"

func TestFizzBuzz1(t *testing.T) {
	want := "1"
	input := 1
	if FizzBuzz(input) != want {
		t.Errorf("want %s, got %s", want, FizzBuzz(input))
	}
}
