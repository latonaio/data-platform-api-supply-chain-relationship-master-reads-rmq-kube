package requests

type FreightPaymentRelation struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        int     `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID int     `json:"SupplyChainRelationshipFreightBillingID"`
	SupplyChainRelationshipFreightPaymentID int     `json:"SupplyChainRelationshipFreightPaymentID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	FreightPartner                          int     `json:"FreightPartner"`
	FreightBillToParty                      int     `json:"FreightBillToParty"`
	FreightBillFromParty                    int     `json:"FreightBillFromParty"`
	FreightPayer                            int     `json:"FreightPayer"`
	FreightPayee                            int     `json:"FreightPayee"`
	DefaultRelation                         *bool   `json:"DefaultRelation"`
	PayerHouseBank                          *string `json:"PayerHouseBank"`
	PayerHouseBankAccount                   *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank                          *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount                   *string `json:"PayeeHouseBankAccount"`
	CreationDate                            *string `json:"CreationDate"`
	LastChangeDate                          *string `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
}
