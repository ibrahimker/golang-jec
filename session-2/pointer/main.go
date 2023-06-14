package main

import (
	"fmt"
)

func main() {
	/*	var a int
		a = 8
		fmt.Println(a)
		fmt.Println("a disimpan di", &a)
		fmt.Println("isi value &a", *&a)
	*/

	// model pertama
	a := 8
	a = addOne(a)
	fmt.Println(a)
	addOnex(&a)
	fmt.Println(a)
}

func addOne(n int) int {
	return n + 1
}

func addOnex(n *int) {
	*n = 10
}
