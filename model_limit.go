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

type Limit struct {
	Feature *Feature `json:"feature,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
