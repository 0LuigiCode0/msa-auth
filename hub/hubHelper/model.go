package hubHelper

import (
	"net/http"

	"github.com/0LuigiCode0/msa-auth/core/database"
	"github.com/0LuigiCode0/msa-auth/handlers/grpcHandler/grpcHelper"
	"github.com/0LuigiCode0/msa-auth/handlers/rootsHandler/rootsHelper"
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
	Grps() grpcHelper.MSA
}

type HandlerForHelper interface {
	database.DBForHandler
	Roots() rootsHelper.Handler
	Grps() grpcHelper.MSA
	Config() *helper.Config
}

type help struct {
	HandlerForHelper
}
