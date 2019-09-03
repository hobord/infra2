package main

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"

	"github.com/hobord/infra2/infra/parampeel"
	"github.com/hobord/infra2/infra/proxy"
	"github.com/hobord/infra2/infra/redirect"
	"github.com/hobord/infra2/infra/requestId"
	"github.com/hobord/infra2/infra/session"
	log "github.com/hobord/infra2/log"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	demoURL, err := url.Parse(os.Getenv("DEFAULT_WEBSERVER"))
	if err != nil {
		log.Logger.Fatal(err)
	}
	// proxyHandler := httputil.NewSingleHostReverseProxy(demoURL)
	proxyHandler := proxy.NewProxyHandler(demoURL, nil)

	redirectHandler := redirect.RedirectHandler(proxyHandler)
	paramsHandler := parampeel.ParamsHandler(redirectHandler)
	sessionHandler := session.SessionHandler(paramsHandler)
	requestIDHandler := requestId.RequestIDHandler(sessionHandler)

	logHandler := log.HttpLoggerHandler(requestIDHandler)

	log.Logger.Fatal(http.ListenAndServe(":"+httpPort, logHandler))

	// https
	// log.Logger.Fatal(http.ListenAndServeTLS(":"+httpPort, "cert.pem", "key.pem" , proxy))
}
