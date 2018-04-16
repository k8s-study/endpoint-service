package endpoints

import (

	"net/http"

	"github.com/labstack/echo"
	"github.com/k8s-study/endpoint-service/database"

)


type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

type SearchInput struct {
	SkipCount      int `query:"skipCount"`
	MaxResultCount int `query:"maxResultCount"`
}

// e.GET("/checkList", getAllEndpoints)
func GetAllEndpoints(c echo.Context) error {

	var v SearchInput
	if err := c.Bind(&v); err != nil {

	}
	if v.MaxResultCount == 0 {
		v.MaxResultCount = 15
	}

	totalCount, items, err := database.Endpoint{}.GetAll(c.Request().Context(), v.SkipCount, v.MaxResultCount)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"TotalCount":     totalCount,
		"MaxResultCount": v.MaxResultCount,
		"Items":          items,
	})
}

