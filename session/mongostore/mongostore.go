package mongostore

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/hobord/infra2/log"
)

type MongoStore struct {
	client *mongo.Client
}

func getMongoClient() (*mongo.Client, error) {
	mongoURL := os.Getenv("DB_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Logger.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Logger.Fatal(err)
	}
	return client, nil
}

func CreateMongoDBSessionStore(client *mongo.Client) *MongoStore {
	var err error
	if client == nil {
		client, err = getMongoClient()
		if err != nil {
			log.Logger.Fatal("Cant get Mongodb Client")
		}
	}
	return &MongoStore{
		client: client,
	}
}

// CreateSession is create a new session with ttl, if ttl is 0 then the session is eternal
func (s *MongoStore) CreateSession(ttl int64) (string, error) {
	var err error

	ctx := context.Background()
	err = s.client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// defer conn.Close() // TODO: ???
	uuid := uuid.New()

	if ttl > 0 {
		ttlstr := fmt.Sprintf("%d", ttl)
		err = s.addValueToSession(uuid.String(), "__TTL", ttlstr)
		if err != nil {
			return "", err
		}
	} else {
		err = s.addValueToSession(uuid.String(), "__TTL", "0")
		if err != nil {
			return "", err
		}
	}

	return uuid.String(), nil
}

func (s *MongoStore) createSession() error {
	return nil
}

func (s *MongoStore) addValueToSession(sessionID, key, value string) error {
	return nil
}

/*
func AddValueToSession(id, key, value string) error                    {}
func AddValuesToSession(id string, values session.SessionValues) error {}
func GetSessionValues(id string) (session.SessionValues, error)        {}
func InvalidateSession(id string) error                                {}
func InvalidateSessionValue(id, key string) error                      {}
func InvalidateSessionValues(id string, keys []string) error           {}
*/
