package main

import (
	"C"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"ID"`
	Title  string  `json:"Title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"Price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Heavier things", Artist: "John Mayer", Price: 9.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func main() {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	MODE := os.Getenv("GIN_MODE")
	PORT := os.Getenv("PORT")
	if MODE != "release" {

		fmt.Println("port:", PORT)
		fmt.Println("MODE is: ", MODE)
	}

	gin.SetMode(MODE)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/",

		func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from %v status %v", "Gin", http.StatusOK)
		})

	fmt.Println("server running")
	router.Run()
}
