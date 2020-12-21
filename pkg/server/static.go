package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type srvMode int64

const (
	// HTTPMode for server listening on http port without TLS termination
	HTTPMode srvMode = 0
	// HTTPSMode for server listening on http port with TLS termination.
	// In this mode you have to provide a key and certificate to setting the TLS.
	// You can enable an automatic redirection from http to https in TLSConf parameter.
	HTTPSMode srvMode = 1
)

// ServerConfig define the HTTP server configuration.
type ServerConfig struct {
	// ServerMode define the server type, http or https
	// If a proxy with SSL termination is set in the front of this server the server should probably configured in http mode.
	ServerMode srvMode
	// HTTPAddr define the listening address for the http server "0.0.0.0:80", "127.0.0.1:80", ":80"
	// When EnableHTTPSredirect is set to true, this setting will be used for the http listener.
	// If the server is setting with HTTPSMode and EnableHTTPSredirect is set to false, this parameter will be ignored.
	HTTPAddr string
	// HTTPSAddr define the listening port for the https server "0.0.0.0:443", "127.0.0.1:443", ":443"
	// If the server is setting with HTTPMode, this parameter will be ignored.
	HTTPSAddr string
	// EnableHTTPSredirect when set to true, create an http server which will automatically redirect all incoming request to https server
	// If the server is setting with HTTPMode, this parameter will be ignored.
	EnableHTTPSredirect bool
	// Path to private key file for the server setting with HTTPS mode.
	// If the server is setting with HTTPMode, this parameter will be ignored.
	KeyPath string
	// Path to certificate file for the server setting with HTTPS mode.
	// If the server is setting with HTTPMode, this parameter will be ignored.
	CertPath string
	// EnableMTLS enable Mutual TLS authentication
	EnableMTLS bool
	// ClientCAsPath list of CA cert used to signe client certificates fot MTLS.
	// If EnableMTLS parameter is set to false, this setting will be ignored.
	ClientCAsPath []string
}

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// Serve create new web server and start to listen
func Serve(srvConf ServerConfig) {
	srvCount := 1
	router := mux.NewRouter()

	// Server rest API routes
	ServeAPI(router)

	// Create and start ws
	hub := newHub()
	go hub.run()
	// Serve ws
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// Serve single page application
	spa := spaHandler{
		staticPath: "build",
		indexPath:  "index.html",
	}
	router.PathPrefix("/").Handler(spa)

	// Create server instance
	var httpServer = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	var httpsServer = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if srvConf.ServerMode == HTTPMode {
		httpServer.Handler = router
		httpServer.Addr = srvConf.HTTPAddr
	} else if srvConf.ServerMode == HTTPSMode {
		httpsServer.Handler = router
		httpsServer.Addr = srvConf.HTTPSAddr

		// Setting redirect http to https
		if srvConf.EnableHTTPSredirect {
			httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, httpsPort, err := net.SplitHostPort(srvConf.HTTPSAddr)
				if err != nil {
					log.Fatalln(err)
				}
				http.Redirect(w, r, "https://"+r.Host+":"+httpsPort+r.URL.String(), http.StatusMovedPermanently)
			})
			srvCount++
		}

		// Setting MTLS
		if srvConf.EnableMTLS {
			caCertPool := x509.NewCertPool()
			// Create a CA certificate pool with all Client CAs listed in srvConf.ClientCAsPath
			for _, caPath := range srvConf.ClientCAsPath {
				caCert, err := ioutil.ReadFile(caPath)
				if err != nil {
					log.Fatal(err)
				}
				caCertPool.AppendCertsFromPEM(caCert)
			}

			// Create the TLS Config with the CA pool and enable Client certificate validation
			tlsConfig := &tls.Config{
				ClientCAs:  caCertPool,
				ClientAuth: tls.RequireAndVerifyClientCert,
			}
			tlsConfig.BuildNameToCertificate()
			httpsServer.TLSConfig = tlsConfig
		}
	}

	// Create waitGroup to wait until all http server stop
	wg := new(sync.WaitGroup)
	wg.Add(srvCount)

	// Start http server
	if srvConf.ServerMode == HTTPMode || srvConf.EnableHTTPSredirect {
		go func() {
			log.Printf("Listening on %s\n", srvConf.HTTPAddr)
			log.Fatal(httpServer.ListenAndServe())
			wg.Done()
		}()
	}

	// Start https server
	if srvConf.ServerMode == HTTPSMode {
		go func() {
			log.Printf("Listening on %s\n", srvConf.HTTPSAddr)
			log.Fatal(httpsServer.ListenAndServeTLS(srvConf.CertPath, srvConf.KeyPath))
			wg.Done()
		}()
	}
	wg.Wait()
}
