package server

import (
	"fmt"
)

type Server struct {
	healthz bool
	mode    string
}

type PreparedServer struct {
	*Server
}

func (s *Server) init() {
	fmt.Println("[Server] Init")
}

func (s *Server) PrepareRun() *PreparedServer {
	fmt.Println("[Server] PrepareRun")

	return &PreparedServer{
		s,
	}
}

func (s *PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run on mode:", s.mode)

	if s.healthz {
		fmt.Println("[Server] Activate healthz option")
	}

	return nil
}
