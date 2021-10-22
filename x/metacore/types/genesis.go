package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		SendList:              []*Send{},
		NodeAccountList:       []*NodeAccount{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in send
	sendIndexMap := make(map[string]bool)

	for _, elem := range gs.SendList {
		if _, ok := sendIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for send")
		}
		sendIndexMap[elem.Index] = true
	}


	// Check for duplicated index in nodeAccount
	nodeAccountIndexMap := make(map[string]bool)

	for _, elem := range gs.NodeAccountList {
		if _, ok := nodeAccountIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for nodeAccount")
		}
		nodeAccountIndexMap[elem.Index] = true
	}


	return nil
}
