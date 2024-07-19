package middleware

import (
	"net/http"
	"students/api/token"
	"students/models"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Errors{
				Error: "",
			})
			return
		}

		claims, err := token.ExtractClaimsAccess(authHeader)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Errors{
				Error: "invalid token",
			})
			return
		}

		ok, err := enforcer.Enforce(claims.Role, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.JSON(http.StatusForbidden, models.Errors{
				Error: err.Error(),
			})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.Errors{
				Error: "user can only get",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
