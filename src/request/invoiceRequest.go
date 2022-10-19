package request

type InvoiceRequest struct {
	TaxInfo     *TaxInfo     `json:"tax_info"`
	InvoiceInfo *InvoiceInfo `json:"invoice_info"`
	Details     []*Detail    `json:"details"`
	Payments    []*Payment   `json:"payments"`
}

type TaxInfo struct {
	Environment  string `json:"enviroment"`
	BusinessName string `json:"business_name"`
	Ruc          string `json:"ruc"`
	Sequencial   string `json:"sequencial"`
}

type InvoiceInfo struct {
	EmissionDate       string `json:"emission_date"`
	IdentificationType string `json:"identification_type"`
	Identification     string `json:"identification"`
	Address            string `json:"address"`
}

type Detail struct {
	ProductCode string  `json:"product_code"`
	Description string  `json:"description"`
	Amount      string  `json:"amount"`
	UnitPrice   float64 `json:"unit_price"`
	Discount    float64 `json:"discount"`
	Tax         float64 `json:"tax"`
}

type Payment struct {
	PaymentForm string  `json:"payment_form"`
	Total       float64 `json:"total"`
	SalesTime   int     `json:"sales_time"`
	UnitTime    string  `json:"unit_time"`
}
