package main

import (
	"fmt"
	"math/big"
	"time"
)

func fiboIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		temp := b
		b = a + b
		a = temp
	}

	return b
}

func fiboMemoized() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		if val, ok := cache[n]; ok {
			return val
		}

		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}

	return fib
}

func fiboBigInt(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 2; i <= n; i++ {
		temp := new(big.Int).Set(b)

		b.Add(a, b)

		a.Set(temp)
	}

	return b
}

func fiboMatrix(n int) int {
	if n <= 1 {
		return n
	}

	A := [2][2]int{{1, 1}, {1, 0}}

	result := matrixPower(A, n-1)

	return result[0][0]
}

func matrixMultiply(A, B [2][2]int) [2][2]int {
	C := [2][2]int{}

	C[0][0] = A[0][0]*B[0][0] + A[0][1]*B[1][0]
	C[0][1] = A[0][0]*B[0][1] + A[0][1]*B[1][1]
	C[1][0] = A[1][0]*B[0][0] + A[1][1]*B[1][0]
	C[1][1] = A[1][0]*B[0][1] + A[1][1]*B[1][1]

	return C
}

func matrixPower(A [2][2]int, n int) [2][2]int {
	if n == 1 {
		return A
	}

	if n%2 == 0 {
		half := matrixPower(A, n/2)
		return matrixMultiply(half, half)
	} else {
		return matrixMultiply(A, matrixPower(A, n-1))
	}
}

func main() {
	testValues := []int{10, 20, 30, 40, 45}

	fmt.Println("=== Fixed Iterative Implementation ===")
	for _, n := range testValues[:3] {
		fmt.Printf("Fibonacci(%d) = %d\n", n, fiboIterative(n))
	}

	fmt.Println("\n=== Memoized Implementation ===")
	fib := fiboMemoized()
	for _, n := range testValues {
		start := time.Now()
		result := fib(n)
		duration := time.Since(start)
		fmt.Printf("Fibonacci(%d) = %d (took %v)\n", n, result, duration)
	}

	fmt.Println("\n=== Big.Int Implementation ===")
	for _, n := range testValues {
		start := time.Now()
		result := fiboBigInt(n)
		duration := time.Since(start)
		fmt.Printf("Fibonacci(%d) = %s (took %v)\n", n, result.String(), duration)
	}

	fmt.Println("\n=== Matrix Exponentiation Implementation (O(log n)) ===")
	for _, n := range testValues {
		start := time.Now()
		result := fiboMatrix(n)
		duration := time.Since(start)
		fmt.Printf("Fibonacci(%d) = %d (took %v)\n", n, result, duration)
	}

	n := 100
	fmt.Println("\n=== Calculating a Large Fibonacci Number ===")
	fmt.Printf("Fibonacci(%d) = %s\n", n, fiboBigInt(n).String())

	n = 1000000
	fmt.Println("\n=== Speed Demonstration of Matrix Method ===")
	start := time.Now()
	result := fiboBigInt(n)
	duration := time.Since(start)
	fmt.Printf("Fibonacci(%d) has %d digits (calculated in %v)\n",
		n, len(result.String()), duration)
}
