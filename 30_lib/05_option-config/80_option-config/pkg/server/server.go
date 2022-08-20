package server

import (
	"fmt"
)

type Server struct {
	healthz bool
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

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	if s.healthz {
		fmt.Println("[Server] Activate healthz option")
	}

	return nil
}
