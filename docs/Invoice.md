# Invoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Number** | **string** |  | [optional] [default to null]
**Currency** | **string** |  | [optional] [default to null]
**Customer** | [***Customer**](Customer.md) |  | [optional] [default to null]
**TaxTotal** | **int32** |  | [optional] [default to null]
**SubTotal** | **int32** |  | [optional] [default to null]
**AmountDue** | **int32** |  | [optional] [default to null]
**Paid** | **bool** |  | [optional] [default to null]
**PayLink** | **string** |  | [optional] [default to null]
**CreatedAt** | **string** |  | [optional] [default to null]
**PaidAt** | **string** |  | [optional] [default to null]
**DueDate** | **string** |  | [optional] [default to null]
**BillerAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Payeeaddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Lines** | [**[]InvoiceLines**](Invoice_lines.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

