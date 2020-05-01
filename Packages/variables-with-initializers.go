package main

import "fmt"

// 変数代入
var i, j int = 1, 2

func main() {
	// 型を省略
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
