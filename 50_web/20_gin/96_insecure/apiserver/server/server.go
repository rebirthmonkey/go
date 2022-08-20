package server

import (
	"fmt"

	ginInstance "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/server/gin"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/gin"
)

type Server struct {
	ginServer *gin.Server
}

type PreparedServer struct {
	preparedGinServer *gin.PreparedServer
}

func NewServer(opts *Options) (*Server, error) {
	config := NewConfig()
	opts.ApplyTo(config)
	serverInstance, err := config.Complete().New()

	return serverInstance, err
}

func (s *Server) PrepareRun() PreparedServer {
	fmt.Println("[Server] PrepareRun")

	ginInstance.InitRouter(s.ginServer.Engine)

	return PreparedServer{
		preparedGinServer: s.ginServer.PrepareRun(),
	}
}

func (s PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	return s.preparedGinServer.Run()
}
