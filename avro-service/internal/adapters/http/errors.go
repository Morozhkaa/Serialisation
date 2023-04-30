package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func (a *Adapter) ErrorHandler(ctx *gin.Context, err error) {
	l := zapctx.Logger(ctx)
	l.Sugar().Errorf("request failed: %s", err.Error())

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
