package gin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Server struct {
	middlewares []string
	healthz     bool

	*gin.Engine

	// Insecure holds configuration of the insecure HTTP server.
	Insecure *InsecureConfig

	insecureServer *http.Server
}

type PreparedServer struct {
	*Server
}

func (s *Server) PrepareRun() PreparedServer {
	fmt.Println("[Server] PrepareRun")

	return PreparedServer{s}
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *Server) Run() error {
	fmt.Println("[Server] Run")
	s.Insecure.Address = "127.0.0.1:8080"
	fmt.Println("[Server] Run", s.Insecure.Address)

	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.Insecure.Address,
		Handler: s,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s", s.Insecure.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.Insecure.Address)

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *Server) initGinServer() {
	fmt.Println("[GinServer] Init")

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
	s.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "aaa": "bbb"}))

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
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}

	s.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Version 0.1")
	})

}
