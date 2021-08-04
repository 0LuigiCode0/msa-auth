package roots_handler

import (
	"net/http"
	corehelper "x-msa-core/helper"

	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
)

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	login := vars["login"]
	pwd := vars["pwd"]

	user, err := h.MongoStore().UserStore().SelectByLogin(login)
	if err != nil {
		logger.Log.Warningf("%v user: %v", corehelper.KeyErrorSave, err)
		h.respError(w, corehelper.ErrorSave, corehelper.KeyErrorSave)
		return
	}

	if user.Password != pwd {
		logger.Log.Warningf("%v password wrong", corehelper.KeyErorrAccessDenied)
		h.respError(w, corehelper.ErorrAccessDeniedParams, corehelper.KeyErorrAccessDenied)
		return
	}

	resp, err := h.Helper().GenerateJwt(user.ID)
	if err != nil {
		logger.Log.Warningf("%v jwt %v", corehelper.KeyErrorGenerate, err)
		h.respError(w, corehelper.ErrorGenerate, corehelper.KeyErrorGenerate)
		return
	}

	h.respOk(w, resp)
}
