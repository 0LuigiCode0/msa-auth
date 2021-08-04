package hub

import (
	"fmt"
	"net/http"
	"x-msa-auth/core/database"
	"x-msa-auth/handlers/grpc_handler"
	"x-msa-auth/handlers/grpc_handler/grpc_helper"
	"x-msa-auth/handlers/roots_handler"
	"x-msa-auth/handlers/roots_handler/roots_helper"
	"x-msa-auth/helper"
	"x-msa-auth/hub/hub_helper"

	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
)

const (
	_roots = "roots"
	_wss   = "wss"
	_grpc  = "grpc"
)

type Hub interface {
	GetHandler() http.Handler
	Close()
}

type hub struct {
	database.DB
	helper  hub_helper.Helper
	router  *mux.Router
	handler http.Handler
	config  *helper.Config

	_roots roots_helper.Handler
	_grpc  grpc_helper.Handler
}

func InitHub(db database.DB, conf *helper.Config) (H Hub, err error) {
	hh := &hub{
		DB:     db,
		router: mux.NewRouter(),
		config: conf,
	}
	H = hh
	hh.SetHandler(hh.router)

	hh.helper = hub_helper.InitHelper(hh)

	if err = hh.intiDefault(); err != nil {
		logger.Log.Warningf("initializing default is failed: %v", err)
		err = fmt.Errorf("handler not initializing: %v", err)
		return
	}
	logger.Log.Service("initializing default")

	if v, ok := conf.Handlers[_roots]; ok {
		hh._roots, err = roots_handler.InitHandler(hh, v)
		if err != nil {
			err = fmt.Errorf("handler %q not initializing: %v", _roots, err)
			return
		}
		logger.Log.Servicef("handler %q initializing", _roots)
	} else {
		err = fmt.Errorf("config %q not found", _roots)
		return
	}

	if v, ok := conf.Handlers[_grpc]; ok {
		hh._grpc, err = grpc_handler.InitHandler(hh, v)
		if err != nil {
			err = fmt.Errorf("handler %q not initializing: %v", _grpc, err)
			return
		}
		logger.Log.Servicef("handler %q initializing", _grpc)
	} else {
		err = fmt.Errorf("config %q not found", _grpc)
		return
	}

	logger.Log.Service("handler initializing")
	return
}

func (h *hub) Config() *helper.Config      { return h.config }
func (h *hub) Helper() hub_helper.Helper   { return h.helper }
func (h *hub) Router() *mux.Router         { return h.router }
func (h *hub) GetHandler() http.Handler    { return h.handler }
func (h *hub) SetHandler(hh http.Handler)  { h.handler = hh }
func (h *hub) Roots() roots_helper.Handler { return h._roots }
func (h *hub) Grps() grpc_helper.MSA       { return h._grpc }
func (h *hub) Close() {
	if h._grpc != nil {
		h._grpc.Close()
	}
}

func (h *hub) intiDefault() error {
	// Users collection create

	return nil
}
