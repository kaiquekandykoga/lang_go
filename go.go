package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("ありがとう！")
	fmt.Println("[time] time.Now() =", time.Now())
	fmt.Println("[math/rand] rand.Intn(9) =", rand.Intn(9))
	fmt.Printf("[math] math.Sqrt(16) = %g \n", math.Sqrt(16))
	fmt.Println("[math] math.Pi =", math.Pi)

	fmt.Println("add_0(3, 5) =", add_0(3, 5))
	fmt.Println("add_1(3, 5) =", add_1(3, 5))

	koga, kandy := swap("Koga", "Kandy")
	fmt.Println("swap(\"Koga\", \"Kandy\") =", koga, kandy)
}

func add_0(a int, b int) int {
	return a + b
}

// Go's shorthand where consecutive parameters of the same type
// can share a single type declaration
func add_1(x, y int) int {
	return x + y
}

// Go's multiple return values
func swap(x, y string) (string, string) {
	return y, x
}
