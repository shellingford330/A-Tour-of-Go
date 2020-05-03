// +build ignore
package main

import "fmt"

func main() {
	// 変数の初期値を与えない→0
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
