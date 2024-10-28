package auth

import (
	"github.com/EduardValentin/course-platform/feature/auth/authenticator"
	"github.com/EduardValentin/course-platform/feature/auth/handler/callback"
	"github.com/EduardValentin/course-platform/feature/auth/handler/login"
	"github.com/EduardValentin/course-platform/feature/auth/handler/logout"
	"github.com/EduardValentin/course-platform/feature/auth/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, authenticator *authenticator.Authenticator) {

	router.GET("/login", login.Handler(authenticator))
	router.GET("/callback", callback.Handler(authenticator))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)

}
