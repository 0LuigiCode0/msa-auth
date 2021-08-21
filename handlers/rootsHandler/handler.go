package rootsHandler

import (
	"encoding/json"
	"net/http"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/msa-auth/handlers/rootsHandler/rootsHelper"
	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/hub/hubHelper"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hubHelper.HelperForHandler
}

func InitHandler(hub hubHelper.HelperForHandler, conf *helper.HandlerConfig) (H rootsHelper.Handler, err error) {
	h := &handler{HelperForHandler: hub}
	H = h

	hUser := h.Router().PathPrefix("/core").Subrouter()
	hUser.HandleFunc("/auth", h.auth).Queries("login", "{login}", "pwd", "{pwd}")

	h.SetHandler(applyCORS(h.Router()))
	return
}

func (h *handler) respOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &coreHelper.ResponseModel{
		Success: true,
		Result:  data,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		logger.Log.Warningf(coreHelper.KeyErrorParse+": josn: %v", err)
		h.respError(w, coreHelper.ErrorParse, coreHelper.KeyErrorParse+": josn")
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		logger.Log.Warningf(coreHelper.KeyErrorWrite+": response: %v", err)
		h.respError(w, coreHelper.ErrorWrite, coreHelper.KeyErrorWrite+": response")
		return
	}
}

func (h *handler) respError(w http.ResponseWriter, code coreHelper.ErrCode, msg string) {
	w.Header().Set("Content-Type", "application/json")
	resp := &coreHelper.ResponseModel{
		Success: false,
		Result: &coreHelper.ResponseError{
			Code: code,
			Msg:  msg,
		},
	}
	buf, _ := json.Marshal(resp)
	_, err := w.Write(buf)
	if err != nil {
		logger.Log.Warningf(coreHelper.KeyErrorWrite+": response: %v", err)
	}
}
