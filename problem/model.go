package problem

type problem struct {
	Id      int
	Title   string
	Content string
}

type testcase struct {
	Id        int
	ProblemId int
	Input     string
	Output    string
}
