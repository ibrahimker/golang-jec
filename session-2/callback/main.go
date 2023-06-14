package main

import (
	"fmt"
)

func main() {
	fmt.Println("dengan callback 1")
	loginUser("Airell", callback)
	fmt.Println("============")
	fmt.Println("dengan callback 2")
	loginUser("Budi", callback2)
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
