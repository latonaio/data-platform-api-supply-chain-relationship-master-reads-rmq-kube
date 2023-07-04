package dpfm_api_output_formatter

import (
	"data-platform-api-supply-chain-relationship-master-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

		data := pm
		general = append(general, General{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			CreationDate:              data.CreationDate,
			LastChangeDate:            data.LastChangeDate,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}

func ConvertToTransaction(rows *sql.Rows) (*[]Transaction, error) {
	defer rows.Close()
	transaction := make([]Transaction, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Transaction{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.TransactionCurrency,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Incoterms,
			&pm.AccountAssignmentGroup,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.QuotationIsBlocked,
			&pm.OrderIsBlocked,
			&pm.DeliveryIsBlocked,
			&pm.BillingIsBlocked,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &transaction, err
		}

		data := pm
		transaction = append(transaction, Transaction{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			TransactionCurrency:       data.TransactionCurrency,
			PaymentTerms:              data.PaymentTerms,
			PaymentMethod:             data.PaymentMethod,
			Incoterms:                 data.Incoterms,
			AccountAssignmentGroup:    data.AccountAssignmentGroup,
			CreationDate:              data.CreationDate,
			LastChangeDate:            data.LastChangeDate,
			QuotationIsBlocked:        data.QuotationIsBlocked,
			OrderIsBlocked:            data.OrderIsBlocked,
			DeliveryIsBlocked:         data.DeliveryIsBlocked,
			BillingIsBlocked:          data.BillingIsBlocked,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &transaction, nil
	}

	return &transaction, nil
}

func ConvertToDeliveryRelation(rows *sql.Rows) (*[]DeliveryRelation, error) {
	defer rows.Close()
	deliveryRelation := make([]DeliveryRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := DeliveryRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DefaultRelation,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
			&pm.BuyerName,
			&pm.SellerName,
			&pm.DeliverToPartyName,
			&pm.DeliverFromPartyName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &deliveryRelation, err
		}

		data := pm
		deliveryRelation = append(deliveryRelation, DeliveryRelation{
			SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
			Buyer:                             data.Buyer,
			Seller:                            data.Seller,
			DeliverToParty:                    data.DeliverToParty,
			DeliverFromParty:                  data.DeliverFromParty,
			DefaultRelation:                   data.DefaultRelation,
			CreationDate:                      data.CreationDate,
			LastChangeDate:                    data.LastChangeDate,
			IsMarkedForDeletion:               data.IsMarkedForDeletion,
			BuyerName:                         data.BuyerName,
			SellerName:                        data.SellerName,
			DeliverToPartyName:                data.DeliverToPartyName,
			DeliverFromPartyName:              data.DeliverFromPartyName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &deliveryRelation, nil
	}

	return &deliveryRelation, nil
}

func ConvertToBillingRelation(rows *sql.Rows) (*[]BillingRelation, error) {
	defer rows.Close()
	billingRelation := make([]BillingRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BillingRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.DefaultRelation,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.IsExportImport,
			&pm.TransactionTaxCategory,
			&pm.TransactionTaxClassification,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &billingRelation, err
		}

		data := pm
		billingRelation = append(billingRelation, BillingRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			DefaultRelation:                  data.DefaultRelation,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			IsExportImport:                   data.IsExportImport,
			TransactionTaxCategory:           data.TransactionTaxCategory,
			TransactionTaxClassification:     data.TransactionTaxClassification,
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &billingRelation, nil
	}

	return &billingRelation, nil
}

func ConvertToPaymentRelation(rows *sql.Rows) (*[]PaymentRelation, error) {
	defer rows.Close()
	paymentRelation := make([]PaymentRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.Payer,
			&pm.Payee,
			&pm.DefaultRelation,
			&pm.PayerHouseBank,
			&pm.PayerHouseBankAccount,
			&pm.PayeeHouseBank,
			&pm.PayeeHouseBankAccount,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentRelation, err
		}

		data := pm
		paymentRelation = append(paymentRelation, PaymentRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			DefaultRelation:                  data.DefaultRelation,
			PayerHouseBank:                   data.PayerHouseBank,
			PayerHouseBankAccount:            data.PayerHouseBankAccount,
			PayeeHouseBank:                   data.PayeeHouseBank,
			PayeeHouseBankAccount:            data.PayeeHouseBankAccount,
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &paymentRelation, nil
	}

	return &paymentRelation, nil
}

func ConvertToDeliveryPlantRelation(rows *sql.Rows) (*[]DeliveryPlantRelation, error) {
	defer rows.Close()
	deliveryPlantRelation := make([]DeliveryPlantRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := DeliveryPlantRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.DefaultRelation,
			&pm.MRPType,
			&pm.MRPController,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
			&pm.BuyerName,
			&pm.SellerName,
			&pm.DeliverToPartyName,
			&pm.DeliverFromPartyName,
			&pm.DeliverToPlantName,
			&pm.DeliverFromPlantName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &deliveryPlantRelation, err
		}

		data := pm
		deliveryPlantRelation = append(deliveryPlantRelation, DeliveryPlantRelation{
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			DefaultRelation:                        data.DefaultRelation,
			MRPType:                                data.MRPType,
			MRPController:                          data.MRPController,
			CreationDate:                           data.CreationDate,
			LastChangeDate:                         data.LastChangeDate,
			IsMarkedForDeletion:                    data.IsMarkedForDeletion,
			BuyerName:                              data.BuyerName,
			SellerName:                             data.SellerName,
			DeliverToPartyName:                     data.DeliverToPartyName,
			DeliverFromPartyName:                   data.DeliverFromPartyName,
			DeliverToPlantName:                     data.DeliverToPlantName,
			DeliverFromPlantName:                   data.DeliverFromPlantName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &deliveryPlantRelation, nil
	}

	return &deliveryPlantRelation, nil
}

func ConvertToDeliveryPlantRelationProduct(rows *sql.Rows) (*[]DeliveryPlantRelationProduct, error) {
	defer rows.Close()
	deliveryPlantRelationProduct := make([]DeliveryPlantRelationProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryPlantRelationProduct{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.Product,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlantStorageLocation,
			&pm.DeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantityInBaseUnit,
			&pm.PlanningTimeFenceInDays,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.SafetyDurationUnit,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinimumDeliveryQuantityInBaseUnit,
			&pm.MinimumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &deliveryPlantRelationProduct, err
		}

		data := pm
		deliveryPlantRelationProduct = append(deliveryPlantRelationProduct, DeliveryPlantRelationProduct{
			SupplyChainRelationshipID:                 data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:         data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:    data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                     data.Buyer,
			Seller:                                    data.Seller,
			DeliverToParty:                            data.DeliverToParty,
			DeliverFromParty:                          data.DeliverFromParty,
			DeliverToPlant:                            data.DeliverToPlant,
			DeliverFromPlant:                          data.DeliverFromPlant,
			Product:                                   data.Product,
			DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
			DeliveryUnit:                              data.DeliveryUnit,
			QuantityPerPackage:                        data.QuantityPerPackage,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
			MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &deliveryPlantRelationProduct, nil
	}

	return &deliveryPlantRelationProduct, nil
}

func ConvertToDeliveryPlantRelationProductMRPArea(rows *sql.Rows) (*[]DeliveryPlantRelationProductMRPArea, error) {
	defer rows.Close()
	deliveryPlantRelationProductMRPArea := make([]DeliveryPlantRelationProductMRPArea, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryPlantRelationProductMRPArea{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.Product,
			&pm.MRPArea,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlantStorageLocation,
			&pm.DeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantityInBaseUnit,
			&pm.PlanningTimeFenceInDays,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.SafetyDurationUnit,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinimumDeliveryQuantityInBaseUnit,
			&pm.MinimumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &deliveryPlantRelationProductMRPArea, err
		}
		data := pm

		deliveryPlantRelationProductMRPArea = append(deliveryPlantRelationProductMRPArea, DeliveryPlantRelationProductMRPArea{
			SupplyChainRelationshipID:                 data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:         data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:    data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                     data.Buyer,
			Seller:                                    data.Seller,
			DeliverToParty:                            data.DeliverToParty,
			DeliverFromParty:                          data.DeliverFromParty,
			DeliverToPlant:                            data.DeliverToPlant,
			DeliverFromPlant:                          data.DeliverFromPlant,
			Product:                                   data.Product,
			MRPArea:                                   data.MRPArea,
			DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
			DeliveryUnit:                              data.DeliveryUnit,
			QuantityPerPackage:                        data.QuantityPerPackage,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
			MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &deliveryPlantRelationProductMRPArea, nil
	}

	return &deliveryPlantRelationProductMRPArea, nil
}

func ConvertToStockConfPlantRelation(rows *sql.Rows) (*[]StockConfPlantRelation, error) {
	defer rows.Close()
	stockConfPlantRelation := make([]StockConfPlantRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.StockConfPlantRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.DefaultRelation,
			&pm.MRPType,
			&pm.MRPController,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &stockConfPlantRelation, err
		}
		data := pm

		stockConfPlantRelation = append(stockConfPlantRelation, StockConfPlantRelation{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			DefaultRelation:                         data.DefaultRelation,
			MRPType:                                 data.MRPType,
			MRPController:                           data.MRPController,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &stockConfPlantRelation, nil
	}

	return &stockConfPlantRelation, nil
}

func ConvertToStockConfPlantRelationProduct(rows *sql.Rows) (*[]StockConfPlantRelationProduct, error) {
	defer rows.Close()
	stockConfPlantRelationProduct := make([]StockConfPlantRelationProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.StockConfPlantRelationProduct{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.Product,
			&pm.MRPType,
			&pm.MRPController,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &stockConfPlantRelationProduct, err
		}
		data := pm

		stockConfPlantRelationProduct = append(stockConfPlantRelationProduct, StockConfPlantRelationProduct{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			Product:                                 data.Product,
			MRPType:                                 data.MRPType,
			MRPController:                           data.MRPController,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &stockConfPlantRelationProduct, nil
	}

	return &stockConfPlantRelationProduct, nil
}

func ConvertToProductionPlantRelation(rows *sql.Rows) (*[]ProductionPlantRelation, error) {
	defer rows.Close()
	productionPlantRelation := make([]ProductionPlantRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductionPlantRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.DefaultRelation,
			&pm.MRPType,
			&pm.MRPController,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productionPlantRelation, err
		}
		data := pm

		productionPlantRelation = append(productionPlantRelation, ProductionPlantRelation{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			DefaultRelation:                          data.DefaultRelation,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productionPlantRelation, nil
	}

	return &productionPlantRelation, nil
}

func ConvertToProductionPlantRelationProductMRP(rows *sql.Rows) (*[]ProductionPlantRelationProductMRP, error) {
	defer rows.Close()
	productionPlantRelationProductMRP := make([]ProductionPlantRelationProductMRP, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductionPlantRelationProductMRP{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.Product,
			&pm.ProductionPlantStorageLocation,
			&pm.MRPType,
			&pm.MRPController,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productionPlantRelationProductMRP, err
		}
		data := pm

		productionPlantRelationProductMRP = append(productionPlantRelationProductMRP, ProductionPlantRelationProductMRP{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			Product:                                  data.Product,
			ProductionPlantStorageLocation:           data.ProductionPlantStorageLocation,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productionPlantRelationProductMRP, nil
	}

	return &productionPlantRelationProductMRP, nil
}

func ConvertToFreightRelation(rows *sql.Rows) (*[]FreightRelation, error) {
	defer rows.Close()
	freightRelation := make([]FreightRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.FreightRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.FreightPartner,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &freightRelation, err
		}

		data := pm
		freightRelation = append(freightRelation, FreightRelation{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			CreationDate:              data.CreationDate,
			LastChangeDate:            data.LastChangeDate,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &freightRelation, nil
	}

	return &freightRelation, nil
}

func ConvertToFreightTransaction(rows *sql.Rows) (*[]FreightTransaction, error) {
	defer rows.Close()
	freightTransaction := make([]FreightTransaction, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.FreightTransaction{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipFreightID,
			&pm.Buyer,
			&pm.Seller,
			&pm.FreightPartner,
			&pm.TransactionCurrency,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Incoterms,
			&pm.AccountAssignmentGroup,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.QuotationIsBlocked,
			&pm.OrderIsBlocked,
			&pm.DeliveryIsBlocked,
			&pm.BillingIsBlocked,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &freightTransaction, err
		}

		data := pm
		freightTransaction = append(freightTransaction, FreightTransaction{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipFreightID: data.SupplyChainRelationshipFreightID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			FreightPartner:                   data.FreightPartner,
			TransactionCurrency:              data.TransactionCurrency,
			PaymentTerms:                     data.PaymentTerms,
			PaymentMethod:                    data.PaymentMethod,
			Incoterms:                        data.Incoterms,
			AccountAssignmentGroup:           data.AccountAssignmentGroup,
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
			QuotationIsBlocked:               data.QuotationIsBlocked,
			OrderIsBlocked:                   data.OrderIsBlocked,
			DeliveryIsBlocked:                data.DeliveryIsBlocked,
			BillingIsBlocked:                 data.BillingIsBlocked,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &freightTransaction, nil
	}

	return &freightTransaction, nil
}

func ConvertToFreightBillingRelation(rows *sql.Rows) (*[]FreightBillingRelation, error) {
	defer rows.Close()
	freightBillingRelation := make([]FreightBillingRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.FreightBillingRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipFreightID,
			&pm.SupplyChainRelationshipFreightBillingID,
			&pm.Buyer,
			&pm.Seller,
			&pm.FreightPartner,
			&pm.FreightBillToParty,
			&pm.FreightBillFromParty,
			&pm.DefaultRelation,
			&pm.FreightBillToCountry,
			&pm.FreightBillFromCountry,
			&pm.IsExportImport,
			&pm.TransactionTaxCategory,
			&pm.TransactionTaxClassification,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &freightBillingRelation, err
		}

		data := pm
		freightBillingRelation = append(freightBillingRelation, FreightBillingRelation{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipFreightID:        data.SupplyChainRelationshipFreightID,
			SupplyChainRelationshipFreightBillingID: data.SupplyChainRelationshipFreightBillingID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			FreightPartner:                          data.FreightPartner,
			FreightBillToParty:                      data.FreightBillToParty,
			FreightBillFromParty:                    data.FreightBillFromParty,
			DefaultRelation:                         data.DefaultRelation,
			FreightBillToCountry:                    data.FreightBillToCountry,
			FreightBillFromCountry:                  data.FreightBillFromCountry,
			IsExportImport:                          data.IsExportImport,
			TransactionTaxCategory:                  data.TransactionTaxCategory,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &freightBillingRelation, nil
	}

	return &freightBillingRelation, nil
}

func ConvertToFreightPaymentRelation(rows *sql.Rows) (*[]FreightPaymentRelation, error) {
	defer rows.Close()
	freightPaymentRelation := make([]FreightPaymentRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.FreightPaymentRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipFreightID,
			&pm.SupplyChainRelationshipFreightBillingID,
			&pm.SupplyChainRelationshipFreightPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.FreightPartner,
			&pm.FreightBillToParty,
			&pm.FreightBillFromParty,
			&pm.FreightPayer,
			&pm.FreightPayee,
			&pm.DefaultRelation,
			&pm.PayerHouseBank,
			&pm.PayerHouseBankAccount,
			&pm.PayeeHouseBank,
			&pm.PayeeHouseBankAccount,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &freightPaymentRelation, err
		}

		data := pm
		freightPaymentRelation = append(freightPaymentRelation, FreightPaymentRelation{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipFreightID:        data.SupplyChainRelationshipFreightID,
			SupplyChainRelationshipFreightBillingID: data.SupplyChainRelationshipFreightBillingID,
			SupplyChainRelationshipFreightPaymentID: data.SupplyChainRelationshipFreightPaymentID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			FreightPartner:                          data.FreightPartner,
			FreightBillToParty:                      data.FreightBillToParty,
			FreightBillFromParty:                    data.FreightBillFromParty,
			FreightPayer:                            data.FreightPayer,
			FreightPayee:                            data.FreightPayee,
			DefaultRelation:                         data.DefaultRelation,
			PayerHouseBank:                          data.PayerHouseBank,
			PayerHouseBankAccount:                   data.PayerHouseBankAccount,
			PayeeHouseBank:                          data.PayeeHouseBank,
			PayeeHouseBankAccount:                   data.PayeeHouseBankAccount,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &freightPaymentRelation, nil
	}

	return &freightPaymentRelation, nil
}
