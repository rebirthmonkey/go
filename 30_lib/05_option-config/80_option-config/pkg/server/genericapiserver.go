package server

import (
	"fmt"
)

type GenericAPIServer struct {
	healthz bool
}

func initGenericAPIServer(s *GenericAPIServer) {
	fmt.Println("[GenericAPIServer] Init")
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	fmt.Println("[GenericAPIServer] Run")

	fmt.Println("s.healthz", s.healthz)

	if s.healthz {
		fmt.Println("[GenericAPIServer] Activate healthz option")
	}

	return nil
}
