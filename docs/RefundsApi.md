# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRefundsForCustomer**](RefundsApi.md#GetRefundsForCustomer) | **Get** /customer/{customerId}/refund | List Customer Refunds
[**ListRefund**](RefundsApi.md#ListRefund) | **Get** /refund | List
[**ShowRefundById**](RefundsApi.md#ShowRefundById) | **Get** /refund/{refundId} | Detail

# **GetRefundsForCustomer**
> InlineResponse2002 GetRefundsForCustomer(ctx, customerId, optional)
List Customer Refunds

List Customer Refund

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
 **optional** | ***RefundsApiGetRefundsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundsApiGetRefundsForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **name** | **optional.String**| The name to search for | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRefund**
> InlineResponse2002 ListRefund(ctx, optional)
List

List all refund

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundsApiListRefundOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundsApiListRefundOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **name** | **optional.String**| The name to search for | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowRefundById**
> Refund ShowRefundById(ctx, refundId)
Detail

Info for a specific Refund

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refundId** | **string**| The id of the refund | 

### Return type

[**Refund**](Refund.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

