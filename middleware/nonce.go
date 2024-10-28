package middleware

import (
	"fmt"

	"github.com/EduardValentin/course-platform/util"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Nonce() gin.HandlerFunc {
	return func(c *gin.Context) {
		nonce := utils.GenerateSecureNonce(128)
		c.Header("Content-Security-Policy", fmt.Sprintf("script-src 'nonce-%s'", nonce))

		ctx := templ.WithNonce(c.Request.Context(), nonce)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
