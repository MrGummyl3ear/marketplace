package handler

import (
	"fmt"
	"marketplace/pkg/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxPriceLimit = 99999
	pageSizeLimit = 20
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
	params, err := ParseQuery(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	//fmt.Printf("%+v\n", params)

	items, err := h.services.GetAllItems(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
	/*
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"posts": items,
		})
	*/
}

func ParseQuery(c *gin.Context) (model.QueryParam, error) {
	var params model.QueryParam
	var err error

	params.Order = c.DefaultQuery("order", "date")
	if (params.Order != "date") && (params.Order != "price") {
		return params, fmt.Errorf("incorrect format of order")
	}
	params.Sort = c.DefaultQuery("sort", "asc")
	if (params.Sort != "asc") && (params.Sort != "desc") {
		return params, fmt.Errorf("incorrect format of sort")
	}

	params.MaxPrice, err = strconv.ParseUint(c.DefaultQuery("maxPrice", "99999"), 10, 64)
	if (err != nil) || (params.MaxPrice >= maxPriceLimit) {
		return params, fmt.Errorf("incorrect format of max price")
	}
	params.MinPrice, err = strconv.ParseUint(c.DefaultQuery("minPrice", "99999"), 10, 64)
	if (err != nil) || (params.MinPrice >= params.MaxPrice) {
		return params, fmt.Errorf("incorrect format of min price")
	}
	params.Page, err = strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		return params, err
	}
	if params.Page == 0 {
		params.Page = 1
	}
	params.PageSize, err = strconv.ParseUint(c.DefaultQuery("pageSize", "10"), 10, 64)
	if (err != nil) || (params.PageSize >= pageSizeLimit) || (params.PageSize == 0) {
		return params, err
	}
	if params.PageSize == 0 {
		params.PageSize = 20
	}

	return params, nil
}
