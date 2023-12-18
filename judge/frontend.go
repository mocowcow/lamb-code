package judge

import (
	"github.com/gin-gonic/gin"
)

func frontendRouting(r *gin.Engine) {
	r.LoadHTMLGlob("./judge/assets/*.html")
	r.Static("code_template", "./judge/assets/code_template")

	r.GET("/index", getIndex)
	r.GET("/problems/:id", getProblemPage)
}

func getIndex(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func getProblemPage(ctx *gin.Context) {
	ctx.HTML(200, "problems.html", nil)
}
