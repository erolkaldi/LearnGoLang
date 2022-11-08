package main

import (
	"net/http"

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
func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.Run(":8080")

}
