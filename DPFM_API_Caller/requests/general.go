package requests

type General struct {
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
}
