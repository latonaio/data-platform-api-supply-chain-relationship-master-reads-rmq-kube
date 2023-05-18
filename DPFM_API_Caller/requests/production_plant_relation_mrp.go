package requests

type ProductionPlantRelationProductMRP struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	Product                                  string  `json:"Product"`
	ProductionPlantStorageLocation           *string `json:"ProductionPlantStorageLocation"`
	MRPType                                  *string `json:"MRPType"`
	MRPController                            *string `json:"MRPController"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}
