package judge

import (
	"fmt"
	"lamb-code/config"
	"strings"

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
	host := fmt.Sprintf("%s:%s",
		config.GetString("service.problem.host"),
		config.GetString("service.problem.port"),
	)
	var data gin.H
	req := client.R().
		SetPathParam("host", host).
		SetPathParam("problemId", problemId).
		SetResult(&data)

	req.Get("http://{host}/problems/{problemId}")

	respStrs := make([]string, 0)
	respStrs = append(respStrs, data["Title"].(string))
	respStrs = append(respStrs, data["Content"].(string))

	ctx.String(200, strings.Join(respStrs, "\n"))
}
