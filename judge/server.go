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

	r.Run(config.JUDGE_PORT)
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
			ctx.String(400, "Some Error")
			return
		}
		resStr := strings.Join(res, "\n")
		if resStr != tc.Output {
			respStrs := make([]string, 0)
			respStrs = append(respStrs, "Wrong Answer")
			respStrs = append(respStrs, "Input :")
			respStrs = append(respStrs, tc.Input)
			respStrs = append(respStrs, "Output :")
			respStrs = append(respStrs, resStr)
			respStrs = append(respStrs, "Expected :")
			respStrs = append(respStrs, tc.Output)
			ctx.String(200, strings.Join(respStrs, "\n"))
			return
		}
	}

	ctx.String(200, "Accepted")
}

func getTestcases(problemId int) []testcase {
	client := resty.New()
	var testcases []testcase

	req := client.R().
		SetPathParam("host", config.PROBLEM_HOST).
		SetPathParam("problemId", strconv.Itoa(problemId)).
		SetResult(&testcases)

	req.Get("http://{host}/problems/{problemId}/testcases")

	return testcases
}
