package main

import "testing"

func TestGetWord(t *testing.T) {
	t.Run("Should Return Fizz", func(t *testing.T) {
		got := GetWord(3)
		want := "fizz"
		if got != want {
			t.Errorf("GetWord() = %v, want %v", got, want)
		}
	})

	t.Run("Should Return Buzz", func(t *testing.T) {
		got := GetWord(5)
		want := "buzz"
		if got != want {
			t.Errorf("GetWord() = %v, want %v", got, want)
		}
	})

	t.Run("Should Return FizzBuzz", func(t *testing.T) {
		got := GetWord(15)
		want := "fizzbuzz"
		if got != want {
			t.Errorf("GetWord() = %v, want %v", got, want)
		}
	})

	t.Run("Should Return Number", func(t *testing.T) {
		got := GetWord(4)
		want := "4"
		if got != want {
			t.Errorf("GetWord() = %v, want %v", got, want)
		}
	})
}
