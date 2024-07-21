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

type CheckoutSubscriptions struct {
	Plan string `json:"plan,omitempty"`
	Price string `json:"price,omitempty"`
	SeatNumber int32 `json:"seat_number,omitempty"`
}