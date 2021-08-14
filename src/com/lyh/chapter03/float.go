package chapter03

import (
	"fmt"
	"math"
)

func FloatPrint(){
	f := float32(16777216)
	fmt.Println(f == f + 1)

	for x := 0; x < 8; x++{
		fmt.Printf("x = %d ex = %8.3f\n",x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z,  math.IsInf(-1/z,0), math.IsNaN(z/z))

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}