package redirect

import (
	"context"
	"fmt"

	api "github.com/hobord/infra2/api/grpc/redirect"
	config "github.com/hobord/infra2/redirect/config"
)

// GrpcServer is base struct
type GrpcServer struct {
	configState *config.State
}

// CreateGrpcServer make an instace of GrpcServer
func CreateGrpcServer() *GrpcServer {
	configState := &config.State{}
	configState.LoadConfigs("configs")

	srv := &GrpcServer{
		configState: configState,
	}

	return srv
}

// GetRedirection is implementing RedirectService rcp function
func (s *GrpcServer) GetRedirection(ctx context.Context, in *api.GetRedirectionMessage) (*api.GetRedirectionResponse, error) {
	fmt.Printf("Get redirection: %v", in)
	sessionValues := &SessionValues{} // TODO: get session
	redirections := make(map[string]int32)

	request := Request{
		URL:         in.Url,
		HTTPMethod:  in.HttpMethod,
		HTTPHeaders: in.Headers,
		RequestID:   in.RequestID,
		SessionID:   in.SessionID,
	}

	response, err := CalculateRedirections(ctx, s.configState, request, sessionValues, redirections)
	if err != nil {
		return &response, err
	}
	r := ParamPeeling(ctx, s.configState, Request{
		URL:         response.Location,
		HTTPMethod:  in.HttpMethod,
		HTTPHeaders: in.Headers,
		RequestID:   in.RequestID,
		SessionID:   in.SessionID,
	})
	if r.HttpStatusCode != 200 {
		return &r, nil
	}
	fmt.Printf("Response: %v", response)
	return &response, err
}
