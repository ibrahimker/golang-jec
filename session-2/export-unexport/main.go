package main

import "github.com/ibrahimker/golang-jec/session-2/export-unexport/service"

func main() {
	service.Create("nama")
	//service.get() // will get error due to private function
}
