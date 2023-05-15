package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	place "github.com/seapvnk/jukebox/domain/place"
)

func HandleHttp() {
	route := gin.Default()
	
	// places
	route.GET("/places", place.PlaceGetController)
	route.GET("/places/:id", place.PlaceGetOneController)
	route.POST("/places", place.PlaceCreateController)

	// start
	fmt.Println("starting server at 8080")
	route.Run(":8080")
}