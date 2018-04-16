package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/k8s-study/endpoint-service/database"
	"log"
	"time"
	"strconv"
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
	UserId		   string
}

func GetUsersEndpointsList(c echo.Context) error {

	var v SearchInput
	if err := c.Bind(&v); err != nil {
		log.Fatalln(err)
	}
	if v.MaxResultCount == 0 {
		v.MaxResultCount = 15
	}

	v.UserId = "aaaaa-bbbbb-ccccc"

	totalCount, items, err := database.Endpoint{}.GetUsersAll(c.Request().Context(), v.UserId, v.SkipCount, v.MaxResultCount)
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

	var input EndpointInput

	if err := c.Bind(&input); err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	endpoint, err := input.ToModel()
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	endpoint.UserId = "aaaaa-bbbbb-ccccc"

	_, err = endpoint.Create(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, endpoint)
}

func UpdateUsersEndpoints(c echo.Context) error {

	var input EndpointInput

	if err := c.Bind(&input); err != nil {
		return c.Redirect(http.StatusFound, "/111")
	}

	endpoint, err := input.ToModel()
	if err != nil {
		return c.Redirect(http.StatusFound, "/333")
	}

	id, err := strconv.ParseInt(c.Param("endpointId"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, input)
	}

	endpoint.UserId = "aaaaa-bbbbb-ccccc"

	endpoint.ID = id
	if err := endpoint.Update(c.Request().Context()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, endpoint)

}

func DeleteUsersEndpoints(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("endpointId"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT found endpoint")
	}

	endpoint := &database.Endpoint{
		ID : id,
		UserId : "aaaaa-bbbbb-ccccc",
	}


	if err := endpoint.Delete(c.Request().Context()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, endpoint)

	//return c.String(http.StatusOK, "delete user endpoint")
}
