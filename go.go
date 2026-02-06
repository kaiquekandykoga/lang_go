package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("ありがとう！")
	fmt.Println("time package = ", time.Now())
	fmt.Println("math/rand package = ", rand.Intn(9))
	fmt.Printf("math package = %g \n", math.Sqrt(16))
}
