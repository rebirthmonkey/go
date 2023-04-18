package server

import (
	"fmt"

	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_server/pkg/server"
)

type Server struct {
	genericServer *server.Server
}

type PreparedServer struct {
	preparedGenericServer *server.PreparedServer
}

func NewServer(opts *Options) (*Server, error) {
	config := NewConfig()
	opts.ApplyTo(config)
	serverInstance, err := config.Complete().New()

	return serverInstance, err
}

func (s *Server) PrepareRun() *PreparedServer {
	fmt.Println("[Server] PrepareRun")

	return &PreparedServer{
		preparedGenericServer: s.genericServer.PrepareRun(),
	}
}

func (s *PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	return s.preparedGenericServer.Run()
}
