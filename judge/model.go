package judge

type submitCodeInput struct {
	ProblemId int
	Code      string
}

type testcase struct {
	Input  string
	Output string
}

var CODE_TEMPLATE map[string]string

func init() {
	initGolangTemplate()
}

func initGolangTemplate() {
	CODE_TEMPLATE["go"] = `
	package main
	
	import (
		"fmt"
	)
	
	func main() {
		//// receive inputs from stdin
		// var a, b int
	
		//// output the answer
		// ans := a + b
		// fmt.Printf("ans=%d", ans)
	}
	`
}
