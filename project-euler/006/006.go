package main

// Find the difference between the sum of the squares of the first one hundred natural
// numbers and the square of the sum.
import "fmt"

const N int = 100

func main() {
	var sumOfSquares int
	squareOfSum := N * (N + 1) / 2
	squareOfSum *= squareOfSum
	for i := 1; i <= N; i++ {
		sumOfSquares += i * i
	}
	fmt.Println(squareOfSum - sumOfSquares)
}
