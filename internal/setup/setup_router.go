package setup

import (
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/restapi"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.POST("/register", restapi.RegisterHandler)

	return router
}
