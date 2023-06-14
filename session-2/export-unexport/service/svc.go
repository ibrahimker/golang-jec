package service

import "fmt"

var user string

func init() {
	user = "ibam"
	fmt.Println("init dipanggil pertama")
}

func init() {
	user = "ibam"
	fmt.Println("init dipanggil ketiga")
}

func Create(name string) {
	fmt.Println("this is public (exported) function", user)
}

func get() {
	fmt.Println("this is private (unexported) function")
}
