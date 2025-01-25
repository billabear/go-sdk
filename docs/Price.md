# Price

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Amount** | **int32** |  | [default to null]
**Currency** | **string** | Three-letter ISO currency code. Must be upper-case | [default to null]
**ExternalReference** | **string** |  | [optional] [default to null]
**Recurring** | **bool** |  | [default to null]
**Schedule** | **string** | Required if recurring is true | [optional] [default to null]
**IncludingTax** | **bool** | If the price is including tax. If false tax will be added on top of the price. | [optional] [default to null]
**Public** | **bool** |  | [optional] [default to null]
**Metric** | [***Metric**](Metric.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

