package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedrichaase/shorty/internal/database"
	"github.com/cedrichaase/shorty/internal/generator"
	"github.com/cedrichaase/shorty/internal/helper"

	"github.com/gin-gonic/gin"
)

type CreateShortcutRequestBody struct {
	Url    string             `json:"url" binding:"required"`
	Format generator.AlgoName `json:"format,omitempty"`
}

func CreateShortcutHandler(c *gin.Context) {
	var body CreateShortcutRequestBody

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing JSON"})
		return
	}

	shortcut, error := generator.Generate(body.Format)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}

	database.AddShortcut(shortcut, body.Url)
	c.JSON(http.StatusOK, gin.H{"location": fmt.Sprintf("/%v", shortcut)})
}

func AccessShortcutHandler(c *gin.Context) {
	var shortcut = c.Param("shortcut")

	var url = database.FindUrlByShortcut(shortcut)

	if len(url) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No URL found for given shortcut"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func main() {
	router := gin.Default()

	router.POST("/", CreateShortcutHandler)
	router.GET("/:shortcut", AccessShortcutHandler)

	var port = helper.GetEnv("HTTP_PORT", "8080")
	router.Run(fmt.Sprint(":", port))
}
