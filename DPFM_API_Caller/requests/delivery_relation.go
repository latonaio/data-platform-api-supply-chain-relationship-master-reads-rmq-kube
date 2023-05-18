package requests

type DeliveryRelation struct {
	SupplyChainRelationshipID         int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int     `json:"Buyer"`
	Seller                            int     `json:"Seller"`
	DeliverToParty                    int     `json:"DeliverToParty"`
	DeliverFromParty                  int     `json:"DeliverFromParty"`
	DefaultRelation                   *bool   `json:"DefaultRelation"`
	CreationDate                      *string `json:"CreationDate"`
	LastChangeDate                    *string `json:"LastChangeDate"`
	IsMarkedForDeletion               *bool   `json:"IsMarkedForDeletion"`
}
