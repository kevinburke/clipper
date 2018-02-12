// +build !appengine

package main

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kevinburke/handlers"
)

func init() {
	logger = handlers.Logger
}

func main() {
	start := time.Now()
	c, mux := commonMain()
	addr := ":" + strconv.Itoa(*c.Port)
	if c.HTTPOnly {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Error("Error listening", "addr", addr, "err", err)
			os.Exit(2)
		}
		logger.Info("Started server", "time", time.Since(start).Round(100*time.Microsecond),
			"protocol", "http", "port", *c.Port)
		http.Serve(ln, mux)
	} else {
		mux = handlers.STS(mux) // set Strict-Transport-Security header
		if c.CertFile == "" {
			c.CertFile = "certs/leaf.pem"
		}
		if _, err := os.Stat(c.CertFile); os.IsNotExist(err) {
			logger.Error("Could not find a cert file; generate using 'make generate_cert'", "file", c.CertFile)
			os.Exit(2)
		}
		if c.KeyFile == "" {
			c.KeyFile = "certs/leaf.key"
		}
		if _, err := os.Stat(c.KeyFile); os.IsNotExist(err) {
			logger.Error("Could not find a key file; generate using 'make generate_cert'", "file", c.KeyFile)
			os.Exit(2)
		}
		logger.Info("Starting server", "time", time.Since(start).Round(100*time.Microsecond), "protocol", "https", "port", *c.Port)
		listenErr := http.ListenAndServeTLS(addr, c.CertFile, c.KeyFile, mux)
		logger.Error("server shut down", "err", listenErr)
	}
}
