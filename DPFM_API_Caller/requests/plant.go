package requests

type Plant struct {
	BusinessPartner               *int     `json:"BusinessPartner"`
	PurchasingInfoRecord          *int     `json:"PurchasingInfoRecord"`
	Plant                         string   `json:"Plant"`
	Supplier                      *int     `json:"Supplier"`
	Product                       string   `json:"Product"`
	ProductGroup                  string   `json:"ProductGroup"`
	OrderQuantityUnit             string   `json:"OrderQuantityUnit"`
	MinimumPurchaseOrderQuantity  *float32 `json:"MinimumPurchaseOrderQuantity"`
	StandardPurchaseOrderQuantity *float32 `json:"StandardPurchaseOrderQuantity"`
	ProductPlannedDeliveryDurn    *int     `json:"ProductPlannedDeliveryDurn"`
	InvoiceIsGoodsReceiptBased    *bool    `json:"InvoiceIsGoodsReceiptBased"`
	TaxCode                       string   `json:"TaxCode"`
	Incoterms                     string   `json:"Incoterms"`
	MaximumOrderQuantity          *float32 `json:"MaximumOrderQuantity"`
	IsRelevantForAutomSrcg        *bool    `json:"IsRelevantForAutomSrcg"`
	IsOrderAcknRqd                *bool    `json:"IsOrderAcknRqd"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
