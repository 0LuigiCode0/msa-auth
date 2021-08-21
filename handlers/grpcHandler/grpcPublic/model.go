package grpcpublic

import (
	"net/http"

	"github.com/0LuigiCode0/msa-core/service/client"
	"github.com/0LuigiCode0/msa-core/service/server"

	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/store/mongo/model"
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
