package main

import "fmt"

func main() {
	fmt.Println(contohRekursif(10))
}

func contohRekursif(a int) int {
	if a == 0 {
		return 0
	}
	fmt.Println(a)
	return contohRekursif(a - 1)
}
