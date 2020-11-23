//+build wireinject

package handler

import (
	"github.com/fukuyama012/go-server-layered-sample/db"
	"github.com/fukuyama012/go-server-layered-sample/domain/repository"
	"github.com/fukuyama012/go-server-layered-sample/infra/mysql"
	"github.com/fukuyama012/go-server-layered-sample/usecase"
	"github.com/google/wire"
)

func initExample(hint string) usecase.Example {
	wire.Build(
		db.NewDB,
		mysql.NewExample,
		usecase.NewExample,
		wire.Bind(new(repository.Example), new(*mysql.Example)),
	)
	return nil
}
