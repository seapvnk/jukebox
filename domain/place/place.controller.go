package place

import (
	"net/http"

    "github.com/gin-gonic/gin"
)

// POST: /places
func PlaceCreateController(c *gin.Context) {
	var dto CreatePlaceDto
	c.Bind(&dto)

	places := CreatePlaceService(dto)

	c.JSON(http.StatusCreated, gin.H{
        "places": places,
        "message": "place created",
    })
}

// GET: /places
func PlaceGetController(c *gin.Context) {
	places := GetAllPlacesService()

	c.JSON(http.StatusOK, gin.H{
        "places": places,
        "message": "places fetched",
    })
}

// GET: /places/:id
func PlaceGetOneController(c *gin.Context) {
	var placeDto GetPlaceDto
	
	if err := c.ShouldBindUri(&placeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"places": []int{},
			"message": err,
		})
		return
	}

	places := GetPlaceService(placeDto.ID)

	if len(places) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"places": places,
			"message": "place not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "places": places,
        "message": "place found",
    })
}