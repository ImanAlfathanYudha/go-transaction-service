package cmd

import (
	"fmt"
	"go-transaction-service/common/response"
	"go-transaction-service/config"
	"go-transaction-service/constants"
	"go-transaction-service/controllers"
	"go-transaction-service/middlewares"
	"go-transaction-service/repositories"
	"go-transaction-service/routes"
	"go-transaction-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(c *cobra.Command, args []string) {
		config.Init()
		repository := repositories.NewRepositoryRegistry()
		service := services.NewServiceRegistry(repository)
		controller := controllers.NewControllerRegistry(service)

		router := gin.Default()
		router.Use(middlewares.HandlePanic())

		// NoRoute handler
		router.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, response.Response{
				Status:  constants.Error,
				Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
			})
		})

		// Example root route
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.Response{
				Status:  constants.Success,
				Message: "Welcome to Transaction Service",
			})
		})

		// API group
		group := router.Group("api/v1")
		route := routes.NewRouterRegistry(controller, group)
		route.Serve()

		// Run server
		port := fmt.Sprintf(":%d", config.Config.Port)
		fmt.Printf("Server running on port %s\n", port)
		router.Run(port)
	},
}

func Run() {
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}
