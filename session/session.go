package session

type SessionValues map[string]string

type SessionStore interface {
	CreateSession(ttl int64) (string, error)
	AddValueToSession(id, key, value string) error
	AddValuesToSession(id string, values SessionValues) error
	GetSessionValues(id string) (SessionValues, error)
	InvalidateSession(id string) error
	InvalidateSessionValue(id, key string) error
	InvalidateSessionValues(id string, keys []string) error
}
