package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	place "github.com/seapvnk/jukebox/domain/place"
	person "github.com/seapvnk/jukebox/domain/person"
)

func HandleHttp() {
	route := gin.Default()
	
	// places
	route.GET("/places", place.PlaceGetController)
	route.GET("/places/:id", place.PlaceGetOneController)
	route.POST("/places", place.PlaceCreateController)

	// people
	route.GET("/people", person.PersonGetController)
	route.GET("/people/:id", person.PersonGetOneController)
	route.POST("/people", person.PersonCreateController)

	// start
	fmt.Println("starting server at 8080")
	route.Run(":8080")
}