package problem

import (
	"lamb-code/config"
	"lamb-code/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func RunServer() {
	db, _ = utils.GetDB("problem")

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/problems", getProblems)
	api.GET("/problems/:id", getProblem)
	api.GET("/problems/:id/testcases", getTestcases)

	addr := ":" + config.GetString("service.problem.port")
	r.Run(addr)

}

func getProblems(ctx *gin.Context) {
	var problems []problem
	db.Find(&problems)

	ctx.JSON(
		http.StatusOK,
		problems,
	)
}

func getProblem(ctx *gin.Context) {
	id := ctx.Param("id")
	var problem problem
	db.Find(&problem, id)

	ctx.JSON(
		http.StatusOK,
		problem,
	)
}

func getTestcases(ctx *gin.Context) {
	id := ctx.Param("id")
	var testcases []testcase
	db.Where("problem_id = ?", id).Find(&testcases)

	ctx.JSON(
		http.StatusOK,
		testcases,
	)
}
