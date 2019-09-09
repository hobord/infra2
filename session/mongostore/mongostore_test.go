package mongostore

import (
	"os"
	"reflect"
	"testing"
)

func TestCreateMongoDBSessionStore(t *testing.T) {
	s := CreateMongoDBSessionStore(nil)
	st := reflect.TypeOf(s).String()

	if st != "*mongostore.MongoStore" {
		t.Errorf("I Got %v", st)
	}
}

func TestCreateSession(t *testing.T) {
	// SET TEST DB CONNECTION
	os.Setenv("DB_URL", "mongodb://10.20.153.18:27017")
	s := CreateMongoDBSessionStore(nil)

	sessionID, err := s.CreateSession(10)
	if err != nil {
		t.Errorf("CreateSession error %v", err)
	}

	if sessionID == "" {
		t.Errorf("Bad SessionID")
	}

}

func TestAddValueToSession(t *testing.T) {
	// SET TEST DB CONNECTION
	os.Setenv("DB_URL", "mongodb://10.20.153.18:27017")
	s := CreateMongoDBSessionStore(nil)

	sessionID, err := s.CreateSession(0)
	if err != nil {
		t.Errorf("CreateSession error %v", err)
	}
	s.AddValueToSession(sessionID, "foo", "baar")

}

func TestAddValuesToSession(t *testing.T) {
	// SET TEST DB CONNECTION
	os.Setenv("DB_URL", "mongodb://10.20.153.18:27017")
	s := CreateMongoDBSessionStore(nil)

	var values map[string]string
	values = make(map[string]string)
	values["key"] = "value"
	values["key2"] = "value2"

	sessionID, err := s.CreateSession(0)
	if err != nil {
		t.Errorf("CreateSession error %v", err)
	}
	err = s.AddValuesToSession(sessionID, values)
	if err != nil {
		t.Errorf("AddValuesToSession error %v", err)
	}

	sessionID, err = s.CreateSession(0)
	if err != nil {
		t.Errorf("CreateSession error %v", err)
	}
	updateResult, err := s.addValuesToSession(sessionID, values)
	if err != nil {
		t.Errorf("AddValuesToSession error %v", err)
	}

	if updateResult.MatchedCount != 1 {
		t.Errorf("updateResult.MatchedCount is not 1, it is: %v", updateResult.MatchedCount)
	}

	if updateResult.ModifiedCount != 1 {
		t.Errorf("updateResult.ModifiedCount is not 1, it is: %v", updateResult.ModifiedCount)
	}

	if updateResult.UpsertedCount != 0 {
		t.Errorf("updateResult.ModifiedCount is not 0, it is: %v", updateResult.UpsertedCount)
	}

	if updateResult.UpsertedID != nil {
		t.Errorf("updateResult.UpsertedID is not nil, it is: %v", updateResult.UpsertedID)
	}
}
