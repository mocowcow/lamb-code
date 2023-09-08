package judge

import (
	"fmt"
	"lamb-code/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func RunServer() {
	r := gin.Default()
	r.POST("/submit", submitCode)
	r.Run(config.JUDGE_ADDR)
}

func submitCode(ctx *gin.Context) {
	var input submitCodeInput
	ctx.BindJSON(&input)
	fmt.Println("receive", input)

	testcases := getTestcases(input.ProblemId)
	fmt.Println("TCs", testcases)

	// do RPC for each TC
	for _, tc := range testcases {
		playgroudRPC(tc)
	}

	ctx.JSON(200,
		nil,
	)
}

func getTestcases(problemId int) []testcase {
	client := resty.New()
	var testcases []testcase

	req := client.R().
		SetPathParam("host", config.PROBLEM_ADDR).
		SetPathParam("problemId", strconv.Itoa(1)).
		SetResult(&testcases)

	req.Get("http://{host}/problems/{problemId}/testcases")

	return testcases
}

func playgroudRPC(tc testcase) {
	fmt.Printf("calling RPC for TC %+v\n", tc)
}
