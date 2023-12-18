package judge

type submitCodeInput struct {
	ProblemId int
	Code      string
}

type testcase struct {
	Input  string
	Output string
}

type problem struct {
	Id      int
	Title   string
	Content string
}
