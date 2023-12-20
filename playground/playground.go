package playground

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

var langMap = map[string]languageStrategy{}

func init() {
	langMap["invalid"] = invalid{}
	langMap["go"] = golang{}
}

type languageStrategy interface {
	Run(userCode string, inputs []string) []string
}

type invalid struct {
}

func (invalid) Run(userCode string, inputs []string) []string {
	return []string{"Invalid language"}
}

type golang struct {
}

func (golang) Run(userCode string, inputs []string) []string {
	// write to temp file
	sourceCodePath := path.Join(CODE_FOLDER, "user_code.go")
	executablePath := path.Join(CODE_FOLDER, "run.exe")
	os.MkdirAll(CODE_FOLDER, os.ModePerm)
	os.WriteFile(sourceCodePath, []byte(userCode), os.ModePerm)

	// complie
	// go build -o playground/temp/user_code.exe ./playground/temp/user_code.go
	cmd := exec.Command("go", "build", "-o", executablePath, sourceCodePath)
	b, err := cmd.CombinedOutput()
	if err != nil {
		s := string(b)
		fmt.Println("complie failed:\n", s)
		return strings.Split(s, "\n")
	}

	// run executable
	run := exec.Command(executablePath)
	in, _ := run.StdinPipe()
	out, _ := run.StdoutPipe()
	run.Start()

	// input
	for _, s := range inputs {
		fmt.Println("input: ", s)
		_, err := io.WriteString(in, s)
		if err != nil {
			fmt.Println("input failed:", err)
		}
	}
	in.Close()

	// output
	ouputs := make([]string, 0)
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		aLine := scanner.Text()
		fmt.Println("output:", aLine)
		ouputs = append(ouputs, aLine)
	}

	return ouputs
}
