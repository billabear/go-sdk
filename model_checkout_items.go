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

type CheckoutItems struct {
	Description string `json:"description,omitempty"`
	Amount int32 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	IncludeTax bool `json:"include_tax,omitempty"`
	TaxType string `json:"tax_type,omitempty"`
}
