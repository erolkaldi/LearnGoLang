package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var community = []person{
	{ID: 1, Name: "Adam", Age: 40},
	{ID: 2, Name: "Joe", Age: 30},
	{ID: 3, Name: "Tom", Age: 20},
	{ID: 4, Name: "Jack", Age: 10},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, community)
}

func createPerson(c *gin.Context) {
	var newPerson person
	if err := c.BindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exists := false
	for _, p := range community {
		if p.Name == newPerson.Name {
			exists = true
			break
		}
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person already exists"})
		return
	}
	newPerson.ID = community[len(community)-1].ID + 1
	community = append(community, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, p := range community {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Person not found"})

}
func deletePersonById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idx := 0
	for i, p := range community {
		if p.ID == id {
			idx = i
		}
	}
	if idx != 0 {
		community = append(community[:idx], community[idx+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Person removed"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Person not found"})

}
func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.GET("/people/:id", getPersonById)
	router.POST("/people", createPerson)
	router.DELETE("/people/:id", deletePersonById)
	router.Run(":8080")

}
