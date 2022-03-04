package main

import (
	"fmt"

	"github.com/DevBadi/restgo/controller"

	"github.com/gin-gonic/gin"
	"github.com/pilinux/gorest/lib/middleware"
)

func main() {
	router, err := SetupRouter()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = router.Run(":" + "1200")
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SetupRouter ...
func SetupRouter() (*gin.Engine, error) {
	router := gin.Default()

	router.Use(middleware.CORS())

	// API:v1.0
	v1 := router.Group("/api/v1/")
	{
		v1.GET("greetings", controller.Greetings)
	}

	return router, nil
}
