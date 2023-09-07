package playground

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

func Run(userCode string, inputs []string) []string {
	if userCode == "" {
		userCode = EXAMPLE_CODE
		inputs = []string{"5 9"}
	}

	// write to temp file
	sourceCodePath := path.Join(CODE_FOLDER, FILE_NAME)
	executablePath := path.Join(CODE_FOLDER, "run.exe")
	os.MkdirAll(CODE_FOLDER, os.ModePerm)
	os.WriteFile(sourceCodePath, []byte(EXAMPLE_CODE), os.ModePerm)

	// complie
	// go build -o playground/temp/user_code.exe ./playground/temp/user_code.go
	cmd := exec.Command("go", "build", "-o", executablePath, sourceCodePath)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("complie failed:", err)
		return []string{err.Error()}
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
