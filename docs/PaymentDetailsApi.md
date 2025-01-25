# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteFrontendPaymentDetails**](PaymentDetailsApi.md#CompleteFrontendPaymentDetails) | **Post** /customer/{customerId}/payment-methods/frontend-payment-token | Complete Frontend Detail Collection
[**DeletePaymentDetails**](PaymentDetailsApi.md#DeletePaymentDetails) | **Delete** /payment-methods/{paymentDetailsId} | Delete
[**DeletePaymentDetailsCustomer**](PaymentDetailsApi.md#DeletePaymentDetailsCustomer) | **Delete** /customer/{customerId}/payment-methods/{paymentDetailsId} | Delete With Customer
[**GetPaymentDetails**](PaymentDetailsApi.md#GetPaymentDetails) | **Get** /payment-methods/{paymentDetailsId} | Fetch
[**ListPaymentDetails**](PaymentDetailsApi.md#ListPaymentDetails) | **Get** /customer/{customerId}/payment-methods | List Customer&#x27;s Payment Details
[**MakeDefaultPaymentDetails**](PaymentDetailsApi.md#MakeDefaultPaymentDetails) | **Post** /payment-methods/{paymentDetailsId}/default | Make Default
[**MakeDefaultPaymentDetailsCustomer**](PaymentDetailsApi.md#MakeDefaultPaymentDetailsCustomer) | **Post** /customer/{customerId}/payment-methods/{paymentDetailsId}/default | Make Default With Customer
[**StartFrontendPaymentDetails**](PaymentDetailsApi.md#StartFrontendPaymentDetails) | **Get** /customer/{customerId}/payment-methods/frontend-payment-token | Start Frontend Detail Collection

# **CompleteFrontendPaymentDetails**
> PaymentDetails CompleteFrontendPaymentDetails(ctx, body, customerId)
Complete Frontend Detail Collection

Complete frontend payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FrontendToken**](FrontendToken.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**PaymentDetails**](PaymentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentDetails**
> DeletePaymentDetails(ctx, paymentDetailsId)
Delete

Delete Payment Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentDetailsId** | **string**| The id of the payment details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentDetailsCustomer**
> DeletePaymentDetailsCustomer(ctx, customerId, paymentDetailsId)
Delete With Customer

Delete Payment Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
  **paymentDetailsId** | **string**| The id of the payment details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentDetails**
> PaymentDetails GetPaymentDetails(ctx, paymentDetailsId)
Fetch

Fetch the payment cards

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentDetailsId** | **string**| The id of the payment details | 

### Return type

[**PaymentDetails**](PaymentDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentDetails**
> InlineResponse2007 ListPaymentDetails(ctx, customerId)
List Customer's Payment Details

List all customers <br><br>Added in version 1.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeDefaultPaymentDetails**
> MakeDefaultPaymentDetails(ctx, customerId, paymentDetailsId)
Make Default

Delete Payment Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
  **paymentDetailsId** | **string**| The id of the payment details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeDefaultPaymentDetailsCustomer**
> MakeDefaultPaymentDetailsCustomer(ctx, customerId, paymentDetailsId)
Make Default With Customer

Delete Payment Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
  **paymentDetailsId** | **string**| The id of the payment details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartFrontendPaymentDetails**
> FrontendToken StartFrontendPaymentDetails(ctx, customerId)
Start Frontend Detail Collection

Start frontend payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**FrontendToken**](FrontendToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

