package main

import (
	"fmt"
	"structs/customer"
)

func main() {
	user := customer.New("John", 25, 322, 0, true)
	// user, err := user.CalculateDiscount()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("User: %+v", user)

	err := user.CalculateDiscountPointer()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("User: %+v", user)
}
