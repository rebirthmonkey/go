package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type GenericAPIServer struct {
	middlewares []string

	*gin.Engine
	healthz         bool
	enableMetrics   bool
	enableProfiling bool

	// InsecureInfo holds configuration of the insecure HTTP server.
	InsecureInfo *InsecureConfig

	insecureServer *http.Server
}

func initGenericAPIServer(s *GenericAPIServer) {
	fmt.Println("[GenericAPIServer] Init")

	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	fmt.Println("[GenericAPIServer] Run")

	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.InsecureInfo.Address,
		Handler: s,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.InsecureServingInfo.Address)

		return nil
	})

	//// Ping the server to make sure the router is working.
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if s.healthz {
	//	if err := s.ping(ctx); err != nil {
	//		return err
	//	}
	//}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

// Setup do some setup work for gin engine.
func (s *GenericAPIServer) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

// InstallMiddlewares install generic middlewares.
func (s *GenericAPIServer) InstallMiddlewares() {
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
func (s *GenericAPIServer) InstallAPIs() {
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}

	// install version handler
	s.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Version 0.1")
	})

}
