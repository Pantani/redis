package redis

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

// GetAllHMObjects get all objects from a hash map table.
// It returns the objects and an error if occurs.
func (db *Redis) GetAllHMObjects(entity string) (map[string]string, error) {
	cmd := db.client.HGetAll(entity)
	if cmd.Err() != nil {
		return nil, errors.E("not found", cmd.Err())
	}
	return cmd.Val(), nil
}

// GetHMObject get an object from a hash map table.
// It returns an error if occurs.
func (db *Redis) GetHMObject(entity, key string, value interface{}) error {
	cmd := db.client.HMGet(entity, key)
	if cmd.Err() != nil {
		return errors.E("not found", cmd.Err(), errors.Params{"key": key})
	}
	val, ok := cmd.Val()[0].(string)
	if !ok {
		return errors.E("not found")
	}
	err := json.Unmarshal([]byte(val), value)
	if err != nil {
		return errors.E("not found", err, errors.Params{"key": key})
	}
	return nil
}

// AddHMObject add an object to a hash map table.
// It returns an error if occurs.
func (db *Redis) AddHMObject(entity, key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"key": key})
	}
	cmd := db.client.HMSet(entity, map[string]interface{}{key: j})
	if cmd.Err() != nil {
		return errors.E("not stored", cmd.Err(), errors.Params{"key": key})
	}
	return nil
}

// DeleteHMObject delete an object from a hash map table.
// It returns an error if occurs.
func (db *Redis) DeleteHMObject(entity, key string) error {
	cmd := db.client.HDel(entity, key)
	if cmd.Err() != nil {
		return errors.E("not deleted", cmd.Err(), errors.Params{"key": key})
	}
	return nil
}
