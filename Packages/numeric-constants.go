// +build ignore
package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	/* int は64-bitの整数を保持できますが、
	それでは足りないことが時々あります。
	そういったときにconstを活用しましょう */
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmr.Println(needFloat(Big))
}
