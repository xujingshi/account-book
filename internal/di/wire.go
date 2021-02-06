// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"account-book/internal/dao"
	"account-book/internal/service"
	"account-book/internal/server/grpc"
	"account-book/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
