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

type SubscriptionIdPlanBody struct {
	When string `json:"when"`
	// The ID for the subscription plan to be used
	Plan string `json:"plan,omitempty"`
	// The ID for the price to be used
	Price string `json:"price"`
}
