package redis

import (
	"encoding/json"
	"github.com/Pantani/errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

// Init initialize the database passing the host url.
// It returns an error if occurs.
func (db *Redis) Init(host string) error {
	options, err := redis.ParseURL(host)
	if err != nil {
		return errors.E(err, "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		return errors.E(err, "Redis connection test failed")
	}
	db.client = client
	return nil
}

// GetObject get object from key.
// It returns an error if occurs.
func (db *Redis) GetObject(key string, value interface{}) error {
	cmd := db.client.Get(key)
	if cmd.Err() != nil {
		return errors.E("not found", cmd.Err(), errors.Params{"key": key})
	}
	err := json.Unmarshal([]byte(cmd.Val()), value)
	if err != nil {
		return errors.E("not found", err, errors.Params{"key": key})
	}
	return nil
}

// AddObject add object for a key.
// It returns an error if occurs.
func (db *Redis) AddObject(key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"key": key})
	}
	cmd := db.client.Set(key, j, 0)
	if cmd.Err() != nil {
		return errors.E("not stored", cmd.Err(), errors.Params{"key": key})
	}
	return nil
}

// DeleteObject delete object from key.
// It returns an error if occurs.
func (db *Redis) DeleteObject(key string) error {
	cmd := db.client.Del(key)
	if cmd.Err() != nil {
		return errors.E("not deleted", cmd.Err(), errors.Params{"key": key})
	}
	return nil
}

// IsReady verify the database is ready.
// It returns true if ready.
func (db *Redis) IsReady() bool {
	if db.client == nil {
		return false
	}
	return db.client.Ping().Err() == nil
}
