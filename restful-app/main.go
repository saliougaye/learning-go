package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Piove Ancora", Artist: "Silent Bob", Price: 20.00},
	{ID: "2", Title: "Dawn FM", Artist: "The Weeknd", Price: 50.50},
	{ID: "3", Title: "My Turn", Artist: "Lil Baby", Price: 30.00},
	{ID: "4", Title: "X2", Artist: "Sick Luke", Price: 15.00},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func insertAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getSpecificAlbum(c *gin.Context) {

	var id = c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getSpecificAlbum)
	router.POST("/albums", insertAlbum)

	router.Run("localhost:8080")
}
