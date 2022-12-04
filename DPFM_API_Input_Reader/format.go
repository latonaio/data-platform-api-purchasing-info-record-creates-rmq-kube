package dpfm_api_input_reader

import (
	"data-platform-api-purchasing-info-record-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.PurchasingInfoRecord
	return &requests.General{
		BusinessPartner:          data.BusinessPartner,
		PurchasingInfoRecord:     data.PurchasingInfoRecord,
		Supplier:                 data.Supplier,
		Product:                  data.Product,
		ProductGroup:             data.ProductGroup,
		OrderQuantityUnit:        data.OrderQuantityUnit,
		SupplierProductNumber:    data.SupplierProductNumber,
		ValidityStartDate:        data.ValidityStartDate,
		ValidityEndDate:          data.ValidityEndDate,
		CreationDate:             data.CreationDate,
		PurchasingInfoRecordDesc: data.PurchasingInfoRecordDesc,
		LastChangeDateTime:       data.LastChangeDateTime,
		IsDeleted:                data.IsDeleted,
	}
}

func (sdc *SDC) ConvertToPlant() *requests.Plant {
	dataPurchasingInfoRecord := sdc.PurchasingInfoRecord
	data := sdc.PurchasingInfoRecord.Plant
	return &requests.Plant{
		BusinessPartner:               dataPurchasingInfoRecord.BusinessPartner,
		PurchasingInfoRecord:          dataPurchasingInfoRecord.PurchasingInfoRecord,
		Plant:                         data.Plant,
		Supplier:                      data.Supplier,
		Product:                       data.Product,
		ProductGroup:                  data.ProductGroup,
		OrderQuantityUnit:             data.OrderQuantityUnit,
		MinimumPurchaseOrderQuantity:  data.MinimumPurchaseOrderQuantity,
		StandardPurchaseOrderQuantity: data.StandardPurchaseOrderQuantity,
		ProductPlannedDeliveryDurn:    data.ProductPlannedDeliveryDurn,
		InvoiceIsGoodsReceiptBased:    data.InvoiceIsGoodsReceiptBased,
		TaxCode:                       data.TaxCode,
		Incoterms:                     data.Incoterms,
		MaximumOrderQuantity:          data.MaximumOrderQuantity,
		IsRelevantForAutomSrcg:        data.IsRelevantForAutomSrcg,
		IsOrderAcknRqd:                data.IsOrderAcknRqd,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}
