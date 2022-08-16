package apiserver

import (
	"fmt"

	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/apiserver/config"
	genericapiserver "github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/pkg/server"
)

type apiServer struct {
	genericAPIServer *genericapiserver.GenericAPIServer
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

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	fmt.Println("[PreparedAPIServer] Run")

	return s.genericAPIServer.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()

	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}
