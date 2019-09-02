package redimp

import (
	"context"
	"errors"

	log "github.com/hobord/infra2/log"

	"github.com/golang/protobuf/proto"
	st "github.com/golang/protobuf/ptypes/struct"
	pb "github.com/hobord/infra2/api/grpc/session"
	session "github.com/hobord/infra2/session"
)

type GrpcServer struct {
	store session.SessionStore
}

func (s *GrpcServer) getStore() session.SessionStore {
	return s.store
}

// CreateGrpcServer is a constructor for Session GrpcServer
func CreateGrpcServer(SessionStore session.SessionStore) (*GrpcServer, error) {
	if SessionStore == nil {
		log.Logger.Error("there is no session store")
		return nil, errors.New("We need some session store")
	}
	return &GrpcServer{
		store: SessionStore,
	}, nil
}

// CreateSession is create a new empty session
func (s *GrpcServer) CreateSession(ctx context.Context, in *pb.CreateSessionMessage) (*pb.SessionResponse, error) {
	id, err := s.store.CreateSession(in.Ttl)
	if err != nil {
		log.Logger.Errorf("can't create a session, err: %v", err)
		return &pb.SessionResponse{}, err
	}

	return &pb.SessionResponse{Id: id, Values: nil}, nil
}

// AddValueToSession is add value into the existing session
func (s *GrpcServer) AddValueToSession(ctx context.Context, in *pb.AddValueToSessionMessage) (*pb.SessionResponse, error) {
	data := proto.MarshalTextString(in.Value)
	err := s.store.AddValueToSession(in.Id, in.Key, data)
	if err != nil {
		log.Logger.Errorf("can't add value into session: sessionId: %v, key: %v, err: %v", in.Id, in.Key, err)
		return &pb.SessionResponse{}, err
	}

	response, err := s.getSessionValues(in.Id)
	if err != nil {
		log.Logger.Errorf("can't get values from session: sessionId: %v, err: %v", in.Id, err)
		return nil, err
	}

	return response, nil
}

// AddValuesToSession is add multiple values into the session
func (s *GrpcServer) AddValuesToSession(ctx context.Context, in *pb.AddValuesToSessionMessage) (*pb.SessionResponse, error) {
	var values map[string]string
	values = make(map[string]string)
	for key, val := range in.Values {
		values[key] = proto.MarshalTextString(val)
	}
	err := s.store.AddValuesToSession(in.Id, values)
	if err != nil {
		log.Logger.Errorf("can't add values into session: sessionId: %v, values: %v, err: %v", in.Id, values, err)
		return nil, err
	}

	response, err := s.getSessionValues(in.Id)
	if err != nil {
		log.Logger.Errorf("can't get values from session: sessionId: %v, err: %v", in.Id, err)
		return nil, err
	}

	return response, nil
}

func (s *GrpcServer) getSessionValues(id string) (*pb.SessionResponse, error) {
	values, err := s.store.GetSessionValues(id)
	if err != nil {
		return nil, err
	}
	response := &pb.SessionResponse{Id: id, Values: make(map[string]*st.Value)}
	for key, hval := range values {
		if key != "__TTL" {
			val := st.Value{}
			err := proto.UnmarshalText(hval, &val)
			if err != nil {
				return response, err
			}
			response.Values[key] = &val
		}
	}
	return response, nil
}

// GetSession return the session by id
func (s *GrpcServer) GetSession(ctx context.Context, in *pb.GetSessionMessage) (*pb.SessionResponse, error) {
	response, err := s.getSessionValues(in.Id)
	if err != nil {
		log.Logger.Errorf("can't get values from session: sessionId: %v, err: %v", in.Id, err)
		return nil, err
	}

	return response, nil
}

// InvalidateSession is delete the session
func (s *GrpcServer) InvalidateSession(ctx context.Context, in *pb.InvalidateSessionMessage) (*pb.SuccessMessage, error) {
	err := s.store.InvalidateSession(in.Id)
	if err != nil {
		log.Logger.Errorf("can't invalidate session: sessionId: %v, err: %v", in.Id, err)
		return &pb.SuccessMessage{Successfull: false}, err
	}

	return &pb.SuccessMessage{Successfull: true}, nil
}

// InvalidateSessionValue is remove one key from the session
func (s *GrpcServer) InvalidateSessionValue(ctx context.Context, in *pb.InvalidateSessionValueMessage) (*pb.SuccessMessage, error) {
	err := s.store.InvalidateSessionValue(in.Id, in.Key)
	if err != nil {
		log.Logger.Errorf("can't invalidate session value: sessionId: %v, key: %v, err: %v", in.Id, in.Key, err)
		return &pb.SuccessMessage{Successfull: false}, err
	}

	return &pb.SuccessMessage{Successfull: true}, nil
}

// InvalidateSessionValues is remove multiple keys from the session
func (s *GrpcServer) InvalidateSessionValues(ctx context.Context, in *pb.InvalidateSessionValuesMessage) (*pb.SuccessMessage, error) {
	err := s.store.InvalidateSessionValues(in.Id, in.Keys)
	if err != nil {
		return &pb.SuccessMessage{Successfull: false}, err
	}

	return &pb.SuccessMessage{Successfull: true}, nil
}
