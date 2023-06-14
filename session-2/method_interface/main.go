package main

import (
	"fmt"
	"github.com/ibrahimker/golang-jec/session-2/method_interface/service"
)

func main() {
	var db []service.User
	userSvc := service.NewUserService(db)

	u1 := service.User{Name: "budi"}
	u2 := service.User{Name: "dwi"}
	userSvc.Register(u1)
	userSvc.Register(u2)

	results := userSvc.Get()
	for _, value := range results {
		fmt.Println(value)
	}
}
