package person

import (
	"net/http"

    "github.com/gin-gonic/gin"
)

// POST: /people
func PersonCreateController(c *gin.Context) {
	var dto CreatePersonDto
	c.Bind(&dto)

	people := CreatePersonService(dto)

	c.JSON(http.StatusCreated, gin.H{
        "people": people,
        "message": "person created",
    })
}

// GET: /people
func PersonGetController(c *gin.Context) {
	people := GetAllPeopleService()

	c.JSON(http.StatusOK, gin.H{
        "people": people,
        "message": "people fetched",
    })
}

// GET: /people/:id
func PersonGetOneController(c *gin.Context) {
	var personDto GetPersonDto
	
	if err := c.ShouldBindUri(&personDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"people": []int{},
			"message": err,
		})
		return
	}

	people := GetPersonService(personDto.ID)

	if len(people) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"people": people,
			"message": "person not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "people": people,
        "message": "person found",
    })
}