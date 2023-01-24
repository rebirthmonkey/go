package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	Middlewares []string
	Healthz     bool

	*gin.Engine

	Insecure       *InsecureConfig
	insecureServer *http.Server

	Secure       *SecureConfig
	secureServer *http.Server
}

type PreparedServer struct {
	*Server
}

func (s *Server) PrepareRun() *PreparedServer {
	log.Info("[GinServer] PrepareRun")

	return &PreparedServer{s}
}

func (s *PreparedServer) Run() error {
	log.Info("[GinServer] Run")

	s.insecureServer = &http.Server{
		Addr:    s.Insecure.Address,
		Handler: s,
	}

	s.secureServer = &http.Server{
		Addr:    s.Secure.Address,
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("[GinServer] Listen on HTTP: %s", s.Insecure.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start Gin HTTP server: %s", err.Error())

			return err
		}

		log.Infof("[GinServer] Server on %s stopped", s.Insecure.Address)

		return nil
	})

	eg.Go(func() error {
		key, cert := s.Secure.KeyFile, s.Secure.CertFile
		if cert == "" || key == "" {
			return nil
		}

		log.Infof("[GinServer] Listen on HTTPS : %s", s.Secure.Address)

		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil {
			log.Fatalf("failed to start Gin HTTPS server: %s", err.Error())

			return err
		}

		log.Infof("[GinServer] Server on %s stopped", s.Secure.Address)

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Info(err.Error())
	}

	return nil
}

func (s *Server) init() {
	log.Info("[GinServer] Init")

	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// Setup do some setup work for gin engine.
func (s *Server) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

// InstallMiddlewares install generic middlewares.
func (s *Server) InstallMiddlewares() {
	// necessary middlewares
	//s.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "aaa": "bbb"}))

	//// install custom middlewares
	//for _, m := range s.middlewares {
	//	mw, ok := middleware.Middlewares[m]
	//	if !ok {
	//		log.Warnf("can not find middleware: %s", m)
	//
	//		continue
	//	}
	//
	//	log.Infof("install middleware: %s", m)
	//	s.Use(mw)
	//}
}

// InstallAPIs install generic apis.
func (s *Server) InstallAPIs() {
	if s.Healthz {
		s.GET("/healthz", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}

	s.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Version 0.1")
	})

}
