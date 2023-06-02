package v1

import "github.com/gin-gonic/gin"

func InitRouter(app *gin.Engine) {
    app.POST("/produce", produceHandler)
}