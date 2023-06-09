package main

import (
	"app/controller"
	"fmt"
)

func main() {

	users := controller.GenerateUser(100)

	search := controller.SearchPhoneNumber(users, "102")
	for _, user := range search {
		fmt.Println(user)
	}
}
