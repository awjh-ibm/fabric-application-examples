/*
SPDX-License-Identifier: Apache-2.0
*/
package papernet

import (
	"ledger"
	"ledger/datautils"
	"reflect"
)

// CommercialPaperState set of enums for the handling of
// commercial paper states
type CommercialPaperState int

// Enum values for CommercialPaperState
const (
	ISSUED CommercialPaperState = iota + 1
	TRADING
	REDEEMED
)

// CommercialPaper a struct to represent a CommercialPaper
// Embeds ledger.State to enable use in world state handling
type CommercialPaper struct {
	ledger.State
	Issuer           string               `json:"issuer"`
	Owner            string               `json:"owner"`
	PaperNumber      int                  `json:"paperNumber"`
	IssueDateTime    string               `json:"issueDateTime"`
	MaturityDateTime string               `json:"maturityDateTime"`
	FaceValue        int                  `json:"faceValue"`
	CurrentState     CommercialPaperState `json:"currentState"`
}

// GetType returns the identifying type for commercial
// paper
func (cp *CommercialPaper) GetType() string {
	return "org.papernet.commercialpaper"
}

// GetCurrentState returns the current state of the
// commercial paper. Returns interface to meet StateInterface
func (cp *CommercialPaper) GetCurrentState() interface{} {
	return cp.CurrentState
}

// SetIssued sets the current state of the commercial
// paper to be ISSUED
func (cp *CommercialPaper) SetIssued() {
	cp.CurrentState = ISSUED
}

// SetTrading sets the current state of the commercial
// paper to be TRADING
func (cp *CommercialPaper) SetTrading() {
	cp.CurrentState = TRADING
}

// SetRedeemed sets the current state of the commercial
// paper to be REDEEMED
func (cp *CommercialPaper) SetRedeemed() {
	cp.CurrentState = REDEEMED
}

// IsIssued returns true if commercial paper current state
// is ISSUED
func (cp *CommercialPaper) IsIssued() bool {
	return cp.CurrentState == ISSUED
}

// IsTrading returns true if commercial paper current state
// is TRADING
func (cp *CommercialPaper) IsTrading() bool {
	return cp.CurrentState == TRADING
}

// IsRedeemed returns true if commercial paper current state
// is REDEEMED
func (cp *CommercialPaper) IsRedeemed() bool {
	return cp.CurrentState == REDEEMED
}

// Deserialize converts JSON bytes to a CommercialPaper
func Deserialize(data []byte) (*CommercialPaper, error) {
	iFace, err := datautils.DeserializeToType(data, reflect.TypeOf(CommercialPaper{}))

	if err != nil {
		return nil, err
	}

	return iFace.(*CommercialPaper), nil
}
