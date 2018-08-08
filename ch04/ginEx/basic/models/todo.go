package models

import (
		"time"
		"github.com/rs/xid"
)

type Base struct {
	ID string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (b *Base)BeforeCreate() (err error) {
	b.ID = xid.New().String()
	return
}

type Todo struct {
	Base
	Title string
	Done bool
}
