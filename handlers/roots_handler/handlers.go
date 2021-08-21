package roots_handler

import (
	"encoding/json"
	"net/http"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/msa-auth/handlers/roots_handler/roots_helper"
	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/hub/hub_helper"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hub_helper.HelperForHandler
}

func InitHandler(hub hub_helper.HelperForHandler, conf *helper.HandlerConfig) (H roots_helper.Handler, err error) {
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
