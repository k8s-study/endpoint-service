package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/k8s-study/endpoint-service/docs"
	"github.com/k8s-study/endpoint-service/endpoints"
	"github.com/k8s-study/endpoint-service/users"
	"net/http"
	"github.com/swaggo/echo-swagger"
	"github.com/k8s-study/endpoint-service/database"
	"fmt"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	DB, err:= database.Init()

	if err != nil {
		fmt.Println(err)
	}

	defer DB.Close()

	e.GET("/checkList", endpoints.GetAllEndpoints)
	e.GET("/endpoints", users.GetUsersEndpointsList)
	e.POST("/endpoints/", users.CreateUsersEndpoints)

	e.GET("/endpoints/:endpointId", users.GetUsersOneEndpoint)
	e.PUT("/endpoints/:endpointId", users.UpdateUsersEndpoints)
	e.DELETE("/endpoints/:endpointId", users.DeleteUsersEndpoints)
	//
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}