package problem

import (
	"lamb-code/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/problems", getProblems)
	r.GET("/problems/:id", getProblem)
	r.Run(config.PROBLEM_ADDR)
}

func getProblems(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"return": "problem list",
		},
	)
}
func getProblem(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"return": "problems " + id,
		},
	)
}
