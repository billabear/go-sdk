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

type Subscription struct {
	Id string `json:"id,omitempty"`
	SeatNumber int32 `json:"seat_number,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	Status string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	EndedAt string `json:"ended_at,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	MainExternalReference string `json:"main_external_reference,omitempty"`
	ChildExternalReference string `json:"child_external_reference,omitempty"`
	Price *Price `json:"price,omitempty"`
	Plan *SubscriptionPlan `json:"plan,omitempty"`
}
