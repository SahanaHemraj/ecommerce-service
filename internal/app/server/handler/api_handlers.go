package handler

import (
	r "ecommerce-service/internal/app/server/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// actions actions
	limit  int32
	offset int32
}

type actions interface {
}

func NewHandler(limit, offset int32) *Handler {
	return &Handler{
		// actions: actions,
		limit:  limit,
		offset: offset,
	}
}

const (
	statusPath   = "/status"
	ordersPath   = "/v1beta1/orders"
	productsPath = "/v1beta1/products"
)

func (h *Handler) status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (h *Handler) Routes() r.Routes {
	return []r.Route{
		{
			Name:        "Status",
			Method:      http.MethodGet,
			Pattern:     statusPath,
			HandlerFunc: h.status,
		},
		{
			Name:        "Get orders",
			Method:      http.MethodGet,
			Pattern:     ordersPath,
			HandlerFunc: h.getOrders,
		},
		{
			Name:        "Create orders",
			Method:      http.MethodPost,
			Pattern:     ordersPath,
			HandlerFunc: h.createOrders,
		},
		{
			Name:        "Delete orders",
			Method:      ordersPath,
			Pattern:     ordersPath,
			HandlerFunc: h.deleteOrders,
		},
	}
}
