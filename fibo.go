package main

import "fmt"

func main() {
	fmt.Println(fibo(10))
	fmt.Println(fibo(20))
	fmt.Println(fibo(30))
	fmt.Println(fibo(40))
	fmt.Println(fibo(50))
	fmt.Println(fibo(60))
	fmt.Println(fibo(70))
	fmt.Println(fibo(80))
	fmt.Println(fibo(90))
	fmt.Println(fibo(100))
}

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
