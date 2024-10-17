package middleware

import (
	"awesomeProject/internal/app/auth"
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Request.Header["Authorization"]
		if !ok || value[0] == "" {
			c.JSON(http.StatusUnauthorized, pkg.BuildResponse(constant.Unauthorized, "Unauthorized", pkg.Null()))
			c.Abort()
			return
		}
		token := value[0]
		claims, err := auth.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, pkg.BuildResponse(constant.Unauthorized, "Unauthorized", pkg.Null()))
			c.Abort()
			return
		}

		role := claims.Role
		c.Set("role", role)

		c.Next()
	}
}
