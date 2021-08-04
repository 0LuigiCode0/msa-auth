package grpcpublic

import (
	"net/http"
	"x-msa-auth/helper"
	"x-msa-auth/store/mongo/model"
	"x-msa-core/service/client"
	"x-msa-core/service/server"
)

type AuthServices interface {
	Auth() Auth
}

type authServices struct {
	server.ServiceServer
}

type Auth interface {
	Error() error

	AuthGuard(r *http.Request, roles ...helper.Role) (*model.UserModel, error)
}

type auth struct {
	client.ServiceClient
	err error
}
