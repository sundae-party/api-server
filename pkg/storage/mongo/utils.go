package mongo

import (
	"errors"
	"strings"
)

// Decodekey extract integration and entity name from key and return integration name and entity name
// ex /integration01/entityA
func decodeKey(key string) (integrationName string, entityName string, err error) {
	entityInfos := strings.Split(key, "/")
	if len(entityInfos) != 2 {
		return "", "", errors.New("Invalid key format.")
	}
	return entityInfos[0], entityInfos[1], nil
}
