package session

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	sessionApi "github.com/hobord/infra2/api/grpc/session"
	log "github.com/hobord/infra2/log"

	st "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc"
)

type Values map[string]*st.Value
type sessionKey int

const sessionIDKey sessionKey = 0

var sessionConn *grpc.ClientConn

func init() {
	serverAddr := os.Getenv("SESSION_GRPC_SERVER")
	if serverAddr == "" {
		serverAddr = "10.20.35.111:30645"
		serverAddr = "127.0.0.1:50051"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Fatal(err)
	}
	sessionConn = conn
}

// SessionHandler is a middleware handler
func SessionHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Logger.Println("Session Handler")

		ctx := newContextWithSessionID(r.Context(), r)
		sessionID := SessionIDFromContext(ctx)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)

		log.Logger.Println("SessionID:" + sessionID)
		log.Logger.Println("Session NEXT")
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Logger.Println("Session NEXT END")
	}
}

// SessionIDFromContext return session id
func SessionIDFromContext(ctx context.Context) string {
	return ctx.Value(sessionIDKey).(string)
}

// Checking the HTTP GET,POST parameters and the COOKIE.
// If sessionID is not prosented then create a new session
func newContextWithSessionID(ctx context.Context, r *http.Request) context.Context {
	var sessionID string
	sessionID = ""
	if r.Method == http.MethodPost {
		sessionID = checkBody(r)
	}
	if sessionID == "" {
		sessionID = checkGetParams(r)
		if sessionID == "" {
			sessionID = checkCookie(r)
			if sessionID == "" {
				ttl := getSessionTTL(r)
				sessionID = createSession(ttl)
				return context.WithValue(ctx, sessionIDKey, sessionID)
			}
		}
	}

	// Check the session is exists, if not create a new
	sess, err := GetSession(sessionID)
	if err != nil || sess == nil || len(sess.Values) == 0 {
		ttl := getSessionTTL(r)
		sessionID = createSession(ttl)
	}
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func checkBody(r *http.Request) string {
	var sessionID string
	err := r.ParseForm()
	if err != nil {
		sessionID = ""
	} else {
		sessionID = r.FormValue(getSessionRequestPostKey(r))
	}
	return sessionID
}

func checkGetParams(r *http.Request) string {
	var sessionID string
	sessionID = r.URL.Query().Get(getSessionRequestGetKey(r))
	return sessionID
}

func checkCookie(r *http.Request) string {
	var sessionID string
	cookie, err := r.Cookie(getSessionCookieKey(r))
	if err != nil || cookie == nil || cookie.Value == "" {
		sessionID = ""
	} else {
		sessionID = cookie.Value
	}
	return sessionID
}

func createSession(ttl int64) string {
	var client sessionApi.SessionServiceClient
	client = sessionApi.NewSessionServiceClient(sessionConn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsession, err := client.CreateSession(ctx, &sessionApi.CreateSessionMessage{Ttl: ttl})
	if err != nil {
		// TODO: errorhandling
		log.Logger.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	return dsession.Id
}

func getSessionTTL(r *http.Request) int64 {
	ttlstr := os.Getenv("SESSION_TTL")
	if ttlstr == "" {
		ttlstr = "0"
	}
	ttl, err := strconv.ParseInt(ttlstr, 10, 64)
	if err != nil {
		panic(err)
	}
	return ttl
}

func getSessionCookieKey(r *http.Request) string {
	ck := os.Getenv("SESSION_COOKIE_KEY")
	if ck == "" {
		ck = "session"
	}
	return ck
}

func getSessionRequestGetKey(r *http.Request) string {
	ck := os.Getenv("SESSION_REQUEST_GET_KEY")
	if ck == "" {
		ck = "session"
	}
	return ck
}
func getSessionRequestPostKey(r *http.Request) string {
	ck := os.Getenv("SESSION_REQUEST_POST_KEY")
	if ck == "" {
		ck = "session"
	}
	return ck
}

func GetSession(sessionID string) (*sessionApi.SessionResponse, error) {
	var client sessionApi.SessionServiceClient
	client = sessionApi.NewSessionServiceClient(sessionConn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsession, err := client.GetSession(ctx, &sessionApi.GetSessionMessage{Id: sessionID})
	if err != nil {
		return &sessionApi.SessionResponse{}, err
	}

	return dsession, nil
}

func AddValuesToSession(sessionID string, values Values) (*sessionApi.SessionResponse, error) {
	var client sessionApi.SessionServiceClient
	client = sessionApi.NewSessionServiceClient(sessionConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &sessionApi.AddValuesToSessionMessage{Id: sessionID, Values: values}
	response, err := client.AddValuesToSession(ctx, message)
	if err != nil {
		return &sessionApi.SessionResponse{}, err
	}
	return response, nil
}

func AddValueToSession(sessionID string, value *st.Value) (*sessionApi.SessionResponse, error) {
	var client sessionApi.SessionServiceClient
	client = sessionApi.NewSessionServiceClient(sessionConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &sessionApi.AddValueToSessionMessage{Id: sessionID, Value: value}
	response, err := client.AddValueToSession(ctx, message)
	if err != nil {
		return &sessionApi.SessionResponse{}, err
	}
	return response, nil
}

func InvalidateSession(sessionID string) (*sessionApi.SuccessMessage, error) {
	var client sessionApi.SessionServiceClient
	client = sessionApi.NewSessionServiceClient(sessionConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &sessionApi.InvalidateSessionMessage{Id: sessionID}
	response, err := client.InvalidateSession(ctx, message)
	if err != nil {
		return &sessionApi.SuccessMessage{}, err
	}
	return response, nil
}
