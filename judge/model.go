package judge

type submitCodeInput struct {
	ProblemId int
	Code      string
}

type testcase struct {
	Input  string
	Output string
}

var CODE_TEMPLATE map[string]string = map[string]string{}

func init() {
	initGolangTemplate()
}

func initGolangTemplate() {
	CODE_TEMPLATE["go"] = `package main

import (
	"fmt"
)

func main() {
	//// receive inputs from stdin
	var a, b int
	fmt.Scan(&a, &b)
	
	//// output the answer
	fmt.Println(a+b)
}
`
}

type problem struct {
	Id      int
	Title   string
	Content string
}
