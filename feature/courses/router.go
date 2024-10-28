package courses

import (
	"github.com/EduardValentin/course-platform/feature/auth/authenticator"
	programmingfundamentals "github.com/EduardValentin/course-platform/feature/courses/programming-fundamentals"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, authenticator *authenticator.Authenticator) {

	router.GET("/programming-fundamentals", programmingfundamentals.Index())
}
