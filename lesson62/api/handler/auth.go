package handler

import (
	"fmt"
	"net/http"
	"students/api/token"
	"students/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
	var user models.Register
	role := ctx.GetHeader("role")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}

	resp, err := h.Authentification.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(role)
	_, eror := h.Enforcer.AddGroupingPolicy(user.Username, role)
	if eror != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error adding policy",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(ctx *gin.Context) {
	var user models.LoginRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}

	resp, err := h.Authentification.Login(models.Register{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if resp != user.Password {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Errors{
			Error: "passwor xato",
		})
		return
	}

	newToken, eror := token.GenerateAccessJWT(&models.LoginRequest{
		Username: user.Username,
		Role:     user.Role,
	})
	if eror != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: eror.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Login{
		AccessToken: newToken,
	})
}
