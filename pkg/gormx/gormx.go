package gormx

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Client interface {
		Table(name string, args ...interface{}) Client
		Select(query interface{}, args ...interface{}) Client
		Where(query interface{}, args ...interface{}) Client
		Count(count *int64) Client
		Find(dest interface{}, conds ...interface{}) Client
		Create(value interface{}) Client
		Updates(value interface{}) Client
		Begin(opts ...*sql.TxOptions) Client
		Commit() Client
		Rollback() Client
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

func (c *gormx) Where(query interface{}, args ...interface{}) Client {
	return wrap(c.client.Where(query, args...))
}

func (c *gormx) Count(count *int64) Client {
	return wrap(c.client.Count(count))
}

func (c *gormx) Find(dest interface{}, conds ...interface{}) Client {
	return wrap(c.client.Find(dest, conds...))
}

func (c *gormx) Create(value interface{}) Client {
	return wrap(c.client.Create(value))
}

func (c *gormx) Updates(value interface{}) Client {
	return wrap(c.client.Updates(value))
}

func (c *gormx) Begin(opts ...*sql.TxOptions) Client {
	return wrap(c.client.Begin(opts...))
}

func (c *gormx) Commit() Client {
	return wrap(c.client.Commit())
}

func (c *gormx) Rollback() Client {
	return wrap(c.client.Rollback())
}

func (c *gormx) Error() error {
	return c.client.Error
}
