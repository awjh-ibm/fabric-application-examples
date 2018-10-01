/*
SPDX-License-Identifier: Apache-2.0
*/

package datautils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Serialize returns JSON bytes of a given interface
func Serialize(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}

// Deserialize converts JSON into an interface of valid type if data has type that is supported
func Deserialize(data []byte, supportedTypes map[string]reflect.Type) (interface{}, error) {
	var err error

	jsonMap := make(map[string]interface{})

	err = json.Unmarshal(data, &jsonMap)

	if err != nil {
		return nil, errors.New("JSON passed invalid")
	}

	typeIDIFace, ok := jsonMap["type"]

	if !ok {
		return nil, errors.New("JSON missing field \"type\"")
	}

	typeID := fmt.Sprint(typeIDIFace)

	goalType, ok := supportedTypes[typeID]

	if !ok {
		return nil, fmt.Errorf("JSON passed had invalid type \"%s\"", typeID)
	}

	return DeserializeToType(data, goalType)
}

// DeserializeToType converts JSON into an interface that matches the passed type
func DeserializeToType(data []byte, goalType reflect.Type) (interface{}, error) {
	var err error

	jsonMap := make(map[string]interface{})

	err = json.Unmarshal(data, &jsonMap)

	if err != nil {
		return nil, errors.New("JSON passed invalid")
	}

	toReturn := reflect.New(goalType).Interface()

	json.Unmarshal(data, toReturn)

	return toReturn, nil
}
