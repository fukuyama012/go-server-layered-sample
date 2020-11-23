package repository

import (
	"github.com/fukuyama012/go-server-layered-sample/domain/model"
)

// Example repository interface
type Example interface {
	Load(id int64) (*model.Example, error)
}
