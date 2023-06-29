package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	General           General  `json:"SupplyChainRelationship"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type General struct {
	SupplyChainRelationshipID int                       `json:"SupplyChainRelationshipID"`
	Buyer                     int                       `json:"Buyer"`
	Seller                    int                       `json:"Seller"`
	CreationDate              *string                   `json:"CreationDate"`
	LastChangeDate            *string                   `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool                     `json:"IsMarkedForDeletion"`
	Transaction               []Transaction             `json:"Transaction"`
	DeliveryRelation          []DeliveryRelation        `json:"DeliveryRelation"`
	BillingRelation           []BillingRelation         `json:"BillingRelation"`
	ProductionPlantRelation   []ProductionPlantRelation `json:"ProductionPlantRelation"`
	StockConfPlantRelation    []StockConfPlantRelation  `json:"StockConfPlantRelation"`
	FreightRelation           []FreightRelation         `json:"FreightRelation"`
}

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

type DeliveryRelation struct {
	SupplyChainRelationshipID         int                     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int                     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int                     `json:"Buyer"`
	Seller                            int                     `json:"Seller"`
	DeliverToParty                    int                     `json:"DeliverToParty"`
	DeliverFromParty                  int                     `json:"DeliverFromParty"`
	DefaultRelation                   *bool                   `json:"DefaultRelation"`
	CreationDate                      *string                 `json:"CreationDate"`
	LastChangeDate                    *string                 `json:"LastChangeDate"`
	IsMarkedForDeletion               *bool                   `json:"IsMarkedForDeletion"`
	DeliveryPlantRelation             []DeliveryPlantRelation `json:"DeliveryPlantRelation"`
}

type BillingRelation struct {
	SupplyChainRelationshipID        int               `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int               `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int               `json:"Buyer"`
	Seller                           int               `json:"Seller"`
	BillToParty                      int               `json:"BillToParty"`
	BillFromParty                    int               `json:"BillFromParty"`
	DefaultRelation                  *bool             `json:"DefaultRelation"`
	BillToCountry                    *string           `json:"BillToCountry"`
	BillFromCountry                  *string           `json:"BillFromCountry"`
	IsExportImport                   *bool             `json:"IsExportImport"`
	TransactionTaxCategory           *string           `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string           `json:"TransactionTaxClassification"`
	CreationDate                     *string           `json:"CreationDate"`
	LastChangeDate                   *string           `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool             `json:"IsMarkedForDeletion"`
	PaymentRelation                  []PaymentRelation `json:"PaymentRelation"`
}

type PaymentRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	Payer                            int     `json:"Payer"`
	Payee                            int     `json:"Payee"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	PayerHouseBank                   *string `json:"PayerHouseBank"`
	PayerHouseBankAccount            *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank                   *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount            *string `json:"PayeeHouseBankAccount"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelation struct {
	SupplyChainRelationshipID              int                            `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int                            `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int                            `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int                            `json:"Buyer"`
	Seller                                 int                            `json:"Seller"`
	DeliverToParty                         int                            `json:"DeliverToParty"`
	DeliverFromParty                       int                            `json:"DeliverFromParty"`
	DeliverToPlant                         string                         `json:"DeliverToPlant"`
	DeliverFromPlant                       string                         `json:"DeliverFromPlant"`
	DefaultRelation                        *bool                          `json:"DefaultRelation"`
	MRPType                                *string                        `json:"MRPType"`
	MRPController                          *string                        `json:"MRPController"`
	CreationDate                           *string                        `json:"CreationDate"`
	LastChangeDate                         *string                        `json:"LastChangeDate"`
	IsMarkedForDeletion                    *bool                          `json:"IsMarkedForDeletion"`
	DeliveryPlantRelationProduct           []DeliveryPlantRelationProduct `json:"DeliveryPlantRelationProduct"`
}

type DeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID                 int                                   `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int                                   `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int                                   `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int                                   `json:"Buyer"`
	Seller                                    int                                   `json:"Seller"`
	DeliverToParty                            int                                   `json:"DeliverToParty"`
	DeliverFromParty                          int                                   `json:"DeliverFromParty"`
	DeliverToPlant                            string                                `json:"DeliverToPlant"`
	DeliverFromPlant                          string                                `json:"DeliverFromPlant"`
	Product                                   string                                `json:"Product"`
	DeliverToPlantStorageLocation             string                                `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string                                `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string                                `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32                              `json:"QuantityPerPackage"`
	MRPType                                   *string                               `json:"MRPType"`
	MRPController                             *string                               `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32                              `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int                                  `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string                               `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32                              `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32                              `json:"SafetyDuration"`
	SafetyDurationUnit                        *string                               `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32                              `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32                              `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32                              `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32                              `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32                              `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32                              `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32                              `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32                              `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool                                 `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32                              `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string                               `json:"StandardDeliveryDurationUnit"`
	CreationDate                              *string                               `json:"CreationDate"`
	LastChangeDate                            *string                               `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool                                 `json:"IsMarkedForDeletion"`
	DeliveryPlantRelationProductMRPArea       []DeliveryPlantRelationProductMRPArea `json:"DeliveryPlantRelationProductMRPArea"`
}

