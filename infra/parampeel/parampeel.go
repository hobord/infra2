package parampeel

import (
	"net/http"
	"net/url"

	st "github.com/golang/protobuf/ptypes/struct"

	requestId "github.com/hobord/infra2/infra/requestId"
	session "github.com/hobord/infra2/infra/session"
	log "github.com/hobord/infra2/log"
)

// ParamsHandler make parampeeling
func ParamsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Logger.Println("Params Init")

		parameters := filterParametersByURL(r)
		if len(parameters) > 0 {
			values := makeSessionValues(parameters)
			sessionID := session.SessionIDFromContext(r.Context())
			log.Logger.Println("SessionID:" + sessionID)
			session.AddValuesToSession(sessionID, values)
		}

		reqID := requestId.RequestIDFromContext(r.Context())
		log.Logger.Println("Hello request ID:" + reqID)
		if next != nil {
			next.ServeHTTP(w, r)
		}
		log.Logger.Println("Params END")
	})
}

func filterParametersByURL(r *http.Request) url.Values {
	var parameters map[string][]string

	if r.Method == http.MethodPost && r.Method == http.MethodPut && r.Method == http.MethodPatch {
		err := r.ParseForm()
		if err != nil {
			log.Logger.Error(err) // TODO: make nicer
			return make(map[string][]string)
		}
		parameters = r.Form
	} else {
		parameters = r.URL.Query()
	}

	// TODO: make filter logic, for saving different parameters by domain

	return parameters
}

func makeSessionValues(parameters url.Values) session.Values {
	values := make(session.Values)
	for key, value := range parameters {
		log.Logger.Println("i:" + key + " key:" + value[0])
		values[key] = &st.Value{Kind: &st.Value_StringValue{value[0]}} // TODO: Multiple value?
	}
	return values
}
