package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	ID   int    `json:id`
	NAME string `json:name`
}

var info = []data{
	{ID: 1, NAME: "Madhavi asam"},
	{ID: 2, NAME: "Madhavi Reddy"},
}

func getData(c *gin.Context) {
	fmt.Println(c)
	c.IndentedJSON(http.StatusOK, info)
}

func addData(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/getData", getData)
	router.Run("localhost:8000")
}
