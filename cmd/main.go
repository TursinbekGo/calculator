package main

import (
	"app/controller"
)

func main() {

	users := controller.GenerateUser(100)
	// for _, user := range users {
	// 	fmt.Println(user)
	// }
	controller.SearchPhoneNumber(users, "102")
}
