package playground

const CODE_FOLDER string = "playground/temp"
const FILE_NAME string = "user_code.go"
const EXAMPLE_CODE string = `
package main

import (
	"fmt"
)

func main() {
	// receive inputs from stdin
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("scan failed:",err)
	}
	fmt.Printf("a=%d, b=%d\n", a, b)

	// output the answer
	ans := a + b
	fmt.Printf("ans=%d", ans)
}
`

type PlaygroundArgs struct {
	Code   string
	Inputs []string
}
