package requests

type DeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID                 int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int      `json:"Buyer"`
	Seller                                    int      `json:"Seller"`
	DeliverToParty                            int      `json:"DeliverToParty"`
	DeliverFromParty                          int      `json:"DeliverFromParty"`
	DeliverToPlant                            string   `json:"DeliverToPlant"`
	DeliverFromPlant                          string   `json:"DeliverFromPlant"`
	Product                                   string   `json:"Product"`
	DeliverToPlantStorageLocation             string   `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string   `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string   `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32 `json:"QuantityPerPackage"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
    MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
    MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
    StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
    StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
    MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
    MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
    DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration            	  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit			  *string  `json:"StandardDeliveryDurationUnit"`
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}
