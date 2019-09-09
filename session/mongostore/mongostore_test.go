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
