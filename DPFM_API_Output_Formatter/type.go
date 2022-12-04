package dpfm_api_output_formatter

type PurchasingInfoRecord struct {
	BusinessPartner          int    `json:"BusinessPartner"`
	PurchasingInfoRecord     int    `json:"PurchasingInfoRecord"`
	Supplier                 int    `json:"Supplier"`
	Product                  string `json:"Product"`
	ProductGroup             string `json:"ProductGroup"`
	OrderQuantityUnit        string `json:"OrderQuantityUnit"`
	SupplierProductNumber    string `json:"SupplierProductNumber"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	CreationDate             string `json:"CreationDate"`
	PurchasingInfoRecordDesc string `json:"PurchasingInfoRecordDesc"`
	LastChangeDateTime       string `json:"LastChangeDateTime"`
	IsDeleted                bool   `json:"IsDeleted"`
	Plant                    Plant  `json:"Plant"`
}

type Plant struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	PurchasingInfoRecord          int     `json:"PurchasingInfoRecord"`
	Plant                         string  `json:"Plant"`
	Supplier                      int     `json:"Supplier"`
	Product                       string  `json:"Product"`
	ProductGroup                  string  `json:"ProductGroup"`
	OrderQuantityUnit             string  `json:"OrderQuantityUnit"`
	MinimumPurchaseOrderQuantity  float32 `json:"MinimumPurchaseOrderQuantity"`
	StandardPurchaseOrderQuantity float32 `json:"StandardPurchaseOrderQuantity"`
	ProductPlannedDeliveryDurn    int     `json:"ProductPlannedDeliveryDurn"`
	InvoiceIsGoodsReceiptBased    bool    `json:"InvoiceIsGoodsReceiptBased"`
	TaxCode                       string  `json:"TaxCode"`
	Incoterms                     string  `json:"Incoterms"`
	MaximumOrderQuantity          float32 `json:"MaximumOrderQuantity"`
	IsRelevantForAutomSrcg        bool    `json:"IsRelevantForAutomSrcg"`
	IsOrderAcknRqd                bool    `json:"IsOrderAcknRqd"`
	IsMarkedForDeletion           bool    `json:"IsMarkedForDeletion"`
}
