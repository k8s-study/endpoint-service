package main

import (

	"log"
	"net/http"
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/swaggo/echo-swagger"
	_ "github.com/k8s-study/endpoint-service/docs"

	"github.com/k8s-study/endpoint-service/database"
	"github.com/k8s-study/endpoint-service/endpoints"
	"github.com/k8s-study/endpoint-service/users"


)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	driver := "mysql"

	db, err := xorm.NewEngine(driver, connection)

	err = db.Sync(new(database.Endpoint))

	if err != nil {
		log.Fatalln(err)
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