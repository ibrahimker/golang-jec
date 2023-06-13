package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(GetWord(i))
	}
}

// 0%3 = 0
// 1%3 = 1
// 2%3 = 2
// 3%3 = 0
// 4%3 = 1
// dst
// 6%3=0
func GetWord(in int, options ...string) (out string) {
	if in%3 == 0 && in%5 == 0 {
		return "fizzbuzz"
	} else if in%3 == 0 {
		return "fizz"
	} else if in%5 == 0 {
		return "buzz"
	} else if in%4 == 0 {
		return "ho"
	} else {
		return fmt.Sprintf("%d", in)
	}
}
