package rootsHandler

import (
	"net/http"

	"github.com/0LuigiCode0/msa-auth/helper"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
)

func (h *handler) auth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	login := vars["login"]
	pwd := vars["pwd"]

	user, err := h.MongoStore().UserStore().SelectByLogin(login)
	if err != nil {
		logger.Log.Warningf("%v user: %v", coreHelper.KeyErrorSave, err)
		h.respError(w, coreHelper.ErrorSave, coreHelper.KeyErrorSave)
		return
	}

	pwd = coreHelper.Hash(pwd, helper.Secret)

	if user.Password != pwd {
		logger.Log.Warningf("%v password wrong", coreHelper.KeyErorrAccessDenied)
		h.respError(w, coreHelper.ErorrAccessDeniedParams, coreHelper.KeyErorrAccessDenied)
		return
	}

	resp, err := h.Helper().GenerateJwt(user.ID)
	if err != nil {
		logger.Log.Warningf("%v jwt %v", coreHelper.KeyErrorGenerate, err)
		h.respError(w, coreHelper.ErrorGenerate, coreHelper.KeyErrorGenerate)
		return
	}

	h.respOk(w, resp)
}
