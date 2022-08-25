package mysql

import (
	"fmt"
)

type Server struct{}

type PreparedServer struct {
	*Server
}

func (s *Server) PrepareRun() *PreparedServer {
	fmt.Println("[MySQL] PrepareRun")

	return &PreparedServer{s}
}

func (s *PreparedServer) Run() error {
	fmt.Println("[MySQL] Run")

	return nil
}

func (s *Server) init() {
	fmt.Println("[MySQL] Init")
}
