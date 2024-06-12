package handler

import (
	"my_module/gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	var filter models.Filter

	age, err := strconv.Atoi(c.Query("age"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Xatolik")
		return
	}
	filter.Age = age

	filter.Gender = c.Query("gender")
	filter.Nation = c.Query("nation")
	filter.Field = c.Query("field")

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "xatolik")
		return
	}
	filter.Limit = limit

	if ofs, ok := c.GetQuery("offset"); ok {
		offset, err := strconv.Atoi(ofs)
		if err != nil {
			c.JSON(http.StatusBadRequest, "xatolik")
			return
		}
		filter.Offset = offset
	}
	users, err := h.User.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, "xatolik")
		return
	}

	c.JSON(http.StatusOK, users)
}
