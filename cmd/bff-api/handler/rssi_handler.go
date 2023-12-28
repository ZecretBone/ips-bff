package handler

import (
	"github.com/gin-gonic/gin"
)

type RSSIHandler interface {
	Get(ctx *gin.Context)
}

type rssiHandler struct {
}

func ProvideRSSIHandler() RSSIHandler {
	return &rssiHandler{}
}

func (rs *rssiHandler) Get(ctx *gin.Context) {
	ctx.JSON(200, "asd")
}
