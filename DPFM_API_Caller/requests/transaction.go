package requests

type Transaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	Incoterms                 *string `json:"Incoterms"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	QuotationIsBlocked        *bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked            *bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked         *bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked          *bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
