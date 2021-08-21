package hub_helper

import (
	"net/http"

	"github.com/0LuigiCode0/msa-auth/core/database"
	"github.com/0LuigiCode0/msa-auth/handlers/grpc_handler/grpc_helper"
	"github.com/0LuigiCode0/msa-auth/handlers/roots_handler/roots_helper"
	"github.com/0LuigiCode0/msa-auth/helper"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Helper interface {
	GenerateJwt(id primitive.ObjectID) (string, error)
}

type HelperForHandler interface {
	database.DBForHandler
	Helper() Helper
	Config() *helper.Config
	Router() *mux.Router
	SetHandler(hh http.Handler)
	Grps() grpc_helper.MSA
}

type HandlerForHelper interface {
	database.DBForHandler
	Roots() roots_helper.Handler
	Grps() grpc_helper.MSA
	Config() *helper.Config
}

type help struct {
	HandlerForHelper
}
