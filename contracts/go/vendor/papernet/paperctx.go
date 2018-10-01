/*
SPDX-License-Identifier: Apache-2.0
*/

package papernet

import (
	"ledger"
)

// PaperTransactionContext handles management of papers in the world state
type PaperTransactionContext struct {
	ledger.StateTransactionContext
}

func (ptc *PaperTransactionContext) prepare() {
	ptc.SetCurrentList("org.papernet.commercialpaperlist")
	ptc.Use(new(CommercialPaper))
}

// AddPaper adds a paper to the world state
func (ptc *PaperTransactionContext) AddPaper(paper *CommercialPaper) error {
	ptc.prepare()
	return ptc.AddState(paper)
}

// GetPaper gets a paper to the world state
func (ptc *PaperTransactionContext) GetPaper(paperKey string) (*CommercialPaper, error) {
	ptc.prepare()

	si, err := ptc.GetState(paperKey)

	if err != nil {
		return nil, err
	}

	return si.(*CommercialPaper), nil
}

// UpdatePaper adds a paper to the world state
func (ptc *PaperTransactionContext) UpdatePaper(paper *CommercialPaper) error {
	ptc.prepare()

	return ptc.UpdateState(paper)
}
