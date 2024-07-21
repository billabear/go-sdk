# SubscriptionStartBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionPlan** | **string** | The ID for the subscription plan to be used (Can also be the code name) | [default to null]
**PaymentDetails** | **string** | The Id for the customer&#x27;s payment details to be used | [optional] [default to null]
**CardToken** | **string** | A stripe card token that&#x27;s been created using Stripe&#x27;s js sdk. It&#x27;ll create the payment details for the customer. | [optional] [default to null]
**Price** | **string** | The ID for the price to be used | [optional] [default to null]
**Schedule** | **string** | The schedule of the plan that is to be started. Only used if price isn&#x27;t given. Requires currency as well. | [optional] [default to null]
**Currency** | **string** | The currency of the plan that is to be started. Only used if price isn&#x27;t given. Requires schedule as well. | [optional] [default to null]
**SeatNumbrers** | **int32** |  | [optional] [default to null]
**DenyTrial** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

