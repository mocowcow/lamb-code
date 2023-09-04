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
	r.GET("/problems", getProblems)
	r.GET("/problems/:id", getProblem)
	r.Run(config.PROBLEM_ADDR)
}

type problem struct {
	Id      int
	Title   string
	Content string
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
