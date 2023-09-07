package main

import (
	"fmt"
	"lamb-code/playground"
)

var CODE string = `
package main

import (
	"fmt"
)

func main() {
	// receive inputs from stdin
	var a, b int
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		fmt.Println(b * 3)
	}
}
`

var INPUTS []string = []string{"3\n", "99\n", "879\n", "878\n"}

func main() {
	// 3 testcases
	// multiply all elements by 3
	// input [99, 879, 878]
	// ouput [297, 2637, 2634]

	res := playground.Run(CODE, INPUTS)
	println("--------")
	for i, s := range res {
		fmt.Println(i, s)
	}
}
