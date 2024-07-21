# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [default to null]
**TaxNumber** | **string** | The tax number for the customer &lt;strong&gt;Since 1.1&lt;/strong&gt; | [optional] [default to null]
**StandardTaxRate** | **float32** | The tax rate for the customer on for standard services a &lt;strong&gt;Since 1.1&lt;/strong&gt; | [optional] [default to null]
**DigitalTaxRate** | **float32** | The tax rate for the customer on digital services &lt;strong&gt;Since 1.1&lt;/strong&gt; | [optional] [default to null]
**BillingType** | **string** | Choice between card and invoice | [optional] [default to null]
**Type_** | **string** | Choice between &#x27;individual&#x27; and &#x27;business&#x27;. When not provided &#x27;individual&#x27; is used. &lt;strong&gt;Since 1.1&lt;/strong&gt; | [optional] [default to null]
**Reference** | **string** |  | [optional] [default to null]
**ExternalReference** | **string** |  | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**Locale** | **string** | Defaults to &#x27;en&#x27; if not sent. | [optional] [default to null]
**Brand** | **string** | Defaults to &#x27;default&#x27; if not sent. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

