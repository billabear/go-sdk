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

type InvoiceLines struct {
	Description string `json:"description,omitempty"`
	Currency string `json:"currency,omitempty"`
	Total int32 `json:"total,omitempty"`
	SubTotal int32 `json:"sub_total,omitempty"`
	TaxTotal int32 `json:"tax_total,omitempty"`
	TaxRate float64 `json:"tax_rate,omitempty"`
}
