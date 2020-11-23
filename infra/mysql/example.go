package mysql

import (
	"github.com/fukuyama012/go-server-layered-sample/db"
	"github.com/fukuyama012/go-server-layered-sample/domain/model"
)

type (
	// Example struct
	Example struct {
		db *db.DB
	}
)

// NewExample returns example instance
func NewExample(db *db.DB) *Example {
	return &Example{
		db: db,
	}
}

// Load load
func (i *Example) Load(id int64) (*model.Example, error) {
	// 実際はMySQLからSelectする
	m := &model.Example{ID: id, Name: i.db.GetSession()}
	return m, nil
}
