package redirect

import (
	"context"
	"net/http"
	"net/url"

	st "github.com/golang/protobuf/ptypes/struct"
	redirectApi "github.com/hobord/infra2/api/grpc/redirect"
	config "github.com/hobord/infra2/redirect/config"
)

type SessionValues map[string]*st.Value

// Request representation
type Request struct {
	URL         string
	HTTPMethod  string
	HTTPHeaders map[string]*redirectApi.HttpHeader
	RequestID   string
	SessionID   string
}

// CalculateRedirections make recursive the redirections, with infinitive redirect loop detection.
func CalculateRedirections(ctx context.Context, configState *config.State, request Request, sessionValues *SessionValues, redirections map[string]int32) (redirectApi.GetRedirectionResponse, error) {
	//make business logic
	response := redirectApi.GetRedirectionResponse{
		Location:       request.URL,
		HttpStatusCode: http.StatusOK,
	}

	// Apply all rules
	newRedirection, err := applyRedirectionRules(ctx, configState, request, sessionValues)
	if err != nil {
		return response, nil // TODO: it is ok?
	}

	if newRedirection.HttpStatusCode == http.StatusOK {
		return response, nil
	}

	// infinitive redirect loop detection
	if httpStatusCode, ok := redirections[newRedirection.Location]; ok {
		if httpStatusCode == newRedirection.HttpStatusCode {
			return response, nil
		}
	}
	redirections[newRedirection.Location] = newRedirection.HttpStatusCode
	response = newRedirection

	// We have changes, lets make a new loop
	redirectTo := Request{
		SessionID:   request.SessionID,
		RequestID:   request.RequestID,
		URL:         response.Location,
		HTTPHeaders: request.HTTPHeaders,
		HTTPMethod:  request.HTTPMethod}

	r, err := CalculateRedirections(ctx, configState, redirectTo, sessionValues, redirections)
	if err != nil {
		return response, err
	}
	if r.HttpStatusCode != 200 {
		response = r
	}

	return response, nil
}

// applyRedirectionRules is apply the redirection rules
func applyRedirectionRules(ctx context.Context, configState *config.State, request Request, sessionValues *SessionValues) (redirectApi.GetRedirectionResponse, error) {
	response := redirectApi.GetRedirectionResponse{
		Location:       request.URL,
		HttpStatusCode: http.StatusOK,
	}

	u, err := url.Parse(request.URL)
	if err != nil {
		return response, err
	}

	// TODO: make businesslogic to here
	if host, ok := configState.RedirectionHosts[u.Host]; ok {
		if rules, ok := host[u.Scheme]; ok {
			for _, rule := range rules {
				switch rule.Type {
				case "Hash":
					if redirectTo, found := rule.TargetsByURL[u.String()]; found {
						response.Location = redirectTo.Target
						if redirectTo.HTTPStatusCode > 0 {
							response.HttpStatusCode = redirectTo.HTTPStatusCode
						} else {
							response.HttpStatusCode = rule.HTTPStatusCode
						}
						goto End
					}
				case "Regexp":
					if rule.Regexp != nil && rule.Regexp.MatchString(u.String()) {
						response.Location = rule.Target
						response.HttpStatusCode = rule.HTTPStatusCode
						goto End
					}
				case "CustomLogic":
					switch rule.LogicName {
					case "condition":
						// r = functionName(ctx, config, request, sessionValues)
						// if r.HttpStatusCode != 200 {
						// 	response = r
						//   	goto End
						// }
					}
				}
			}
		}
	}
End:
	// END of businesslogic

	return response, nil
}
