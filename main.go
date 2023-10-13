package main

import (
	"example/api/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brow", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/albums", getAlbums)
			eg.GET("/albums/:id", getAlbumByID)
			eg.POST("/albums", postAlbums)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("Server error")
	}
}

// @BasePath /api/v1

// getAlbums godoc
// @Summary getAlbums
// @Schemes
// @Description getAlbum
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} album
// @Router /example/albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums godoc
// @Summary postAlbums
// @Schemes
// @Description postAlbums
// @Tags example
// @Accept json
// @Param album body album true "Album"
// @Produce json
// @Success 201 {object} album
// @Router /example/albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById godoc
// @Summary getAlbumById
// @Schemes
// @Description getAlbumById
// @Tags example
// @Accept json
// @Param id path int true "Id"
// @Produce json
// @Success 200 {album} album
// @Router /example/albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
