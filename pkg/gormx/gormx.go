package gormx

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Client interface {
		Table(name string, args ...interface{}) Client
		Select(query interface{}, args ...interface{}) Client
		Find(dest interface{}, conds ...interface{}) Client
		Error() error
	}
)

type (
	gormx struct {
		client *gorm.DB
	}
)

func wrap(db *gorm.DB) *gormx {
	return &gormx{client: db}
}

func NewClient(connString string) (Client, error) {
	db, err := gorm.Open(postgres.Open(connString))
	return wrap(db), err
}

func (c *gormx) Table(name string, args ...interface{}) Client {
	return wrap(c.client.Table(name, args...))
}

func (c *gormx) Select(query interface{}, args ...interface{}) Client {
	return wrap(c.client.Select(query, args...))
}

func (c *gormx) Find(dest interface{}, conds ...interface{}) Client {
	return wrap(c.client.Find(dest, conds...))
}

func (c *gormx) Error() error {
	return c.client.Error
}
