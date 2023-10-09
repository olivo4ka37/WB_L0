package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"log"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		log.Println("invalid input body")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.services.Repo.Create(context.TODO(), order.OrderUID, order); err != nil {
		log.Println("cant create order")
		newErrorResponse(c, http.StatusInternalServerError, "cant create order")

		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"uid": order.OrderUID,
	})
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	uid := c.Param("orderuid")
	order, err := h.services.Repo.GetById(context.TODO(), uid)
	if err != nil {
		log.Println("cant get order by id")
		newErrorResponse(c, http.StatusBadRequest, "no order with your id")

		return
	}
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, order)
}
