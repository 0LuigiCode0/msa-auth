package grpcpublic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"x-msa-auth/helper"
	"x-msa-auth/store/mongo/model"
	"x-msa-core/grpc/msa_service"
	corehelper "x-msa-core/helper"
	"x-msa-core/service/server"

	goutill "github.com/0LuigiCode0/go-utill"
)

func NewUserServices(service server.ServiceServer) AuthServices {
	return &authServices{ServiceServer: service}
}

func (s *authServices) Auth() Auth {
	if c, err := s.Services().GetFirstByGroup(corehelper.Auth); err != nil {
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
		return nil, fmt.Errorf("%v json: %v", corehelper.KeyErrorParse, err)
	}

	resp, err := u.Call(&msa_service.RequestCall{
		FuncName: helper.AuthGuard,
		Data:     data,
	})
	if err != nil {
		return nil, fmt.Errorf("%v user: %v", corehelper.KeyErrorNotFound, err)
	}
	out := &model.UserModel{}
	if err := goutill.JsonParse(bytes.NewReader(resp.Result), out); err != nil {
		return nil, fmt.Errorf("%v json: %v", corehelper.KeyErrorParse, err)
	}
	return out, nil
}
