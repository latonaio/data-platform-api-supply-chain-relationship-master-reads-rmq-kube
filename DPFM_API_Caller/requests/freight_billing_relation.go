package requests

type FreightBillingRelation struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID         int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightBillingID  int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	FreightPartner                           int     `json:"FreightPartner"`
	FreightBillToParty                       int     `json:"FreightBillToParty"`
	FreightBillFromParty                     int     `json:"FreightBillFromParty"`
	DefaultRelation                          *bool   `json:"DefaultRelation"`
	FreightBillToCountry                     string  `json:"FreightBillToCountry"`
	FreightBillFromCountry                   string  `json:"FreightBillFromCountry"`
	IsExportImport                           *bool   `json:"IsExportImport"`
	TransactionTaxCategory                   *string `json:"TransactionTaxCategory"`
	TransactionTaxClassification             *string `json:"TransactionTaxClassification"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}
