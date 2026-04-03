package midlleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			status := http.StatusInternalServerError

			switch c.Errors.Last().Type {
			case gin.ErrorTypeBind:
				status = http.StatusBadRequest
			case gin.ErrorTypePublic:
				status = http.StatusBadRequest
			case gin.ErrorTypePrivate:
				status = http.StatusInternalServerError
			}

			c.JSON(status, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
