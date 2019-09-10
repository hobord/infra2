package mongostore

import (
	"os"
	"testing"
)

/*
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
*/
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
	err = s.addValuesToSession(sessionID, values)
	if err != nil {
		t.Errorf("AddValuesToSession error %v", err)
	}

	var values2 map[string]string
	values2 = make(map[string]string)
	values2["key2"] = "value2"
	values2["key3"] = "value3"
	err = s.addValuesToSession(sessionID, values2)

	err = s.InvalidateSessionValue(sessionID, "key")

	_, err = s.addValueToSession(sessionID, "newKey", "val")

	_, err = s.addValueToSession(sessionID, "key2", "newval")

}
