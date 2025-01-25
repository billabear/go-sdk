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

type Refund struct {
	Id string `json:"id,omitempty"`
	Amount int32 `json:"amount,omitempty"`
	// Three-letter ISO currency code. Must be upper-case
	Currency string `json:"currency,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Comment string `json:"comment,omitempty"`
	Status string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Payment *Paths1paymentgetresponses200contentapplication1jsonschemapropertiesdataitems `json:"payment,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	BillingAdmin *BillingAdmin `json:"billing_admin,omitempty"`
}
