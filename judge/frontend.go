package judge

import (
	"github.com/gin-gonic/gin"
)

func frontendRouting(r *gin.Engine) {
	r.LoadHTMLGlob("./judge/*.html")
	r.GET("/index/:id", getIndex)
	r.GET("/template/:lang", getTemplate)
}

func getIndex(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func getTemplate(ctx *gin.Context) {
	lang := ctx.Param("lang")
	template := CODE_TEMPLATE[lang]
	ctx.String(200, template)
}
