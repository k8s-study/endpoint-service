package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/k8s-study/endpoint-service/database"
	"log"
	"time"
	"fmt"
)

type EndpointInput struct {
	Url      string `json:"url" valid:"required"`
	Port     int    `json:"port"`
	Enable   bool   `json:"enable"`
	Interval int    `json:"interval"`
}

func (d *EndpointInput) ToModel() (*database.Endpoint, error) {

	// port 값이 없을 경우 80번으로 지정
	if d.Port == 0 {
		d.Port = 80
	}

	// interval 값이 없는 경우 1 로 지정
	if d.Interval == 0 {
		d.Interval = 1
	}

	return &database.Endpoint{
		Url:       d.Url,
		Port:      d.Port,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Enable:    d.Enable,
		Interval:  d.Interval,
		UserId:    d.Url,
	}, nil
}

type SearchInput struct {
	SkipCount      int `query:"skipCount"`
	MaxResultCount int `query:"maxResultCount"`
}

func GetUsersEndpointsList(c echo.Context) error {

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

func GetUsersOneEndpoint(c echo.Context) error {
	endpointId := c.Param("endpointId")
	return c.String(http.StatusOK, endpointId)
}

func CreateUsersEndpoints(c echo.Context) error {

	var v EndpointInput

	if err := c.Bind(&v); err != nil {
		return c.Redirect(http.StatusFound, "/111")
	}

	endpoint, err := v.ToModel()
	if err != nil {
		log.Fatalln(err)
		return c.Redirect(http.StatusFound, "/3333")
	}

	endpoint.UserId = "aaaaa-bbbbb-ccccc"

	ep, err := endpoint.Create(c.Request().Context())
	if err != nil {
		return err
	}

	log.Println(ep)

	//return c.JSON(http.StatusOK, ep)

	return c.Redirect(http.StatusFound, fmt.Sprintf("/endpoints/%d", endpoint.ID))
}

func UpdateUsersEndpoints(c echo.Context) error {
	return c.String(http.StatusOK, "update User endpoint")
}

func DeleteUsersEndpoints(c echo.Context) error {
	return c.String(http.StatusOK, "delete user endpoint")
}
