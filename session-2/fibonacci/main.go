package main

func main() {

}

// Fib adalah function untuk menghitung angka fibonacci ke n
// 1, 1, 2, 3, 5, 8
// fib(0) = 0
// fib(1) = 1
// fib(2) = fib(1) + fib(0) = 1 + 0 = 1
// fib(3) = fib(2) + fib(1) = 1 + 1 = 2
// fib(4) = fib(3) + fib(2) = 2 + 1 = 3
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
