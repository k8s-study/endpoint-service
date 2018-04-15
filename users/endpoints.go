package users

import (
	"github.com/labstack/echo"
	"net/http"
)


type (
	endpoint struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)



func GetUsersEndpointsList(c echo.Context) error {
	return c.String(http.StatusOK, "users endpoints list")
}

func GetUsersOneEndpoint(c echo.Context) error {
	endpointId := c.Param("endpointId")
	return c.String(http.StatusOK, endpointId)
}


func CreateUsersEndpoints(c echo.Context) error {
	return c.String(http.StatusOK, "create user endpoint")
}

func UpdateUsersEndpoints(c echo.Context) error {
	return c.String(http.StatusOK, "update User endpoint")
}

func DeleteUsersEndpoints(c echo.Context) error {
	return c.String(http.StatusOK, "delete user endpoint")
}