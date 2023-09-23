package judge

import (
	"github.com/gin-gonic/gin"
)

func frontendRouting(r *gin.Engine) {
	r.LoadHTMLGlob("./judge/*.html")
	r.GET("/index/:id", getIndex)
}

func getIndex(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
