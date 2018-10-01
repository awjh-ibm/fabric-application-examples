/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"papernet"

	"github.com/hyperledger/fabric/core/chaincode/contractapi"
)

func main() {
	cpc := new(papernet.CommercialPaperContract)
	cpc.SetTransactionContextHandler(new(papernet.PaperTransactionContext))
	cpc.SetNamespace("org.papernet.commercialpaper")

	if err := contractapi.CreateNewChaincode(cpc); err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
