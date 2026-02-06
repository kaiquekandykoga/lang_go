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
}
