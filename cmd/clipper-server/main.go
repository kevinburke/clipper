package main

import (
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/kevinburke/clipper/server"
	"github.com/kevinburke/handlers"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/rest"
	"google.golang.org/appengine"
	applog "google.golang.org/appengine/log"
	yaml "gopkg.in/yaml.v2"
)

var logger log.Logger

func init() {
	logger = handlers.Logger
}

func commonMain() (*FileConfig, http.Handler) {
	flag.Parse()
	data, err := ioutil.ReadFile(*cfg)
	c := new(FileConfig)
	if err == nil {
		if err := yaml.Unmarshal(data, c); err != nil {
			logger.Error("Couldn't parse config file", "err", err)
			os.Exit(2)
		}
	} else {
		logger.Error("Couldn't find config file", "err", err)
		os.Exit(2)
	}
	var key nacl.Key
	if c.SecretKey == "" {
		logger.Warn("No secret key specified, generating a random one")
		key = nacl.NewKey()
	} else {
		key, err = nacl.Load(c.SecretKey)
		if err != nil {
			logger.Error("Error getting secret key", "err", err)
			os.Exit(2)
		}
	}

	if c.Port == nil {
		port, ok := os.LookupEnv("PORT")
		if ok {
			iPort, err := strconv.Atoi(port)
			if err != nil {
				logger.Error("Invalid port", "err", err, "port", port)
				os.Exit(2)
			}
			c.Port = &iPort
		} else {
			c.Port = &DefaultPort
		}
	}
	mux := server.NewServeMux(key)
	mux = handlers.UUID(mux)                              // add UUID header
	mux = handlers.Server(mux, "clipper-server/"+Version) // add Server header
	mux = handlers.Log(mux)                               // log requests/responses
	mux = handlers.Duration(mux)                          // add Duration header
	return c, mux
}

func main() {
	start := time.Now()
	c, mux := commonMain()
	if appengine.IsStandard() {
		rest.RegisterHandler(500, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := appengine.NewContext(r)
			err := rest.CtxErr(r)
			applog.Errorf(ctx, "%s %s: %v", r.Method, r.URL.Path, err)
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(500)
			w.Write([]byte("<html><body>Server Error</body></html>"))
		}))
		http.Handle("/", mux)
		appengine.Main()
		return
	}
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

// DefaultPort is the listening port if no other port is specified.
var DefaultPort = 8540

// The server's Version.
const Version = "0.5"

// FileConfig represents the data in a config file.
type FileConfig struct {
	// SecretKey is used to encrypt sessions and other data before serving it to
	// the client. It should be a hex string that's exactly 64 bytes long. For
	// example:
	//
	//   d7211b215341871968869dontusethisc0ff1789fc88e0ac6e296ba36703edf8
	//
	// That key is invalid - you can generate a random key by running:
	//
	//   openssl rand -hex 32
	//
	// If no secret key is present, we'll generate one when the server starts.
	// However, this means that sessions may error when the server restarts.
	//
	// If a server key is present, but invalid, the server will not start.
	SecretKey string `yaml:"secret_key"`

	// Port to listen on. Set to 0 to choose a port at random. If unspecified,
	// defaults to 7065.
	Port *int `yaml:"port"`

	// Set to true to listen for HTTP traffic (instead of TLS traffic). Note
	// you need to terminate TLS to use HTTP server push.
	HTTPOnly bool `yaml:"http_only"`

	// For TLS configuration.
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`

	// Add other configuration settings here.
}

var cfg = flag.String("config", "config.yml", "Path to a config file")
