package grpc_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0LuigiCode0/msa-core/grpc/msa_service"

	"github.com/0LuigiCode0/msa-auth/helper"
	"github.com/0LuigiCode0/msa-auth/store/mongo/model"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
	"github.com/dgrijalva/jwt-go"
)

func (h *handler) call(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error) {
	var out interface{}
	var err error

	switch req.FuncName {
	case helper.AuthGuard:
		out, err = h.authGuard(req)
	default:
		logger.Log.Warningf("%v func -> %v", coreHelper.KeyErrorNotFound, req.FuncName)
		return nil, fmt.Errorf("%v func -> %q", coreHelper.KeyErrorNotFound, req.FuncName)
	}
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(out)
	if err != nil {
		logger.Log.Warningf("%v json: %v", coreHelper.KeyErrorParse, err)
		return nil, err
	}
	return &msa_service.ResponseCall{Result: data}, nil
}

func (h *handler) authGuard(req *msa_service.RequestCall) (*model.UserModel, error) {
	in := &model.RequsetAuthGRPC{}
	if err := goutill.JsonParse(bytes.NewReader(req.Data), in); err != nil {
		logger.Log.Warningf("%v json: %v", coreHelper.KeyErrorParse, err)
		return nil, fmt.Errorf("%v json: %v", coreHelper.KeyErrorParse, err)
	}

	if !strings.HasPrefix(in.Jwt, "Bearer ") {
		logger.Log.Warningf("%v jwt", coreHelper.KeyErrorInvalidParams)
		return nil, fmt.Errorf("%v jwt", coreHelper.KeyErrorInvalidParams)
	}
	token, err := jwt.ParseWithClaims(strings.TrimPrefix(in.Jwt, "Bearer "), &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method == jwt.SigningMethodHS256 {
			err := token.Claims.Valid()
			if err == nil {
				return []byte(helper.Secret), nil
			}
			return nil, err
		}
		return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
	})
	if err != nil {
		logger.Log.Warning("%v jwt: %v", coreHelper.KeyErrorParse, err)
		return nil, err
	}
	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
		user, err := h.MongoStore().UserStore().SelectByID(claims.ID)
		if err != nil {
			logger.Log.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
			return nil, err
		}
		if in.Roles != nil {
			for _, r := range in.Roles {
				if user.Role == r {
					return user, nil
				}
			}
			logger.Log.Warning("wrong role")
			return nil, fmt.Errorf("wrong role")
		}
		return user, nil
	}

	logger.Log.Warningf("%v claims", coreHelper.KeyErrorInvalidParams)
	return nil, fmt.Errorf("%v claims", coreHelper.KeyErrorInvalidParams)
}
