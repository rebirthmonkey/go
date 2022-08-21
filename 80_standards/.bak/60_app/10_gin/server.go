package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type GINServer struct {
	mode            	string
	healthz         	bool
	//middlewares     	[]string
	//enableProfiling 	bool
	//enableMetrics   	bool
	*gin.Engine
	insecureServer		*http.Server
	insecureServerInfo	*InsecureServerInfo
	secureServer 		*http.Server
	secureServerInfo   	*SecureServerInfo
	//Jwt             	*JwtInfo
}

func CreateGINServer(cfg *GINConfig) (*GINServer, error) {

	server := &GINServer{
		insecureServerInfo:	   cfg.InsecureServerInfo,
		secureServerInfo:      cfg.SecureServerInfo,
		mode:                  cfg.Mode,
		healthz:               cfg.Healthz,
		//enableMetrics:       cfg.EnableMetrics,
		//enableProfiling:     cfg.EnableProfiling,
		//middlewares:         cfg.Middlewares,
		Engine:                gin.New(),
	}

	return server, nil
}

func (server *GINServer) Init () error {

	initRouter(server)

	return nil
}

func (server *GINServer) Run() error {
	server.insecureServer = &http.Server{
		Addr:         server.insecureServerInfo.Address,
		Handler:      server,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	server.secureServer = &http.Server{
		Addr:    server.secureServerInfo.Address,
		Handler: server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s", server.insecureServerInfo.Address)

		//if err := server.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		if err := server.insecureServer.ListenAndServe(); err != nil  {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", server.insecureServerInfo.Address)

		return nil
	})

	eg.Go(func() error {
		key, cert := server.secureServerInfo.KeyFile, server.secureServerInfo.CertFile
		//if cert == "" || key == "" || server.SecureServingInfo.BindPort == 0 {
		if cert == "" || key == ""  {
			return nil
		}

		log.Infof("Start to listening the incoming requests on https address: %s", server.secureServerInfo.Address)

		//if err := server.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
		if err := server.secureServer.ListenAndServeTLS(cert, key); err != nil {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", server.secureServerInfo.Address)

		return nil
	})

	// Ping the server to make sure the router is working.
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if server.healthz {
	//	if err := server.ping(ctx); err != nil {
	//		return err
	//	}
	//}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
