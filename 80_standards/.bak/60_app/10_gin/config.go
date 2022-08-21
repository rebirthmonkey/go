// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"
	"strconv"
)

type GINConfig struct {
	Mode            	string
	Healthz         	bool
	Log					*log.Options
	InsecureServerInfo	*InsecureServerInfo
	SecureServerInfo   	*SecureServerInfo
	//Jwt             	*JwtInfo
	//Middlewares     	[]string
	//EnableProfiling 	bool
	//EnableMetrics   	bool
}

type InsecureServerInfo struct {
	Address string
}

type SecureServerInfo struct {
	Address string
	CertFile string
	KeyFile string
}



func CreateGINConfig() *GINConfig {
	return &GINConfig{
		Mode:            	gin.ReleaseMode,
		Healthz:         	true,
		Log:			 	&log.Options{},
		//Middlewares:     []string{},
		//EnableProfiling: true,
		//EnableMetrics:   true,
		InsecureServerInfo: &InsecureServerInfo{},
		SecureServerInfo: 	&SecureServerInfo{},
/*		Jwt: &JwtInfo{
			Realm:      "iam jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},*/
	}
}

func (gcfg *GINConfig) BuildGINConfigFromOptions(opts *Options) error {
	gcfg.Mode = opts.Server.Mode
	gcfg.Healthz, _ = strconv.ParseBool(opts.Server.Healthz)

	gcfg.Log.Level = opts.Log.Level
	gcfg.Log.Format = opts.Log.Format
	gcfg.Log.EnableColor, _ = strconv.ParseBool(opts.Log.EnableColor)
	gcfg.Log.DisableCaller, _ = strconv.ParseBool(opts.Log.DisableCaller)
	gcfg.Log.OutputPaths = append(gcfg.Log.OutputPaths, opts.Log.OutputPaths)
	gcfg.Log.ErrorOutputPaths = append(gcfg.Log.ErrorOutputPaths, opts.Log.ErrorOutputPaths)
	gcfg.InsecureServerInfo.Address = opts.Insecure.Address+":"+opts.Insecure.Port
	gcfg.SecureServerInfo.Address = opts.Secure.Address+":"+opts.Secure.Port
	gcfg.SecureServerInfo.CertFile = opts.Secure.TLS.CertFile
	gcfg.SecureServerInfo.KeyFile = opts.Secure.TLS.PrivateKeyFile
	return nil
}
