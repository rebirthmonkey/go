package server

import (
	"fmt"
)

type GenericAPIServer struct {
	healthz bool
}

func initGenericAPIServer(s *GenericAPIServer) {
	fmt.Println("[Generic APIServer] Init")
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	fmt.Println("[Generic APIServer] Run")

	if s.healthz {
		fmt.Println("[Generic APIServer] Activate healthz option")
	}

	return nil
}
