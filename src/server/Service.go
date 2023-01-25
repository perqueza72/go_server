package go_server

import (
	"log"
	. "server_interfaces"
)

type GoServer struct {
	logger          log.Logger
	user_repository IUserRepository
}

func NewGoServer(logger log.Logger, repository IUserRepository) *GoServer {
	return &GoServer{
		logger:          logger,
		user_repository: repository,
	}
}

func (goServer *GoServer) GetRepository() IUserRepository {
	return goServer.user_repository
}
