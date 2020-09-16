package redis

import (
	"context"
	"encoding/json"

	"github.com/Pantani/errors"
)

// GetAllHMObjects get all objects from a hash map table.
// It returns the objects and an error if occurs.
func (db *Redis) GetAllHMObjects(ctx context.Context, entity string) (map[string]string, error) {
	cmd := db.client.HGetAll(ctx, entity)
	if err := cmd.Err(); err != nil {
		return nil, errors.E("not found", err, errors.Params{"entity": entity})
	}
	return cmd.Val(), nil
}

// GetHMObject get an object from a hash map table.
// It returns an error if occurs.
func (db *Redis) GetHMObject(ctx context.Context, entity, key string, value interface{}) error {
	cmd := db.client.HMGet(ctx, entity, key)
	if err := cmd.Err(); err != nil {
		return errors.E("not found", err, errors.Params{"entity": entity, "key": key})
	}
	val, ok := cmd.Val()[0].(string)
	if !ok {
		return errors.E("not found", errors.Params{"entity": entity, "key": key, "value": val})
	}
	err := json.Unmarshal([]byte(val), value)
	if err != nil {
		return errors.E("fail to unmarshal value", err, errors.Params{"entity": entity, "key": key, "value": val})
	}
	return nil
}

// AddHMObject add an object to a hash map table.
// It returns an error if occurs.
func (db *Redis) AddHMObject(ctx context.Context, entity, key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"key": key})
	}
	cmd := db.client.HMSet(ctx, entity, map[string]interface{}{key: j})
	if err := cmd.Err(); err != nil {
		return errors.E("not stored", err, errors.Params{"entity": entity, "key": key})
	}
	return nil
}

// DeleteHMObject delete an object from a hash map table.
// It returns an error if occurs.
func (db *Redis) DeleteHMObject(ctx context.Context, entity, key string) error {
	cmd := db.client.HDel(ctx, entity, key)
	if err := cmd.Err(); err != nil {
		return errors.E("not deleted", err, errors.Params{"entity": entity, "key": key})
	}
	return nil
}
