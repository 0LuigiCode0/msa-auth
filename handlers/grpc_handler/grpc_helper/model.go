package grpc_helper

import (
	coreHelper "github.com/0LuigiCode0/msa-core/helper"
)

type Handler interface {
	Close()

	AddService(key, addr string, group coreHelper.GroupsType)
	DeleteService(key string, group coreHelper.GroupsType) error
}

type MSA interface {
	AddService(key, addr string, group coreHelper.GroupsType)
	DeleteService(key string, group coreHelper.GroupsType) error
}
