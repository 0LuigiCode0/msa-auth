package grpcpublic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0LuigiCode0/msa-core/grpc/msa_service"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"
	"github.com/0LuigiCode0/msa-core/service/server"

	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/store/mongo/model"

	goutill "github.com/0LuigiCode0/go-utill"
)

func NewUserServices(service server.ServiceServer) AuthServices {
	return &authServices{ServiceServer: service}
}

func (s *authServices) Auth() Auth {
	if c, err := s.Services().GetFirstByGroup(coreHelper.Auth); err != nil {
		return &auth{err: err}
	} else {
		return &auth{ServiceClient: c}
	}
}

func (u *auth) Error() error { return u.err }

func (u *auth) AuthGuard(r *http.Request, roles ...helper.Role) (*model.UserModel, error) {
	if u.err != nil {
		return nil, u.err
	}

	in := &model.RequsetAuthGRPC{
		Jwt:   r.Header.Get("Authorization"),
		Roles: roles,
	}
	data, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("%v json: %v", coreHelper.KeyErrorParse, err)
	}

	resp, err := u.Call(&msa_service.RequestCall{
		FuncName: helper.AuthGuard,
		Data:     data,
	})
	if err != nil {
		return nil, fmt.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
	}
	out := &model.UserModel{}
	if err := goutill.JsonParse(bytes.NewReader(resp.Result), out); err != nil {
		return nil, fmt.Errorf("%v json: %v", coreHelper.KeyErrorParse, err)
	}
	return out, nil
}
