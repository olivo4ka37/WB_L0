package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/olivo4ka37/WB_L0/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("./web/templates/index.html")
	router.Static("/js", "./web/static/js")
	router.Static("/css", "./web/static/css")

	orders := router.Group("/orders")
	{
		orders.GET("", func(c *gin.Context) {
			c.HTML(200, "index.html", map[string]string{"title": "home page"})
		})
		orders.POST("", h.Create)
		orders.GET(":orderuid", h.GetOrderByID)
	}

	return router
}
