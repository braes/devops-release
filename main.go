package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Maker string `json:"maker"`
}

var vehicles = []Vehicle{
	{ID: "1", Model: "Tesla Model Y", Maker: "Tesla"},
	{ID: "2", Model: "Tesla Model 3", Maker: "Tesla"},
	{ID: "3", Model: "Fiat 500e", Maker: "Fiat"},
	{ID: "4", Model: "Peugeot e-208", Maker: "Peugeot"},
	{ID: "5", Model: "Volkswagen ID.4", Maker: "Volkswagen"},
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthz", getHealthz)
	router.GET("/vehicles", getVehicles)
	router.GET("/vehicles/:id", getVehicleByID)
	router.POST("/vehicles", postVehicles)
	return router
}

func main() {
	router := setupRouter()
	router.Run("0.0.0.0:8080")
}

func getHealthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func getVehicles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicles)
}

func getVehicleByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range vehicles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.Writer.WriteHeader(http.StatusNotFound)
}

func postVehicles(c *gin.Context) {
	var newVehicle Vehicle

	if err := c.BindJSON(&newVehicle); err != nil {
		return
	}

	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}
