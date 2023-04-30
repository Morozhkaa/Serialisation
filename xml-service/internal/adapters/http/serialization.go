package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Adapter) getResult(ctx *gin.Context) {
	result, err := a.app.GetResult(ctx)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}
