package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey        string               `json:"connection_key"`
	Result               bool                 `json:"result"`
	RedisKey             string               `json:"redis_key"`
	Filepath             string               `json:"filepath"`
	APIStatusCode        int                  `json:"api_status_code"`
	RuntimeSessionID     string               `json:"runtime_session_id"`
	BusinessPartnerID    *int                 `json:"business_partner"`
	ServiceLabel         string               `json:"service_label"`
	PurchasingInfoRecord PurchasingInfoRecord `json:"PurchasingInfoRecord"`
	APISchema            string               `json:"api_schema"`
	Accepter             []string             `json:"accepter"`
	OrderID              *int                 `json:"order_id"`
	Deleted              bool                 `json:"deleted"`
	SQLUpdateResult      *bool                `json:"sql_update_result"`
	SQLUpdateError       string               `json:"sql_update_error"`
	SubfuncResult        *bool                `json:"subfunc_result"`
	SubfuncError         string               `json:"subfunc_error"`
	ExconfResult         *bool                `json:"exconf_result"`
	ExconfError          string               `json:"exconf_error"`
	APIProcessingResult  *bool                `json:"api_processing_result"`
	APIProcessingError   string               `json:"api_processing_error"`
}

type PurchasingInfoRecord struct {
	BusinessPartner          *int    `json:"BusinessPartner"`
	PurchasingInfoRecord     *int    `json:"PurchasingInfoRecord"`
	Supplier                 *int    `json:"Supplier"`
	Product                  string  `json:"Product"`
	ProductGroup             string  `json:"ProductGroup"`
	OrderQuantityUnit        string  `json:"OrderQuantityUnit"`
	SupplierProductNumber    string  `json:"SupplierProductNumber"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	ValidityEndDate          *string `json:"ValidityEndDate"`
	CreationDate             *string `json:"CreationDate"`
	PurchasingInfoRecordDesc string  `json:"PurchasingInfoRecordDesc"`
	LastChangeDateTime       *string `json:"LastChangeDateTime"`
	IsDeleted                *bool   `json:"IsDeleted"`
	Plant                    Plant   `json:"Plant"`
}

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
