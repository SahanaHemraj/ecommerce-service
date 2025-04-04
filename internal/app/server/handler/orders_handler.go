package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getOrders(ctx *gin.Context) {
	log.Println("Got a rquest to get orders")
	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) createOrders(ctx *gin.Context) {
	log.Println("Got a rquest to get orders")
	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) deleteOrders(ctx *gin.Context) {
	log.Println("Got a rquest to get orders")
	ctx.JSON(http.StatusOK, "")
}

