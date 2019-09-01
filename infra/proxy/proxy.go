package proxy

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/hobord/infra2/log"
	"golang.org/x/net/http2"
)

var demoURL *url.URL

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func NewProxyHandler(demoURL *url.URL, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
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

		if next != nil {
			next.ServeHTTP(rw, req)
		}
	})
}
