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
	*Server
}

func CreateServer(cfg *Config) (*Server, error) {
	ginConfig, err := buildGinConfig(cfg)
	if err != nil {
		return nil, err
	}

	ginServer, err := ginConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &Server{
		ginServer: ginServer,
	}

	return server, nil
}

func (s *Server) PrepareRun() PreparedServer {
	fmt.Println("[Server] PrepareRun")

	s.ginServer.PrepareRun()
	ginInstance.InitRouter(s.ginServer.Engine)

	return PreparedServer{s}
}

func (s PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	return s.ginServer.Run()
}

func buildGinConfig(cfg *Config) (ginConfig *gin.Config, lastErr error) {
	ginConfig = gin.NewConfig()

	return cfg.GinConfig, nil
}
