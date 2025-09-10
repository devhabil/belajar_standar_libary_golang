package main

import (
	"fmt"
	"math"

)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(20, 22))
	fmt.Println(math.Min(20, 22))
}