type DeliveryPlantRelationProductMRPArea struct {
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
	MRPArea                                   string   `json:"MRPArea"`
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
	StandardDeliveryDuration                  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type StockConfPlantRelation struct {
	SupplyChainRelationshipID               int                             `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int                             `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int                             `json:"Buyer"`
	Seller                                  int                             `json:"Seller"`
	StockConfirmationBusinessPartner        int                             `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string                          `json:"StockConfirmationPlant"`
	DefaultRelation                         *bool                           `json:"DefaultRelation"`
	MRPType                                 *string                         `json:"MRPType"`
	MRPController                           *string                         `json:"MRPController"`
	CreationDate                            *string                         `json:"CreationDate"`
	LastChangeDate                          *string                         `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool                           `json:"IsMarkedForDeletion"`
	StockConfPlantRelationProduct           []StockConfPlantRelationProduct `json:"StockConfPlantRelationProduct"`
}

type StockConfPlantRelationProduct struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	StockConfirmationBusinessPartner        int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string  `json:"StockConfirmationPlant"`
	Product                                 string  `json:"Product"`
	MRPType                                 *string `json:"MRPType"`
	MRPController                           *string `json:"MRPController"`
	CreationDate                            *string `json:"CreationDate"`
	LastChangeDate                          *string `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
}

type ProductionPlantRelation struct {
	SupplyChainRelationshipID                int                                 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int                                 `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int                                 `json:"Buyer"`
	Seller                                   int                                 `json:"Seller"`
	ProductionPlantBusinessPartner           int                                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string                              `json:"ProductionPlant"`
	DefaultRelation                          *bool                               `json:"DefaultRelation"`
	MRPType                                  *string                             `json:"MRPType"`
	MRPController                            *string                             `json:"MRPController"`
	CreationDate                             *string                             `json:"CreationDate"`
	LastChangeDate                           *string                             `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool                               `json:"IsMarkedForDeletion"`
	ProductionPlantRelationProductMRP        []ProductionPlantRelationProductMRP `json:"ProductionPlantRelationProductMRP"`
}

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

type FreightRelation struct {
	SupplyChainRelationshipID        int                      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID int                      `json:"SupplyChainRelationshipFreightID"`
	Buyer                            int                      `json:"Buyer"`
	Seller                           int                      `json:"Seller"`
	FreightPartner                   int                      `json:"FreightPartner"`
	CreationDate                     *string                  `json:"CreationDate"`
	LastChangeDate                   *string                  `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool                    `json:"IsMarkedForDeletion"`
	FreightTransaction               []FreightTransaction     `json:"FreightTransaction"`
	FreightBillingRelation           []FreightBillingRelation `json:"FreightBillingRelation"`
	FreightPaymentRelation           []FreightPaymentRelation `json:"FreightPaymentRelation"`
}

type FreightTransaction struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID int     `json:"SupplyChainRelationshipFreightID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	FreightPartner                   int     `json:"FreightPartner"`
	TransactionCurrency              *string `json:"TransactionCurrency"`
	PaymentTerms                     *string `json:"PaymentTerms"`
	PaymentMethod                    *string `json:"PaymentMethod"`
	Incoterms                        *string `json:"Incoterms"`
	AccountAssignmentGroup           *string `json:"AccountAssignmentGroup"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	QuotationIsBlocked               *bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked                   *bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked                *bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked                 *bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type FreightBillingRelation struct {
	SupplyChainRelationshipID               int                      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        int                      `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID int                      `json:"SupplyChainRelationshipBillingID"`
	Buyer                                   int                      `json:"Buyer"`
	Seller                                  int                      `json:"Seller"`
	FreightPartner                          int                      `json:"FreightPartner"`
	FreightBillToParty                      int                      `json:"FreightBillToParty"`
	FreightBillFromParty                    int                      `json:"FreightBillFromParty"`
	DefaultRelation                         *bool                    `json:"DefaultRelation"`
	FreightBillToCountry                    string                   `json:"FreightBillToCountry"`
	FreightBillFromCountry                  string                   `json:"FreightBillFromCountry"`
	IsExportImport                          *bool                    `json:"IsExportImport"`
	TransactionTaxCategory                  *string                  `json:"TransactionTaxCategory"`
	TransactionTaxClassification            *string                  `json:"TransactionTaxClassification"`
	CreationDate                            *string                  `json:"CreationDate"`
	LastChangeDate                          *string                  `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool                    `json:"IsMarkedForDeletion"`
	FreightPaymentRelation                  []FreightPaymentRelation `json:"FreightPaymentRelation"`
}

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
