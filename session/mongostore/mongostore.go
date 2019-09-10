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
	session "github.com/hobord/infra2/session"
)

type sessionEntity struct {
	ID     string
	TTL    int64
	Expire int64
	Values session.SessionValues
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
	store := &MongoStore{
		client: client,
	}

	// frenquencyEnv := os.Getenv("REDIS_MAXTIMEOUT")
	// if frenquencyEnv == "" {
	// 	frenquencyEnv = "240"
	// }
	// fr, err := strconv.Atoi(frenquencyEnv)
	// if err != nil {
	// 	log.Logger.Fatal(err)
	// }
	// store.StartGbCollector(fr)

	return store
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

	sess := &sessionEntity{
		ID:     uuid.String(),
		TTL:    ttl,
		Expire: expire,
		Values: make(session.SessionValues),
	}

	_, err = collection.InsertOne(context.TODO(), sess)
	if err != nil {
		log.Logger.Error(err)
	}

	return uuid.String(), nil
}

func (s *MongoStore) AddValueToSession(id, key, value string) error {
	err := s.addValueToSession(id, key, value)
	return err
}

func (s *MongoStore) addValueToSession(sessionID, key, value string) error {
	var err error
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")

	// set filters and updates
	filter := bson.M{"id": sessionID}
	update := bson.M{"$set": bson.M{"values." + key: value}}

	// update document
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Logger.Error(err)
	}

	return err
}

func (s *MongoStore) AddValuesToSession(id string, values session.SessionValues) error {
	err := s.addValuesToSession(id, values)
	return err
}

func (s *MongoStore) addValuesToSession(sessionID string, values session.SessionValues) error {
	var err error
	for key, value := range values {
		s.addValueToSession(sessionID, key, value)
	}

	return err
}

func (s *MongoStore) GetSessionValues(id string) (session.SessionValues, error) {
	return s.getSessionValues(id)
}

func (s *MongoStore) getSessionValues(id string) (session.SessionValues, error) {
	var err error
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")

	// set filters and updates
	var sess sessionEntity
	filter := bson.M{"id": id}
	if err = collection.FindOne(ctx, filter).Decode(&sess); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return sess.Values, nil
}

func (s *MongoStore) InvalidateSession(id string) error {
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (s *MongoStore) InvalidateSessionValue(id, key string) error {
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")
	filter := bson.M{"id": id}

	update := bson.M{
		"$unset": bson.M{
			"values." + key: "",
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Logger.Error(err)
	}

	return err
}

func (s *MongoStore) InvalidateSessionValues(id string, keys []string) error {
	for _, key := range keys {
		err := s.InvalidateSessionValue(id, key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *MongoStore) StartGbCollector(intervall int) {
	for {
		go s.deleteInvalidSessions()
		time.Sleep(time.Second * time.Duration(intervall*1000))
	}
}

func (s *MongoStore) deleteInvalidSessions() {
	ctx := context.TODO()
	collection := s.client.Database("sessions").Collection("sessions")
	expTime := time.Now().Local().Unix()

	filter := bson.M{"$and": []interface{}{
		bson.M{"expire": bson.M{"$ne": 0}},
		bson.M{"expire": bson.M{"$lte": expTime}},
	}}

	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Logger.Error(err)
	}
}
