package main

import (
	"fmt"
)

func main() {
	//fmt.Println(greet("Airell", "Pancoran", 17))
	//print()
	//print(1, "a", "b")
	//_ = getUser(10)
	fmt.Println("dengan callback 1")
	loginUser("Airell", callback)
	fmt.Println("============")
	fmt.Println("dengan callback 2")
	loginUser("Budi", callback2)
}

func contohRekursif() {
	fmt.Println("halo")
}

func callback() {
	// ceritanya call API
	fmt.Println("notify ke service a kalau airell berhasil masuk")
}

func callback2() {
	// ceritanya call API
	fmt.Println("notify ke service b kalau budi berhasil masuk")
}

// function normal
func loginUser(n string, cb func()) {
	cb()
	fmt.Println("login " + n + " success")
}

func print(id int, names ...string) {
	nickname2("ibam")
	//names = []string{"a","b"}
	fmt.Println("masuk")
	fmt.Println("names", names)
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("masuk 2")
}

func getUser(opts ...int) []string {
	if opts[0] > 0 {
		return []string{"1"}
	}
	return []string{"1", "2"}
}
