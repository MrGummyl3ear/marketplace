package handler

import (
	"marketplace/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	username, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "username not found")
		return
	}

	var input model.Item
	input.CreatedAt = time.Now()
	input.Username = username.(string)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Item.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

type getAllItemsResponse struct {
	Data []model.Item `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.services.GetAllItems()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}
