/*
 * BillaBear
 *
 * The REST API provided by BillaBear
 *
 * API version: 1.0.0
 * Contact: support@billabear.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Price struct {
	Id string `json:"id,omitempty"`
	Amount int32 `json:"amount"`
	// Three-letter ISO currency code. Must be upper-case
	Currency string `json:"currency"`
	ExternalReference string `json:"external_reference,omitempty"`
	Recurring bool `json:"recurring"`
	// Required if recurring is true
	Schedule string `json:"schedule,omitempty"`
	// If the price is including tax. If false tax will be added on top of the price.
	IncludingTax bool `json:"including_tax,omitempty"`
	Public bool `json:"public,omitempty"`
}