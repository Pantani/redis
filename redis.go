package redis

import (
	"context"
	"encoding/json"

	"github.com/Pantani/errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	client  *redis.Client
	context *context.Context
}

// Init initialize the database passing the host url.
// It returns an error if occurs.
func (db *Redis) Init(host string) error {
	options, err := redis.ParseURL(host)
	if err != nil {
		return errors.E(err, "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return errors.E(err, "Redis connection test failed")
	}
	db.client = client
	return nil
}

// GetObject get object from key.
// It returns an error if occurs.
func (db *Redis) GetObject(ctx context.Context, key string, value interface{}) error {
	cmd := db.client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return errors.E("not found", err, errors.Params{"key": key})
	}
	val := cmd.Val()
	err := json.Unmarshal([]byte(val), value)
	if err != nil {
		return errors.E("fail to unmarshal value", err, errors.Params{"key": key, "value": val})
	}
	return nil
}

// AddObject add object for a key.
// It returns an error if occurs.
func (db *Redis) AddObject(ctx context.Context, key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E("fail to marshal value", err, errors.Params{"key": key})
	}
	cmd := db.client.Set(ctx, key, j, 0)
	if err := cmd.Err(); err != nil {
		return errors.E("not stored", err, errors.Params{"key": key})
	}
	return nil
}

// DeleteObject delete object from key.
// It returns an error if occurs.
func (db *Redis) DeleteObject(ctx context.Context, key string) error {
	cmd := db.client.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return errors.E("not deleted", err, errors.Params{"key": key})
	}
	return nil
}

// IsReady verify the database is ready.
// It returns true if ready.
func (db *Redis) IsReady(ctx context.Context) bool {
	if db.client == nil {
		return false
	}
	return db.client.Ping(ctx).Err() == nil
}
