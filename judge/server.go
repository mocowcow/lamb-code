package judge

import (
	"errors"
	"fmt"
	"lamb-code/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func RunServer() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/problems", getProblems)
	api.GET("/problems/:id", getProblem)
	api.POST("/submit", submitCode)

	frontendRouting(r)

	addr := ":" + config.GetString("service.judge.port")
	r.Run(addr)
}

func submitCode(ctx *gin.Context) {
	var input submitCodeInput
	err := ctx.BindJSON(&input)
	if err != nil {
		fmt.Println("invalid input", err)

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"result": err},
		)
		return
	}
	fmt.Println("receive", input)
	fmt.Println("language", input.Lang)

	testcases, err := getTestcases(input.ProblemId)
	if err != nil {
		ctx.String(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	fmt.Println("TCs", testcases)

	// do RPC for each TC
	for _, tc := range testcases {
		res, err := playgroudRPC(input.Lang, input.Code, strings.Split(tc.Input, "\n"))

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

			ctx.String(
				http.StatusOK,
				strings.Join(respStrs, "\n"),
			)
			return
		}
	}

	ctx.String(
		http.StatusOK,
		"Accepted",
	)
}

func getTestcases(problemId int) ([]testcase, error) {
	client := resty.New()
	var testcases []testcase
	host := fmt.Sprintf("%s:%s",
		config.GetString("service.problem.host"),
		config.GetString("service.problem.port"),
	)
	req := client.R().
		SetPathParam("host", host).
		SetPathParam("problemId", strconv.Itoa(problemId)).
		SetResult(&testcases)

	req.Get("http://{host}/api/problems/{problemId}/testcases")

	if len(testcases) == 0 {
		return nil, errors.New("failed to get TCs")
	}

	return testcases, nil
}

func getProblem(ctx *gin.Context) {
	client := resty.New()
	problemId := ctx.Param("id")
	host := fmt.Sprintf("%s:%s",
		config.GetString("service.problem.host"),
		config.GetString("service.problem.port"),
	)
	var data gin.H
	req := client.R().
		SetPathParam("host", host).
		SetPathParam("problemId", problemId).
		SetResult(&data)

	req.Get("http://{host}/api/problems/{problemId}")

	respStrs := make([]string, 0)
	respStrs = append(respStrs, data["Title"].(string))
	respStrs = append(respStrs, data["Content"].(string))

	ctx.String(
		http.StatusOK,
		strings.Join(respStrs, "\n"),
	)
}

func getProblems(ctx *gin.Context) {
	client := resty.New()
	host := fmt.Sprintf("%s:%s",
		config.GetString("service.problem.host"),
		config.GetString("service.problem.port"),
	)
	var problems []problem
	req := client.R().
		SetPathParam("host", host).
		SetResult(&problems)

	req.Get("http://{host}/api/problems")

	ctx.JSON(
		http.StatusOK,
		problems,
	)

}
