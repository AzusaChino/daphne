package router

import (
	v1 "github.com/azusachino/daphne/router/v1"
	v2 "github.com/azusachino/daphne/router/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	v1.InitRouter(app)
	v2.InitRouter(app)
}
