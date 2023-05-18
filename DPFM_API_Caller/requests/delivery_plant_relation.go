package requests

type DeliveryPlantRelation struct {
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	DefaultRelation                        *bool   `json:"DefaultRelation"`
	MRPType                                *string `json:"MRPType"`
	MRPController                          *string `json:"MRPController"`
	CreationDate                           *string `json:"CreationDate"`
	LastChangeDate                         *string `json:"LastChangeDate"`
	IsMarkedForDeletion                    *bool   `json:"IsMarkedForDeletion"`
}
