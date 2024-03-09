package playground

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var langMap = map[string]languageStrategy{}

func init() {
	langMap["invalid"] = invalid{}
	langMap["go"] = golang{timeLimit: GO_TIME_LIMIT}
	langMap["python3"] = python3{timeLimit: PY3_TIME_LIMIT}
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
	timeLimit time.Duration
}

func (code golang) Run(userCode string, inputs []string) []string {
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

	results := make([]string, 0)
	// prevent TLE, kill proc after time limit
	time.AfterFunc(code.timeLimit, func() {
		run.Process.Kill()
		fmt.Println("TLE, kill", run.Process.Pid)
		results = []string{"Time Limit Exceed"}
	})

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
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		aLine := scanner.Text()
		fmt.Println("output:", aLine)
		results = append(results, aLine)
	}

	return results
}

type python3 struct {
	timeLimit time.Duration
}

func (code python3) Run(userCode string, inputs []string) []string {
	sourceCodePath := path.Join(CODE_FOLDER, "user_code")
	os.MkdirAll(CODE_FOLDER, os.ModePerm)
	os.WriteFile(sourceCodePath, []byte(userCode), os.ModePerm)

	// run executable
	run := exec.Command("python", sourceCodePath)
	in, _ := run.StdinPipe()
	out, _ := run.StdoutPipe()
	errout, _ := run.StderrPipe()
	run.Start()

	results := make([]string, 0)
	// prevent TLE, kill proc after time limit
	time.AfterFunc(code.timeLimit, func() {
		run.Process.Kill()
		fmt.Println("TLE, kill", run.Process.Pid)
		results = []string{"Time Limit Exceed"}
	})

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
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		aLine := scanner.Text()
		fmt.Println("output:", aLine)
		results = append(results, aLine)
	}

	// error
	scanner = bufio.NewScanner(errout)
	for scanner.Scan() {
		aLine := scanner.Text()
		fmt.Println("error:", aLine)
		results = append(results, aLine)
	}

	return results
}
