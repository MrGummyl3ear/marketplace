package handler

import (
	"marketplace/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input model.Item
	input.CreatedAt = time.Now()
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
