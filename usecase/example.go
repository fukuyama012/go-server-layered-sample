package usecase

import (
	"github.com/fukuyama012/go-server-layered-sample/domain/repository"
	"github.com/fukuyama012/go-server-layered-sample/response"
)

type (
	// Example usecase interface
	Example interface {
		Get(id int64) (*response.Example, error)
	}

	// ExampleImpl struct
	ExampleImpl struct {
		exampleRepo repository.Example
	}
)

// NewExample returns example instance
func NewExample(exampleRepo repository.Example) Example {
	return &ExampleImpl{
		exampleRepo: exampleRepo,
	}
}

// Get get example response
func (u *ExampleImpl) Get(id int64) (*response.Example, error) {
	m, err := u.exampleRepo.Load(id)
	if err != nil {
		return nil, err
	}
	res := response.NewExample().ConvertFromModel(m)
	return res, nil
}
