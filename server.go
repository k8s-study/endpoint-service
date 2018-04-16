package main

import (

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/go-xorm/xorm"

	_ "github.com/k8s-study/endpoint-service/docs"
	"github.com/k8s-study/endpoint-service/database"
	"github.com/k8s-study/endpoint-service/endpoints"
	"github.com/k8s-study/endpoint-service/users"

	"github.com/swaggo/echo-swagger"

)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := xorm.NewEngine("sqlite3", "./gorm.db")

	err = db.Sync(new(database.Endpoint))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	e.Use(database.DbContext(db))

	e.GET("/checkList", endpoints.GetAllEndpoints)

	e.GET("/endpoints", users.GetUsersEndpointsList)
	e.GET("/endpoints/", users.GetUsersEndpointsList)
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