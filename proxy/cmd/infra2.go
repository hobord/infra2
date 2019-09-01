package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	log "github.com/hobord/infra2/log"
	"golang.org/x/net/http2"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	demoURL, err := url.Parse("http://netlab.hu")
	if err != nil {
		log.Logger.Fatal(err)
	}
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8100"
	}

	// rh := ctxrouter.RouterHandler()
	// rdh := redirect.RedirectHandler(rh)
	// pmh := httparams.ParamsHandler(rdh)
	// sh := session.SessionHandler(pmh)
	// ridh := requestId.RequestIDHandler(sh)
	// httplogger := log.HttpLoggerHandler(ridh)
	// log.Logger.Fatal(http.ListenAndServe(":"+httpPort, httplogger))
	// proxy := httputil.NewSingleHostReverseProxy(demoURL)

	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// make request
		req.Host = demoURL.Host
		req.URL.Host = demoURL.Host
		req.URL.Scheme = demoURL.Scheme
		req.RequestURI = ""

		// Add x-forwarded-for to the header
		s, _, _ := net.SplitHostPort(req.RemoteAddr)
		req.Header.Set("X-Forwarded-For", s)

		// http2
		http2.ConfigureTransport(http.DefaultTransport.(*http.Transport))

		// Do request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			log.Logger.Error(err)
			return
		}

		// copy the response to rw
		// Copy the headers
		for key, values := range resp.Header {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}

		// add stream support
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-time.Tick(10 * time.Millisecond):
					rw.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()

		// Support Trailer
		trailerKeys := []string{}
		for key := range resp.Trailer {
			trailerKeys = append(trailerKeys, key)
		}
		rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)

		// Support Trailer continue
		for key, values := range resp.Trailer {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}

		close(done)
	})

	log.Logger.Fatal(http.ListenAndServe(":"+httpPort, proxy))

	// https
	// log.Logger.Fatal(http.ListenAndServeTLS(":"+httpPort, "cert.pem", "key.pem" , proxy))
}
