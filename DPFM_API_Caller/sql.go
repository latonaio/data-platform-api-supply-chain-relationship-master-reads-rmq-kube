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
		case "DeliveryRelation":
			func() {
				deliveryRelation = c.DeliveryRelation(mtx, input, output, errs, log)
			}()
		case "DeliveryRelations":
			func() {
				deliveryRelation = c.DeliveryRelations(mtx, input, output, errs, log)
			}()
		case "BillingRelation":
			func() {
				billingRelation = c.BillingRelation(mtx, input, output, errs, log)
			}()
		case "PaymentRelation":
			func() {
				paymentRelation = c.PaymentRelation(mtx, input, output, errs, log)
			}()
		case "DeliveryPlantRelation":
			func() {
				deliveryPlantRelation = c.DeliveryPlantRelation(mtx, input, output, errs, log)
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_general_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_general_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_transaction_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_relation_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_relation_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_billing_relation_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_payment_relation_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_plant_rel_prod
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_deliv_plant_rel_prod_mrp
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_stock_conf_plant_rel
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_stock_conf_plant_rel_pro
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_prod_plant_relation_data
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
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_prod_plant_rel_product
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
