package router

import (
	"encoding/gob"
	"net/http"
	"os"
	"strconv"

	"github.com/EduardValentin/course-platform/feature/auth"
	"github.com/EduardValentin/course-platform/feature/auth/authenticator"
	"github.com/EduardValentin/course-platform/feature/dashboard"
	"github.com/EduardValentin/course-platform/feature/landing"
	"github.com/EduardValentin/course-platform/middleware"
	"github.com/EduardValentin/course-platform/renderer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func dev() gin.HandlerFunc {
	return func(c *gin.Context) {

		isDevMode, err := strconv.ParseBool(os.Getenv("DEV"))
		if err != nil {
			return
		}
		if !isDevMode {
			return
		}

		c.Header("Cache-Control", "no-store")

		c.Next()
	}
}

//	func WithNonce(t func() templ.Component) func(*gin.Context) {
//		return func(c *gin.Context) {
//			ctx := templ.WithNonce(c.Request.Context(), c.MustGet("Nonce").(string))
//			c.Render(http.StatusOK, renderer.New(ctx, http.StatusOK, t()))
//		}
//	}
var instance *gin.Engine

func Instance(authenticator *authenticator.Authenticator) *gin.Engine {
	if instance != nil {
		return instance
	}

	router := gin.Default()

	router.Use(dev())
	router.Use(middleware.Nonce())

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]any{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	auth.InitRoutes(router, authenticator)

	router.GET("/", func(ctx *gin.Context) {
		renderer := renderer.New(ctx.Request.Context(), http.StatusOK, landing.Hello())
		ctx.Render(http.StatusOK, renderer)
	})
	router.GET("/dashboard", middleware.IsAuthenticated, func(ctx *gin.Context) {
		renderer := renderer.New(ctx.Request.Context(), http.StatusOK, dashboard.Dashboard())
		ctx.Render(http.StatusOK, renderer)
	})

	instance = router

	return router
}
