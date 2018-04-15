package endpoints

import (
	"github.com/labstack/echo"
	"net/http"
)


type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

// e.GET("/checkList", getAllEndpoints)
func GetAllEndpoints(c echo.Context) error {

	endpoint := &User{
		Name : "test",
		Email : "email@email.com",
	}

	return c.JSON(http.StatusOK, endpoint)
}

