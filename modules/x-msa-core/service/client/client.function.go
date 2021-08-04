package client

import (
	"fmt"
	"x-msa-core/grpc/msa_service"
	"x-msa-core/helper"
)

func NewServiceClient(addr string, group helper.GroupsType) (ServiceClient, error) {
	var err error
	client := &serviceClient{
		group: group,
	}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msa_service.NewServiceClient(client.conn)

	return client, nil
}

func (c *serviceClient) Close() { c.close(); c.conn.Close() }

func (c *serviceClient) Call(req *msa_service.RequestCall) (*msa_service.ResponseCall, error) {
	return c.client.Call(helper.Ctx, req)
}
func (c *serviceClient) AddService(req *msa_service.RequestAddService) (*msa_service.Response, error) {
	return c.client.AddService(helper.Ctx, req)
}
func (c *serviceClient) DeleteService(req *msa_service.RequestDelService) (*msa_service.Response, error) {
	return c.client.DeleteService(helper.Ctx, req)
}
