package server

import (
	"context"
	"fmt"
	"net"
	"x-msa-core/grpc/msa_observer"
	"x-msa-core/monitor/dependents"

	"google.golang.org/grpc"
)

func NewObserverServer(key, addr string) ObserverServer {
	server := &observerServer{
		server:   grpc.NewServer(),
		key:      key,
		addr:     addr,
		monitors: dependents.NewMDM(),
	}
	msa_observer.RegisterObserverServer(server.server, server)

	return server
}

func (s *observerServer) Start() error {
	gListen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("glistener error: %v", err)
	}
	if err := s.server.Serve(gListen); err != nil {
		if err == grpc.ErrServerStopped {
			return nil
		}
		return fmt.Errorf("gserve error: %v", err)
	}

	return nil
}
func (s *observerServer) Close() { s.server.Stop() }

func (s *observerServer) Monitors() dependents.MonitorDependentsManager { return s.monitors }
func (s *observerServer) GetAddr() string                               { return s.addr }
func (s *observerServer) GetKey() string                                { return s.key }

func (s *observerServer) SetPushFirst(f func(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error)) {
	s.pushFirst = f
}

func (s *observerServer) SetPushStats(f func(ctx context.Context, req *msa_observer.RequestPushStats) (*msa_observer.Response, error)) {
	s.pushStats = f
}

func (s *observerServer) SetPushStatus(f func(ctx context.Context, req *msa_observer.RequestPushStatus) (*msa_observer.Response, error)) {
	s.pushStatus = f
}

func (s *observerServer) SetWho(f func(ctx context.Context, req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error)) {
	s.who = f
}

func (s *observerServer) SetRestartService(f func(ctx context.Context, req *msa_observer.RequestRestartService) (*msa_observer.Response, error)) {
	s.restartService = f
}

func (s *observerServer) PushFirst(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error) {
	return s.pushFirst(ctx, req)
}

func (s *observerServer) PushStats(ctx context.Context, req *msa_observer.RequestPushStats) (*msa_observer.Response, error) {
	return s.pushStats(ctx, req)
}

func (s *observerServer) PushStatus(ctx context.Context, req *msa_observer.RequestPushStatus) (*msa_observer.Response, error) {
	return s.pushStatus(ctx, req)
}

func (s *observerServer) Who(ctx context.Context, req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error) {
	return s.who(ctx, req)
}

func (s *observerServer) RestartService(ctx context.Context, req *msa_observer.RequestRestartService) (*msa_observer.Response, error) {
	return s.restartService(ctx, req)
}
