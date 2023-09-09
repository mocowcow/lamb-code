package judge

import (
	"fmt"
	"lamb-code/config"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func RunServer() {
	r := gin.Default()
	r.POST("/submit", submitCode)

	frontendRouting(r)

	r.Run(config.JUDGE_ADDR)
}

func submitCode(ctx *gin.Context) {
	var input submitCodeInput
	err := ctx.BindJSON(&input)
	if err != nil {
		fmt.Println("invalid input", err)
		ctx.JSON(
			400,
			gin.H{"result": err},
		)
		return
	}
	fmt.Println("receive", input)

	testcases := getTestcases(input.ProblemId)
	fmt.Println("TCs", testcases)

	// do RPC for each TC
	for _, tc := range testcases {
		res, err := playgroudRPC(input.Code, strings.Split(tc.Input, "\n"))

		if err != nil {
			ctx.JSON(
				400,
				gin.H{"result": "Some Error"},
			)
			return
		}
		resStr := strings.Join(res, "\n")
		if resStr != tc.Output {
			ctx.JSON(
				200,
				gin.H{
					"result":   "Wrong Answer",
					"input":    tc.Input,
					"output":   resStr,
					"expected": tc.Output,
				},
			)
			return
		}

	}
	fmt.Println("user 200")
	ctx.JSON(200,
		gin.H{
			"result": "Accepted",
		},
	)
}

func getTestcases(problemId int) []testcase {
	client := resty.New()
	var testcases []testcase

	req := client.R().
		SetPathParam("host", config.PROBLEM_ADDR).
		SetPathParam("problemId", strconv.Itoa(problemId)).
		SetResult(&testcases)

	req.Get("http://{host}/problems/{problemId}/testcases")

	return testcases
}
