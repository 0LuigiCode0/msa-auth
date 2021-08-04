package grpc_helper

import (
	corehelper "x-msa-core/helper"
)

type Handler interface {
	Close()

	AddService(key, addr string, group corehelper.GroupsType)
	DeleteService(key string, group corehelper.GroupsType) error
}

type MSA interface {
	AddService(key, addr string, group corehelper.GroupsType)
	DeleteService(key string, group corehelper.GroupsType) error
}