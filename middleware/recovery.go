package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rollbar/rollbar-go"
)

// Recovery - middleware for rollbar error monitoring.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rVal := recover(); rVal != nil {
				log.Print(rVal)
				rollbar.RequestErrorWithStackSkipWithExtrasAndContext(c, rollbar.CRIT, c.Request, fmt.Errorf("%+v", rVal), 4, nil)
				c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
			}
		}()

		c.Next()
	}
}
