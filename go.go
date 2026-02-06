package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Go's variables in package level
var c, python, java bool

func main() {
	fmt.Println("ありがとう！")
	fmt.Println("[time] time.Now() =", time.Now())
	fmt.Println("[math/rand] rand.Intn(9) =", rand.Intn(9))
	fmt.Printf("[math] math.Sqrt(16) = %g \n", math.Sqrt(16))
	fmt.Println("[math] math.Pi =", math.Pi)

	fmt.Println("\nBEGIN add_0() and add_1()")
	fmt.Println("add_0(3, 5) =", add_0(3, 5))
	fmt.Println("add_1(3, 5) =", add_1(3, 5))
	fmt.Println("END add_0() and add_1()")

	fmt.Println("\nBEGIN swap()")
	koga, kandy := swap("Koga", "Kandy")
	fmt.Println("swap(\"Koga\", \"Kandy\") =", koga, kandy)
	fmt.Println("END swap()")

	fmt.Println("\nBEGIN split()")
	x, y := split(17)
	fmt.Println("split(17) =", x, y)
	fmt.Println("END split()")

	fmt.Println("\nBEGIN variables()")
	variables()
	fmt.Println("END variables()")
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

// Go's named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func variables() {
	var i int
	fmt.Println(i, c, python, java) // 0 false false false

	var i_1, j_1 int = 1, 2
	fmt.Println(i_1, j_1) // 1 2

	// Go's shorthand variable declaration
	k := 3
	csharp, ruby, php := true, false, "はい！"
	fmt.Println(k, csharp, ruby, php) // 3 true false はい！
}
