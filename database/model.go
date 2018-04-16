package database

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"
)

type Endpoint struct {
	ID        int64     `json:"id" xorm:"'id' pk autoincr"`
	UserId    string    `json:"userId" xorm:"'userId'"`
	Enable    bool      `json:"enable"`
	Url       string    `json:"url"`
	Port      int       `json:"port" xorm:"default 80"`
	Interval  int       `json:"interval" xorm:"default 1"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Endpoint) Create(ctx context.Context) (int64, error) {

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

func (Endpoint) GetUsersAll(ctx context.Context, userId string, offset, limit int) (totalCount int64, items []Endpoint, err error) {
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
		if err := queryBuilder().Where("userId =? ", userId).Limit(limit, offset).Find(&items); err != nil {
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

func (Endpoint) GetById(ctx context.Context, id int64) (*Endpoint, error) {
	db := ctx.Value(ContextDBName).(*xorm.Session)
	var v Endpoint
	if has, err := db.ID(id).Get(&v); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return &v, nil
}

func (d *Endpoint) Update(ctx context.Context) (err error) {
	db := ctx.Value(ContextDBName).(*xorm.Session)
	_, err = db.ID(d.ID).Where("userId =?", d.UserId).Update(d)
	return
}

func (d Endpoint) Delete(ctx context.Context) (err error) {
	_, err = DB(ctx).ID(d.ID).Where("userId =?", d.UserId).Delete(&Endpoint{})
	return
}
