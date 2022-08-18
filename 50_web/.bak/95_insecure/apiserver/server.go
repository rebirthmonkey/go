package apiserver

import (
	"fmt"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/apiserver/config"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/pkg/server"
)

type apiServer struct {
	genericAPIServer *server.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	fmt.Println("[APIServer] PrepareRun")

	initRouter(s.genericAPIServer.Engine)

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	fmt.Println("[PreparedAPIServer] Run")

	return s.genericAPIServer.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *server.Config, lastErr error) {
	genericConfig = server.NewConfig()

	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}
