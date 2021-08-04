package client

import (
	"fmt"
	"x-msa-core/grpc/msa_observer"
	"x-msa-core/helper"
)

func NewObserverClient(addr string) (ObserverClient, error) {
	var err error
	client := &observerClient{}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msa_observer.NewObserverClient(client.conn)

	return client, nil
}

func (c *observerClient) Close() { c.close(); c.conn.Close() }

func (c *observerClient) PushFirst(req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error) {
	return c.client.PushFirst(helper.Ctx, req)
}
func (c *observerClient) PushStats(req *msa_observer.RequestPushStats) (*msa_observer.Response, error) {
	return c.client.PushStats(helper.Ctx, req)
}
func (c *observerClient) PushStatus(req *msa_observer.RequestPushStatus) (*msa_observer.Response, error) {
	return c.client.PushStatus(helper.Ctx, req)
}
func (c *observerClient) Who(req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error) {
	return c.client.Who(helper.Ctx, req)
}
func (c *observerClient) RestartService(req *msa_observer.RequestRestartService) (*msa_observer.Response, error) {
	return c.client.RestartService(helper.Ctx, req)
}
