package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-supply-chain-relationship-master-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *[]dpfm_api_output_formatter.General
	var transaction *[]dpfm_api_output_formatter.Transaction
	var deliveryRelation *[]dpfm_api_output_formatter.DeliveryRelation
	var billingRelation *[]dpfm_api_output_formatter.BillingRelation
	var paymentRelation *[]dpfm_api_output_formatter.PaymentRelation
	var deliveryPlantRelation *[]dpfm_api_output_formatter.DeliveryPlantRelation
	var deliveryPlantRelationProduct *[]dpfm_api_output_formatter.DeliveryPlantRelationProduct
	var deliveryPlantRelationProductMRPArea *[]dpfm_api_output_formatter.DeliveryPlantRelationProductMRPArea
	var stockConfPlantRelation *[]dpfm_api_output_formatter.StockConfPlantRelation
	var stockConfPlantRelationProduct *[]dpfm_api_output_formatter.StockConfPlantRelationProduct
	var productionPlantRelation *[]dpfm_api_output_formatter.ProductionPlantRelation
	var productionPlantRelationProductMRP *[]dpfm_api_output_formatter.ProductionPlantRelationProductMRP
	var freightRelation *[]dpfm_api_output_formatter.FreightRelation
	var freightTransaction *[]dpfm_api_output_formatter.FreightTransaction
	var freightBillingRelation *[]dpfm_api_output_formatter.FreightBillingRelation
	var freightPaymentRelation *[]dpfm_api_output_formatter.FreightPaymentRelation
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				general = c.Generals(mtx, input, output, errs, log)
			}()
		case "Transaction":
			func() {
				transaction = c.Transaction(mtx, input, output, errs, log)
			}()
		case "TransactionBySCRID":
			func() {
				transaction = c.TransactionBySCRID(mtx, input, output, errs, log)
			}()
		case "DeliveryRelation":
			func() {
				deliveryRelation = c.DeliveryRelation(mtx, input, output, errs, log)
			}()
		case "DeliveryRelations":
			func() {
				deliveryRelation = c.DeliveryRelations(mtx, input, output, errs, log)
			}()
		case "DeliveryRelationsBySCRID":
			func() {
				deliveryRelation = c.DeliveryRelationsBySCRID(mtx, input, output, errs, log)
			}()
		case "BillingRelation":
			func() {
				billingRelation = c.BillingRelation(mtx, input, output, errs, log)
			}()
		case "BillingRelationBySCRID":
			func() {
				billingRelation = c.BillingRelationBySCRID(mtx, input, output, errs, log)
			}()
		case "PaymentRelation":
			func() {
				paymentRelation = c.PaymentRelation(mtx, input, output, errs, log)
			}()
		case "PaymentRelationBySCRID":
			func() {
				paymentRelation = c.PaymentRelation(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelation":
			func() {
				deliveryPlantRelation = c.DeliveryPlantRelation(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelations":
			func() {
				deliveryPlantRelation = c.DeliveryPlantRelations(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelationsBySCRID":
			func() {
				deliveryPlantRelation = c.DeliveryPlantRelationsBySCRID(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelationProduct":
			func() {
				deliveryPlantRelationProduct = c.DeliveryPlantRelationProduct(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelationProductMRPArea":
			func() {
				deliveryPlantRelationProductMRPArea = c.DeliveryPlantRelationProductMRPArea(mtx, input, output, errs, log)
			}()
		case "StockConfPlantRelation":
			func() {
				stockConfPlantRelation = c.StockConfPlantRelation(mtx, input, output, errs, log)
			}()
		case "StockConfPlantRelationProduct":
			func() {
				stockConfPlantRelationProduct = c.StockConfPlantRelationProduct(mtx, input, output, errs, log)
			}()
		case "ProductionPlantRelation":
			func() {
				productionPlantRelation = c.ProductionPlantRelation(mtx, input, output, errs, log)
			}()
		case "ProductionPlantRelationProductMRP":
			func() {
				productionPlantRelationProductMRP = c.ProductionPlantRelationProductMRP(mtx, input, output, errs, log)
			}()
		case "FreightRelation":
			func() {
				freightRelation = c.FreightRelation(mtx, input, output, errs, log)
			}()
		case "FreightTransaction":
			func() {
				freightTransaction = c.FreightTransaction(mtx, input, output, errs, log)
			}()
		case "FreightBillingRelation":
			func() {
				freightBillingRelation = c.FreightBillingRelation(mtx, input, output, errs, log)
			}()
		case "FreightPaymentRelation":
			func() {
				freightPaymentRelation = c.FreightPaymentRelation(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:                             general,
		Transaction:                         transaction,
		DeliveryRelation:                    deliveryRelation,
		BillingRelation:                     billingRelation,
		PaymentRelation:                     paymentRelation,
		DeliveryPlantRelation:               deliveryPlantRelation,
		DeliveryPlantRelationProduct:        deliveryPlantRelationProduct,
		DeliveryPlantRelationProductMRPArea: deliveryPlantRelationProductMRPArea,
		StockConfPlantRelation:              stockConfPlantRelation,
		StockConfPlantRelationProduct:       stockConfPlantRelationProduct,
		ProductionPlantRelation:             productionPlantRelation,
		ProductionPlantRelationProductMRP:   productionPlantRelationProductMRP,
		FreightRelation:                     freightRelation,
		FreightTransaction:                  freightTransaction,
		FreightBillingRelation:              freightBillingRelation,
		FreightPaymentRelation:              freightPaymentRelation,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_general_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Generals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := "WHERE 1 = 1"
	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.General.IsMarkedForDeletion)
	}

	idWhere := ""
	if input.General.Buyer != 0 {
		idWhere = fmt.Sprintf("\nAND Buyer = %d ", input.General.Buyer)
	}
	if input.General.Seller != 0 {
		idWhere = fmt.Sprintf("%s\nAND Seller = %d ", idWhere, input.General.Seller)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_general_data
		` + where + idWhere + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Transaction(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Transaction {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_transaction_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToTransaction(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) TransactionBySCRID(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Transaction {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_transaction_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToTransaction(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	deliveryRelation := input.General.DeliveryRelation

	cnt := 0
	for _, v := range deliveryRelation {
		args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipDeliveryID, buyer, seller, v.DeliverToParty, v.DeliverFromParty)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryRelations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryRelation {
	where := " WHERE 1 != 1"
	for _, v := range input.General.DeliveryRelation {
		where = fmt.Sprintf("%s OR SupplyChainRelationshipID = %d ", where, v.SupplyChainRelationshipID)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_relation_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryRelationsBySCRID(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryRelation {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT d.*,
       			bBuyer.BusinessPartnerFullName AS BuyerName,
       			bSeller.BusinessPartnerFullName AS SellerName,
       			bDeliveryTo.BusinessPartnerFullName AS bDeliveryToPartyName,
       			bDeliveryFrom.BusinessPartnerFullName AS bDeliveryFromPartyName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_relation_data AS d
		LEFT JOIN data_platform_business_partner_general_data AS bBuyer ON d.Buyer = bBuyer.BusinessPartner
		LEFT JOIN data_platform_business_partner_general_data AS bSeller ON d.Seller = bSeller.BusinessPartner
		LEFT JOIN data_platform_business_partner_general_data AS bDeliveryTo ON d.DeliverToParty = bDeliveryTo.BusinessPartner
		LEFT JOIN data_platform_business_partner_general_data AS bDeliveryFrom ON d.DeliverFromParty = bDeliveryFrom.BusinessPartner
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BillingRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BillingRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	billingRelation := input.General.BillingRelation

	cnt := 0
	for _, v := range billingRelation {
		args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipBillingID, buyer, seller, v.BillToParty, v.BillFromParty)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_billing_relation_data
		WHERE (SupplyChainRelationshipID, supplyChainRelationshipBillingID, Buyer, Seller, BillToParty, BillFromParty) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBillingRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BillingRelationBySCRID(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BillingRelation {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_billing_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBillingRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	billingRelation := input.General.BillingRelation

	cnt := 0
	for _, v := range billingRelation {
		paymentRelation := v.PaymentRelation
		for _, w := range paymentRelation {
			args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipBillingID, w.SupplyChainRelationshipPaymentID, buyer, seller, v.BillToParty, v.BillFromParty, w.Payer, w.Payee)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_payment_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID, Buyer, Seller, BillToParty, BillFromParty, Payer, Payee) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentRelationBySCRID(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentRelation {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_payment_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryPlantRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryPlantRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	deliveryRelation := input.General.DeliveryRelation

	cnt := 0
	for _, v := range deliveryRelation {
		for _, w := range v.DeliveryPlantRelation {
			args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipDeliveryID, w.SupplyChainRelationshipDeliveryPlantID, buyer, seller, v.DeliverToParty, v.DeliverFromParty, w.DeliverToPlant, w.DeliverFromPlant)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_plant_rel_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryPlantRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryPlantRelations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryPlantRelation {
	where := " WHERE 1 != 1"
	for _, v := range input.General.DeliveryRelation[0].DeliveryPlantRelation {
		where = fmt.Sprintf("%s OR SupplyChainRelationshipID = %d ", where, v.SupplyChainRelationshipID)
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_plant_relation_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryPlantRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryPlantRelationsBySCRID(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryPlantRelation {
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller

	rows, err := c.db.Query(
		`SELECT d.*,
       			bBuyer.BusinessPartnerFullName AS BuyerName,
       			bSeller.BusinessPartnerFullName AS SellerName,
       			pDeliveryToPlant.PlantFullName AS DeliveryToPlantName,
       			pDeliveryFromPlant.PlantFullName AS DeliveryFromPlantName,
				bDeliveryTo.BusinessPartnerFullName AS DeliveryToPartyName,
       			bDeliveryFrom.BusinessPartnerFullName AS DeliveryFromPartyName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_plant_relation_data AS d
		LEFT JOIN data_platform_business_partner_general_data AS bBuyer ON d.Buyer = bBuyer.BusinessPartner
		LEFT JOIN data_platform_business_partner_general_data AS bSeller ON d.Seller = bSeller.BusinessPartner
		LEFT JOIN data_platform_plant_general_data AS pDeliveryToPlant ON d.DeliverToPlant = pDeliveryToPlant.Plant
		LEFT JOIN data_platform_plant_general_data AS pDeliveryFromPlant ON d.DeliverFromPlant = pDeliveryFromPlant.Plant
		LEFT JOIN data_platform_business_partner_general_data AS bDeliveryTo ON d.DeliverToParty = bDeliveryTo.BusinessPartner
		LEFT JOIN data_platform_business_partner_general_data AS bDeliveryFrom ON d.DeliverFromParty = bDeliveryFrom.BusinessPartner
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?,?,?);`, supplyChainRelationshipID, buyer, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryPlantRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryPlantRelationProduct(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryPlantRelationProduct {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	deliveryRelation := input.General.DeliveryRelation

	cnt := 0
	for _, v := range deliveryRelation {
		for _, w := range v.DeliveryPlantRelation {
			for _, x := range w.DeliveryPlantRelationProduct {
				args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipDeliveryID, w.SupplyChainRelationshipDeliveryPlantID, buyer, seller, v.DeliverToParty, v.DeliverFromParty, w.DeliverToPlant, w.DeliverFromPlant, x.Product)
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_plant_rel_prod
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant, Product) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryPlantRelationProduct(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) DeliveryPlantRelationProductMRPArea(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DeliveryPlantRelationProductMRPArea {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	deliveryRelation := input.General.DeliveryRelation

	cnt := 0
	for _, v := range deliveryRelation {
		for _, w := range v.DeliveryPlantRelation {
			for _, x := range w.DeliveryPlantRelationProduct {
				for _, y := range x.DeliveryPlantRelationProductMRPArea {
					args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipDeliveryID, w.SupplyChainRelationshipDeliveryPlantID, buyer, seller, v.DeliverToParty, v.DeliverFromParty, w.DeliverToPlant, w.DeliverFromPlant, x.Product, y.MRPArea)
				}
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_deliv_plant_rel_prod_mrp
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant, Product, MRPArea) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToDeliveryPlantRelationProductMRPArea(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StockConfPlantRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StockConfPlantRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	stockConfPlantRelation := input.General.StockConfPlantRelation

	cnt := 0
	for _, v := range stockConfPlantRelation {
		args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipStockConfPlantID, buyer, seller, v.StockConfirmationBusinessPartner, v.StockConfirmationPlant)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_stock_conf_plant_rel
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipStockConfPlantID, Buyer, Seller, StockConfirmationBusinessPartner, StockConfirmationPlant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStockConfPlantRelation(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StockConfPlantRelationProduct(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StockConfPlantRelationProduct {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	stockConfPlantRelation := input.General.StockConfPlantRelation

	cnt := 0
	for _, v := range stockConfPlantRelation {
		stockConfPlantRelationProduct := v.StockConfPlantRelationProduct
		for _, w := range stockConfPlantRelationProduct {
			args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipStockConfPlantID, buyer, seller, v.StockConfirmationBusinessPartner, v.StockConfirmationPlant, w.Product)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_stock_conf_plant_rel_pro
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipStockConfPlantID, Buyer, Seller, StockConfirmationBusinessPartner, StockConfirmationPlant, Product) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStockConfPlantRelationProduct(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductionPlantRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductionPlantRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	productionPlantRelation := input.General.ProductionPlantRelation

	cnt := 0
	for _, v := range productionPlantRelation {
		args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipProductionPlantID, buyer, seller, v.ProductionPlantBusinessPartner, v.ProductionPlant)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_prod_plant_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Buyer, Seller, ProductionPlantBusinessPartner, ProductionPlant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductionPlantRelation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductionPlantRelationProductMRP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductionPlantRelationProductMRP {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	productionPlantRelation := input.General.ProductionPlantRelation

	cnt := 0
	for _, v := range productionPlantRelation {
		productionPlantRelationProductMRP := v.ProductionPlantRelationProductMRP
		for _, w := range productionPlantRelationProductMRP {
			args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipProductionPlantID, buyer, seller, v.ProductionPlantBusinessPartner, v.ProductionPlant, w.Product)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_prod_plant_rel_product
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Buyer, Seller, ProductionPlantBusinessPartner, ProductionPlant, Product) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductionPlantRelationProductMRP(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FreightRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FreightRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	buyer := input.General.Buyer
	seller := input.General.Seller
	freightRelation := input.General.FreightRelation
	cnt := 0
	for _, v := range freightRelation {
		args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipFreightID, buyer, seller, v.FreightPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_freight_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionFreightID, Buyer, Seller, FreightPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToFreightRelation(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FreightTransaction(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FreightTransaction {
	var args []interface{}
	freightRelation := input.General.FreightRelation
	cnt := 0
	for _, v := range freightRelation {
		args = append(args, v.SupplyChainRelationshipID, v.SupplyChainRelationshipFreightID, v.Buyer, v.Seller, v.FreightPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_freight_transaction_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionFreightID, Buyer, Seller, FreightPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToFreightTransaction(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FreightBillingRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FreightBillingRelation {
	var args []interface{}
	supplyChainRelationshipID := input.General.SupplyChainRelationshipID
	freightRelation := input.General.FreightRelation
	cnt := 0
	for _, v := range freightRelation {
		freightBillingReration := v.FreightBillingRelation
		for _, w := range freightBillingReration {
			args = append(args, supplyChainRelationshipID, v.SupplyChainRelationshipFreightID, w.SupplyChainRelationshipFreightBillingID, v.Buyer, v.Seller, v.FreightPartner, w.FreightBillToParty, w.FreightBillFromParty)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_freight_billing_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionFreightID, SupplyChainRelationshipFreightBillingID, Buyer, Seller, FreightPartner, FreightBillToParty, FreightBillFromParty) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToFreightBillingRelation(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FreightPaymentRelation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FreightPaymentRelation {
	var args []interface{}
	buyer := input.General.Buyer
	seller := input.General.Seller
	freightRelation := input.General.FreightRelation
	cnt := 0
	for _, v := range freightRelation {
		freightBillingReration := v.FreightBillingRelation
		for _, w := range freightBillingReration {
			freightPaymentReration := w.FreightPaymentRelation
			for _, x := range freightPaymentReration {
				args = append(args, w.SupplyChainRelationshipID, w.SupplyChainRelationshipFreightID, w.SupplyChainRelationshipFreightBillingID,
					x.SupplyChainRelationshipFreightPaymentID, w.Buyer, w.Seller, w.FreightPartner, w.FreightBillToParty, w.FreightBillFromParty, buyer, seller)
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_freight_billing_relation_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipFreightID, SupplyChainRelationshipFreightBillingID, SupplyChainRelationshipFreightPaymentID, Buyer, Seller, FreightPartner, FreightBillToParty, FreightBillFromParty, FreightPayer, FreightPayee) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToFreightPaymentRelation(rows)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
