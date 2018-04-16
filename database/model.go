package database

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ “github.com/go-sql-driver/mysql”

	"github.com/labstack/echo"
	"github.com/go-xorm/xorm"
	"context"
	"time"
)

var ContextDBName = "DB"

type Endpoint struct {
	ID        int       `json:"id" xorm:"'id' pk autoincr"`
	UserId    string    `json:"userId"`
	Enable    bool      `json:"enable"`
	Url       string    `json:"url"`
	Port      int       `json:"port" xorm:"default 80"`
	Interval  int       `json:"interval" xorm:"default 1"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Endpoint) Create(ctx context.Context) (int, error) {

	DB(ctx).Insert(d)

	return d.ID, nil
}

func (Endpoint) GetAll(ctx context.Context, offset, limit int) (totalCount int64, items []Endpoint, err error) {
	queryBuilder := func() xorm.Interface {
		q := DB(ctx)
		return q
	}

	errc := make(chan error)
	go func() {
		v, err := queryBuilder().Count(&Endpoint{})
		if err != nil {
			errc <- err
			return
		}
		totalCount = v
		errc <- nil

	}()

	go func() {
		if err := queryBuilder().Limit(limit, offset).Find(&items); err != nil {
			errc <- err
			return
		}
		errc <- nil
	}()

	if err := <-errc; err != nil {
		return 0, nil, err
	}
	if err := <-errc; err != nil {
		return 0, nil, err
	}
	return
}

func DB(ctx context.Context) xorm.Interface {
	v := ctx.Value(ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	if db, ok := v.(*xorm.Engine); ok {
		return db
	}
	panic("DB is not exist")
}

func DbContext(db *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session := db.NewSession()
			defer session.Close()

			req := c.Request()
			c.SetRequest(req.WithContext(
				context.WithValue(
					req.Context(),
					ContextDBName,
					session,
				),
			))

			switch req.Method {
			case "POST", "PUT", "DELETE":
				if err := session.Begin(); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
				if err := next(c); err != nil {
					session.Rollback()
					return echo.NewHTTPError(500, err.Error())
				}
				if c.Response().Status >= 500 {
					session.Rollback()
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
			default:
				if err := next(c); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
			}

			return nil
		}
	}
}
