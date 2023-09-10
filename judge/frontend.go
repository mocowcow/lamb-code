package judge

import (
	"fmt"
	"lamb-code/config"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func frontendRouting(r *gin.Engine) {
	r.LoadHTMLGlob("./judge/*.html")
	r.GET("/index/:id", getIndex)
	r.GET("/problems/:id", getProblem)
}

func getIndex(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func getProblem(ctx *gin.Context) {
	client := resty.New()
	problemId := ctx.Param("id")
	data := gin.H{}
	req := client.R().
		SetPathParam("host", config.PROBLEM_ADDR).
		SetPathParam("problemId", problemId).
		SetResult(&data)

	req.Get("http://{host}/problems/{problemId}")

	s := fmt.Sprintf("%s : %s", data["Title"], data["Content"])
	ctx.String(200, s)
}
