package memstore

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
)

type memStoreService struct {
	Pool *redis.Pool
}

func NewMemStoreService(port int) service.MemStoreService {
	optDial := redis.DialConnectTimeout(1 * time.Second)
	optRead := redis.DialReadTimeout(1 * time.Second)
	optWrite := redis.DialWriteTimeout(1 * time.Second)
	address := fmt.Sprintf("localhost:%d", port)
	pool := redis.Pool{
		MaxIdle:   1000,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, optDial, optRead, optWrite)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	err := initialize(&pool)
	if err != nil {
		log.Fatalln(err)
	}

	memStoreService := &memStoreService{
		Pool: &pool,
	}
	return memStoreService
}

func (s *memStoreService) Get(key string) ([]byte, error) {
	con := s.Pool.Get()
	defer con.Close()

	result, err := redis.Bytes(con.Do("GET", key))
	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, errors.New("no cache")
	}
	return result, nil
}

func (s *memStoreService) Add(key string, value []byte, sec int) error {
	con := s.Pool.Get()
	defer con.Close()

	_, err := con.Do("SET", key, value, "EX", sec)
	if err != nil {
		return err
	}
	return nil
}
