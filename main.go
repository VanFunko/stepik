package main

import (
	"fmt"
)

func main() {
	var n, res int
	fmt.Scan(&n)
	res = (n-1)%9 + 1

	fmt.Println(res)
	fmt.Println("nu nah")
}
