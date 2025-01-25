/*
 * BillaBear
 *
 * The REST API provided by BillaBear
 *
 * API version: 1.1.0
 * Contact: support@billabear.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package billabear

type Invoice struct {
	Id string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	Currency string `json:"currency,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	TaxTotal int32 `json:"tax_total,omitempty"`
	SubTotal int32 `json:"sub_total,omitempty"`
	AmountDue int32 `json:"amount_due,omitempty"`
	Paid bool `json:"paid,omitempty"`
	PayLink string `json:"pay_link,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	PaidAt string `json:"paid_at,omitempty"`
	DueDate string `json:"due_date,omitempty"`
	BillerAddress *Address `json:"biller_address,omitempty"`
	Payeeaddress *Address `json:"payeeaddress,omitempty"`
	Lines []InvoiceLines `json:"lines,omitempty"`
}
