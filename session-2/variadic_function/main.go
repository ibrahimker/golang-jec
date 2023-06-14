package main

import (
	"fmt"
)

func main() {
	print()
	print("a", "b")
}

func print(names ...string) {
	//names = []string{"a","b"}
	fmt.Println("masuk")
	fmt.Println("names", names)
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("masuk 2")
}
