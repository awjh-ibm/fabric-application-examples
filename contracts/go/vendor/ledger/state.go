/*
SPDX-License-Identifier: Apache-2.0
*/

package ledger

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MakeKey generates a key from a slice
func MakeKey(keyParts ...interface{}) (string, error) {
	strSlice := []string{}
	for i, el := range keyParts {
		bytes, err := json.Marshal(el)

		if err != nil {
			return "", fmt.Errorf("Item %d in slice could not be turned to JSON string", i)
		}

		strSlice = append(strSlice, string(bytes))
	}

	return strings.Join(strSlice, ":"), nil
}

// SplitKey splits a key string on ":"
func SplitKey(key string) []string {
	return strings.Split(key, ":")
}

// StateInterface used to handle accessing the component parts
// of a state. States have a type, unique key, and a lifecycle
// current state
type StateInterface interface {
	GetType() string
	SetKey(string)
	GetKey() string
	GetCurrentState() interface{}
	GetSplitKey() []string
}

// State is a preconfigured struct that helps implement key
// management
type State struct {
	key string
}

// GenerateKey makes a new key and sets it
func (s *State) GenerateKey(keyParts ...interface{}) error {
	key, err := MakeKey(keyParts)

	if err != nil {
		return err
	}

	s.key = key

	return nil
}

// SetKey generates a key from the key parts passed in
// using a : separator
func (s *State) SetKey(key string) {
	s.key = key
}

// GetKey returns the state's key in its full format
func (s *State) GetKey() string {
	return s.key
}

// GetSplitKey the key split into multiple parts on :
func (s *State) GetSplitKey() []string {
	return SplitKey(s.key)
}
