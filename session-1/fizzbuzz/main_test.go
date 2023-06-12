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

func TestGetWord1(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get fizz",
			args: args{
				i: 3,
			},
			want: "fizz",
		},
		{
			name: "should get buzz",
			args: args{
				i: 5,
			},
			want: "buzz",
		},
		{
			name: "should get fizzbuzz",
			args: args{
				i: 15,
			},
			want: "fizzbuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWord(tt.args.i); got != tt.want {
				t.Errorf("GetWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
