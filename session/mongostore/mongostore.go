package mongostore

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/hobord/infra2/log"
)

type sessionValues map[string]string

type session struct {
	Id     string
	Ttl    int64
	Expire int64
	Values sessionValues
}

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
		log.Logger.Error(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Logger.Error(err)
	}

	return client, nil
}

func CreateMongoDBSessionStore(client *mongo.Client) *MongoStore {
	var err error
	if client == nil {
		client, err = getMongoClient()
		if err != nil {
			log.Logger.Error("Cant get Mongodb Client")
		}
	}
	return &MongoStore{
		client: client,
	}
}

// CreateSession is create a new session with ttl, if ttl is 0 then the session is eternal
func (s *MongoStore) CreateSession(ttl int64) (string, error) {
	var err error

	collection := s.client.Database("sessions").Collection("sessions")

	uuid := uuid.New()

	var expire int64
	if ttl > 0 {
		expTime := time.Now().Local().Add(time.Second * time.Duration(ttl))
		expire = expTime.Unix()
	} else {
		expire = 0
	}

	sess := &session{
		Id:     uuid.String(),
		Ttl:    ttl,
		Expire: expire,
		Values: make(sessionValues),
	}

	_, err = collection.InsertOne(context.TODO(), sess)
	if err != nil {
		log.Logger.Error(err)
	}

	return uuid.String(), nil
}

func (s *MongoStore) AddValueToSession(id, key, value string) error {
	_, err := s.addValueToSession(id, key, value)
	return err
}

func (s *MongoStore) addValueToSession(sessionID, key, value string) (*mongo.UpdateResult, error) {
	var err error
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")

	// set filters and updates
	filter := bson.M{"id": sessionID}
	update := bson.M{"$set": bson.M{"values": bson.D{{key, value}}}}

	// update document
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Logger.Error(err)
	}

	return result, err
}

func (s *MongoStore) AddValuesToSession(id string, values sessionValues) error {
	_, err := s.addValuesToSession(id, values)
	return err
}

func (s *MongoStore) addValuesToSession(sessionID string, values sessionValues) (*mongo.UpdateResult, error) {
	var err error
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")

	// set filters and updates
	filter := bson.M{"id": sessionID}
	v := bson.D{}
	for key, value := range values {
		// update
		v = append(v, bson.E{key, value})
	}
	update := bson.M{
		"$set": bson.M{
			"values": v,
		},
	}
	// update document
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Logger.Error(err)
	}

	return updateResult, err
}

func (s *MongoStore) GetSessionValues(id string) (sessionValues, error) {

	return nil, nil
}

func (s *MongoStore) getSessionValues(id string) (sessionValues, error) {
	var err error
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")

	// set filters and updates
	var sess session
	filter := bson.M{"id": id}
	if err = collection.FindOne(ctx, filter).Decode(&sess); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return sess.Values, nil
}

/*
func InvalidateSession(id string) error                                {}
func InvalidateSessionValue(id, key string) error                      {}
func InvalidateSessionValues(id string, keys []string) error           {}
*/
// update.$set.values
