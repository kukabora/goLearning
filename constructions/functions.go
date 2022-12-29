package main

import "fmt"

// func myFunc(name string) {
// 	fmt.Println("Hello, i'm", name)
// }

// func main() {
// 	myFunc("Alex")
// 	myFunc("Alex")
// 	myFunc("Alex")
// 	myFunc("Alex")
// 	myFunc("Alex")
// 	myFunc("Alex")
// }

// func divideTwoNumbers(a int, b int) (int, error) {
// 	if a == 0 || b == 0 {
// 		return 0, errors.New("Zero division")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divideTwoNumbers(0, 3)
// 	fmt.Println(result, err)
// }

// func sum(a int, b int) (res int, err error) {
// 	if a == 0 || b == 0 {
// 		err = errors.New("Some number is zero")
// 		return 0, err
// 	}
// 	return a + b, err
// }

// func sum(a ...int) int {
// 	count := 0
// 	for i := 0; i < len(a); i++ {
// 		count += a[i]
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(sum(3, 2, 5))
// }

// АНОНИМНЫЕ ФУНКЦИИ

// type MyFunc func()

// func main() {
// myfunc := func() {
// 	fmt.Println("Hello from myfunc")
// }
// myfunc()

// var myFunc MyFunc
// myFunc = func() {
// 	fmt.Println("Hello world")
// }
// myFunc()

// }

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
