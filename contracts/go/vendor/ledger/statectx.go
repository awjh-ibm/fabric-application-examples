/*
SPDX-License-Identifier: Apache-2.0
*/

package ledger

import (
	"encoding/json"
	"errors"
	"ledger/datautils"
	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("STATECTX")

// StateTransactionContext context for handling generic world
// state interactions
type StateTransactionContext struct {
	stub           shim.ChaincodeStubInterface
	cid            cid.ClientIdentity
	currentList    string
	supportedTypes map[string]reflect.Type
}

// SetCurrentList sets the name of the list to be used in the world state for the context
func (stc *StateTransactionContext) SetCurrentList(listName string) {
	stc.currentList = listName
}

// SetStub allows the stub to be set in the context
func (stc *StateTransactionContext) SetStub(stub shim.ChaincodeStubInterface) {
	stc.stub = stub
}

// SetClientIdentity allows the cid to be set in the context
func (stc *StateTransactionContext) SetClientIdentity(cid cid.ClientIdentity) {
	stc.cid = cid
}

// AddState adds a state to the list of states. Creates a new state in the world state
// with appropriate composite key.
func (stc *StateTransactionContext) AddState(state StateInterface) error {
	key, err := stc.createCompositeKey(state.GetSplitKey())

	if err != nil {
		return err
	}

	data, err := stc.serializeWithType(state)

	return stc.stub.PutState(key, data)
}

// GetState gets a state from the list using suplied key. Worldstate data is deserialized
// into a State type
func (stc *StateTransactionContext) GetState(key string) (StateInterface, error) {
	splitKey := SplitKey(key)
	compKey, err := stc.createCompositeKey(splitKey)

	if err != nil {
		return nil, err
	}

	data, err := stc.stub.GetState(compKey)

	if err != nil {
		return nil, err
	}

	state, err := datautils.Deserialize(data, stc.supportedTypes)

	if err != nil {
		return nil, err
	}

	logger.Infof("KEY %s, THE DATA RETURNED %s", compKey, data)

	si := state.(StateInterface)
	si.SetKey(key)

	return si, nil
}

// UpdateState updates a state in the list of states. Puts the new state in the world state
// with the appropriate composite key.
func (stc *StateTransactionContext) UpdateState(state StateInterface) error {
	key, err := stc.createCompositeKey(state.GetSplitKey())

	if err != nil {
		return err
	}

	data, err := stc.serializeWithType(state)

	logger.Infof("KEY %s. THE DATA BEING WRITTEN %s", key, string(data))

	return stc.stub.PutState(key, data)
}

// Use adds a state type to the state list
func (stc *StateTransactionContext) Use(state StateInterface) {
	if stc.supportedTypes == nil {
		stc.supportedTypes = make(map[string]reflect.Type)
	}

	stc.supportedTypes[state.GetType()] = reflect.TypeOf(state).Elem()
}

func (stc *StateTransactionContext) serializeWithType(state StateInterface) ([]byte, error) {
	serializedState, err := datautils.Serialize(state)

	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})

	err = json.Unmarshal(serializedState, &jsonMap)

	if err != nil {
		return nil, err
	}

	jsonMap["type"] = state.GetType()

	return datautils.Serialize(jsonMap)
}

func (stc *StateTransactionContext) createCompositeKey(splitKey []string) (string, error) {
	key, err := stc.stub.CreateCompositeKey(stc.currentList, splitKey)

	if err != nil {
		return "", errors.New("Failed to generate composite key")
	}

	return key, nil
}
