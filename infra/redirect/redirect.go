package redirect

import (
	context "context"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	api "github.com/hobord/infra2/api/grpc/redirect"
	requestId "github.com/hobord/infra2/infra/requestId"
	session "github.com/hobord/infra2/infra/session"

	log "github.com/hobord/infra2/log"
)

var redirectConn *grpc.ClientConn

func init() {
	serverAddr := os.Getenv("REDIRECT_GRPC_SERVER")
	if serverAddr == "" {
		serverAddr = "127.0.0.1:50052"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Debug(err)
	}
	redirectConn = conn
}

// RedirectHandler is a redirect midleware
func RedirectHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := session.SessionIDFromContext(r.Context())
		reqID := requestId.RequestIDFromContext(r.Context())
		log.Logger.Println("SessionID:" + sessionID)

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		u := scheme + "://" + r.Host + r.URL.Path
		redirection := requestRedirection(sessionID, reqID, u, r.Method, r.Header)
		if redirection.HttpStatusCode != 200 {
			http.Redirect(w, r, redirection.Location, int(redirection.HttpStatusCode))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func requestRedirection(sessionID, requestID, url, httpMethod string, httpHeaders http.Header) *api.GetRedirectionResponse {
	log.Logger.Println("call redirect service with: " + url)
	var client api.RedirectServiceClient
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client = api.NewRedirectServiceClient(redirectConn)
	defer cancel()

	headers := make(map[string]*api.HttpHeader)
	for headerKey, h := range httpHeaders {
		values := []string{}
		for _, val := range h {
			values = append(values, val)
		}
		headers[headerKey] = &api.HttpHeader{
			Header: values,
		}
	}

	message := &api.GetRedirectionMessage{
		SessionID:  sessionID,
		RequestID:  requestID,
		Url:        url,
		HttpMethod: httpMethod,
		Headers:    headers,
	}
	redirection, err := client.GetRedirection(ctx, message)

	if err != nil {
		// TODO: logs
		redirection = &api.GetRedirectionResponse{
			Location:       url,
			HttpStatusCode: 200,
		}
	}

	log.Logger.Println(redirection)
	return redirection
}
