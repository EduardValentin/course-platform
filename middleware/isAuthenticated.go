package middleware

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		c := context.WithValue(ctx.Request.Context(), "authenticated", true)

		ctx.Request = ctx.Request.WithContext(c)

		ctx.Next()
	}
}
