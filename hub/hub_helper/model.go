package hub_helper

import (
	"net/http"
	"x-msa-auth/core/database"
	"x-msa-auth/handlers/grpc_handler/grpc_helper"
	"x-msa-auth/handlers/roots_handler/roots_helper"
	"x-msa-auth/helper"

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
