package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func goDotEnvVariable(key string) string {

	// // load .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	return os.Getenv(key)
}

var opts = pg.Options{
	Addr:     goDotEnvVariable("POSTGRES_ADDR"),
	User:     goDotEnvVariable("POSTGRES_USER"),
	Password: goDotEnvVariable("POSTGRES_PASSWORD"),
	Database: goDotEnvVariable("POSTGRES_DATABASE"),
}

var db *pg.DB

func getAlbumsHandler(c *gin.Context) {

	albums := selectAllAlbums(db)

	if albums == nil {
		albums = []Album{}
	}

	c.IndentedJSON(http.StatusOK, albums)

}

func postAlbumHandler(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return
	}

	err := insertAlbum(db, newAlbum)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getSpecificAlbumHandler(c *gin.Context) {

	var id = c.Param("id")

	album, err := selectSingleAlbum(db, id)

	if err != nil {
		return
	}

	if album != nil {
		c.IndentedJSON(http.StatusOK, album)

		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumHandle(c *gin.Context) {

	var id = c.Param("id")

	err := deleteAlbum(db, id)

	if err != nil {
		return
	}

	c.Status(http.StatusNoContent)
}

func editAlbumHandle(c *gin.Context) {
	var albumEdited Album

	if err := c.BindJSON(&albumEdited); err != nil {
		fmt.Println(err)
		return
	}

	id := c.Param("id")
	albumEdited.ID = id

	err := editAlbum(db, albumEdited)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, albumEdited)
}

func main() {

	db = pg.Connect(&opts)

	defer db.Close()

	// err := createSchema(db)

	// if err != nil {
	// 	panic(err)
	// }

	router := gin.Default()

	router.GET("/albums", getAlbumsHandler)
	router.GET("/albums/:id", getSpecificAlbumHandler)
	router.POST("/albums", postAlbumHandler)
	router.PUT("/albums/:id", editAlbumHandle)
	router.DELETE("/albums/:id", deleteAlbumHandle)

	router.Run("localhost:8080")

}
