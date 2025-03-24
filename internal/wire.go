//go:build wireinject
// +build wireinject

package internal

import (
	"database/sql"
	"testDBMock/internal/repository"
	"testDBMock/internal/service"

	"github.com/google/wire"
)

func InitializeUserService(db *sql.DB) (service.UserService, error) {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
	)
	return nil, nil
}
