package redistore

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/google/uuid"
	log "github.com/hobord/infra2/log"
	session "github.com/hobord/infra2/session"
)

// RedisStore is a redsi implementation of session store
type RedisStore struct {
	ConnectionPool *redis.Pool
}

// NewRedisPool create a new redis connection pool
func NewRedisPool(server string, dbno int, password string, maxIdle int, idleTimeout int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Second * time.Duration(idleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", dbno)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// CreateRedisStore is the constructor of the redis session store
func CreateRedisStore(connectionPool *redis.Pool) *RedisStore {
	if connectionPool == nil {
		var err error
		rdHost := os.Getenv("REDIS_HOST")
		if rdHost == "" {
			rdHost = "127.0.0.1"
		}
		rdPort := os.Getenv("REDIS_PORT")
		if rdPort == "" {
			rdPort = "6379"
		}

		rdDbEnv := os.Getenv("REDIS_SESSION_DB")
		if rdDbEnv == "" {
			rdDbEnv = "0"
		}
		rdDb, err := strconv.Atoi(rdDbEnv)
		if err != nil {
			log.Logger.Fatalf("Failed to connect parse redis-server DB (%s)", rdDbEnv)
		}

		password := os.Getenv("REDIS_PASSWORD")

		maxIdleEnv := os.Getenv("REDIS_MAXIDLE")
		if maxIdleEnv == "" {
			maxIdleEnv = "3"
		}
		maxIdle, err := strconv.Atoi(maxIdleEnv)
		if err != nil {
			log.Logger.Fatalf("Failed to connect parse redis-server DB (%s)", rdDbEnv)
		}

		maxTimeOutEnv := os.Getenv("REDIS_MAXTIMEOUT")
		if maxTimeOutEnv == "" {
			maxTimeOutEnv = "240"
		}
		maxTimeOut, err := strconv.Atoi(maxTimeOutEnv)
		if err != nil {
			log.Logger.Fatalf("Failed to connect parse redis-server DB (%s)", rdDbEnv)
		}

		rediserver := rdHost + ":" + rdPort

		// Redigo Client
		connectionPool = NewRedisPool(rediserver, rdDb, password, maxIdle, maxTimeOut)
	}
	return &RedisStore{
		ConnectionPool: connectionPool,
	}
}

// CreateSession is create a new session with ttl, if ttl is 0 then the session is eternal
func (s *RedisStore) CreateSession(ttl int64) (string, error) {
	var err error
	conn := s.ConnectionPool.Get()
	defer conn.Close() // TODO: ???
	uuid := uuid.New()
	if ttl > 0 {
		ttlstr := fmt.Sprintf("%d", ttl)
		err = addValueToSession(conn, uuid.String(), "__TTL", ttlstr)
		if err != nil {
			return "", err
		}
		conn.Send("EXPIRE", uuid.String(), ttlstr)
	} else {
		err = addValueToSession(conn, uuid.String(), "__TTL", "0")
		if err != nil {
			return "", err
		}
	}

	err = conn.Flush()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// AddValueToSession is add a value with key to session
func (s *RedisStore) AddValueToSession(id, key, value string) error {
	conn := s.ConnectionPool.Get()
	defer conn.Close() // TODO: ???

	return addValueToSession(conn, id, key, value)
}
func addValueToSession(conn redis.Conn, id, key, value string) error {
	if key != "__TTL" {
		res, err := redis.StringMap(conn.Do("HGETALL", id))
		if err != nil {
			return err
		}
		if len(res) == 0 {
			return errors.New("session already go on")
		}
	}
	return conn.Send("HSET", id, key, value)
}

func (s *RedisStore) AddValuesToSession(id string, values session.SessionValues) error {
	conn := s.ConnectionPool.Get()
	defer conn.Close() // TODO: ???

	for key, val := range values {
		err := addValueToSession(conn, id, key, val)
		if err != nil {
			return err
		}
	}

	err := conn.Flush()
	if err != nil {
		return err
	}
	return nil
}

// GetSessionValues return the session values by id
func (s *RedisStore) GetSessionValues(id string) (session.SessionValues, error) {
	conn := s.ConnectionPool.Get()
	defer conn.Close() // TODO: ???
	return getValuesBySessionID(conn, id)
}

func getValuesBySessionID(conn redis.Conn, id string) (session.SessionValues, error) {
	res, err := conn.Do("HGETALL", id)
	values, err := redis.StringMap(res, err)
	if err != nil {
		return nil, err
	}
	return values, nil
}

// InvalidateSession is invalidate the session by id
func (s *RedisStore) InvalidateSession(id string) error {
	conn := s.ConnectionPool.Get()
	defer conn.Close() // TODO: ???

	err := conn.Send("DEL", id)
	if err != nil {
		log.Logger.Errorf("Something went wrong: %s", err)
		return err
	}

	err = conn.Flush()
	if err != nil {
		log.Logger.Errorf("Something went wrong: %s", err)
		return err
	}

	return nil
}

// InvalidateSessionValue is invalidate one specific value in the session, by session id and key
func (s *RedisStore) InvalidateSessionValue(id, key string) error {
	conn := s.ConnectionPool.Get()
	defer conn.Close()

	err := invalidateSessionValue(conn, id, key)
	if err != nil {
		log.Logger.Errorf("Something went wrong: %s", err)
		return err
	}

	err = conn.Flush()
	if err != nil {
		log.Logger.Errorf("Something went wrong: %s", err)
		return err
	}

	return nil
}

// InvalidateSessionValues is invalidate multiple specific value in the session, by session id and keys
func (s *RedisStore) InvalidateSessionValues(id string, keys []string) error {
	conn := s.ConnectionPool.Get()
	defer conn.Close()

	for _, key := range keys {
		err := invalidateSessionValue(conn, id, key)
		if err != nil {
			log.Logger.Errorf("Something went wrong: %s", err)
			return err
		}
	}

	err := conn.Flush()
	if err != nil {
		log.Logger.Errorf("Something went wrong: %s", err)
		return err
	}

	return nil
}

func invalidateSessionValue(conn redis.Conn, id string, key string) error {
	return conn.Send("HDEL ", id, key)
}
