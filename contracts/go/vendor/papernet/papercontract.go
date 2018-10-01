/*
SPDX-License-Identifier: Apache-2.0
*/

package papernet

import (
	"fmt"
	"ledger"

	"github.com/hyperledger/fabric/core/chaincode/contractapi"
)

// CommercialPaperContract business logic for handling a commercial paper
type CommercialPaperContract struct {
	contractapi.Contract
}

// Issue a commercial paper
func (cpc *CommercialPaperContract) Issue(ctx *PaperTransactionContext, issuer string, paperNumber int, issueDateTime string, maturityDateTime string, faceValue int) error {
	cp := new(CommercialPaper)
	cp.GenerateKey(issuer, paperNumber)
	cp.Issuer = issuer
	cp.Owner = issuer
	cp.PaperNumber = paperNumber
	cp.IssueDateTime = issueDateTime
	cp.MaturityDateTime = maturityDateTime
	cp.FaceValue = faceValue
	cp.SetIssued()

	return ctx.AddPaper(cp)
}

// Buy a commercial paper
func (cpc *CommercialPaperContract) Buy(ctx *PaperTransactionContext, issuer string, paperNumber int, currentOwner string, newOwner string, price int, purchaseDateTime string) error {
	cpKey, err := ledger.MakeKey(issuer, paperNumber)

	if err != nil {
		return err
	}

	cp, err := ctx.GetPaper(cpKey)

	if err != nil {
		return err
	}

	if cp.Owner != currentOwner {
		return fmt.Errorf("Paper %s%d is not owned by %s", issuer, paperNumber, currentOwner)
	}

	if cp.IsIssued() {
		cp.SetTrading()
	}

	if !cp.IsTrading() {
		return fmt.Errorf("Paper %s%d is not trading. Current state = %v", issuer, paperNumber, cp.GetCurrentState())
	}

	cp.Owner = newOwner

	return ctx.UpdatePaper(cp)
}

// Redeem a commercial paper
func (cpc *CommercialPaperContract) Redeem(ctx *PaperTransactionContext, issuer string, paperNumber int, redeemingOwner string, redeemDateTime string) error {
	cpKey, err := ledger.MakeKey(issuer, paperNumber)

	if err != nil {
		return err
	}

	cp, err := ctx.GetPaper(cpKey)

	if err != nil {
		return err
	}

	if cp.Owner != redeemingOwner {
		return fmt.Errorf("Paper %s%d is not owned by %s", issuer, paperNumber, redeemingOwner)
	}

	if cp.IsRedeemed() {
		return fmt.Errorf("Paper %s%d is already redeemed", issuer, paperNumber)
	}

	return ctx.UpdatePaper(cp)
}
