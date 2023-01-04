package main

import (
	"errors"
	"fmt"
	"structs/customer"
)

const DEFAULT_DISCOUNT = 500

func main() {
	user := customer.New("John", 25, 322, 0, true)
	fmt.Println(user)
	user.CalcDicount = func() (int, error) {
		if !user.Discount {
			return 0, errors.New("DISCOUNT IS NOT AVAILABLE")
		}
		result := DEFAULT_DISCOUNT - user.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	fmt.Println(customer.CalcPrice(*user, 500))
}
