# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChargeInvoice**](PaymentsApi.md#ChargeInvoice) | **Post** /invoice/{invoiceId}/charge | Charge Invoice
[**DownloadInvoice**](PaymentsApi.md#DownloadInvoice) | **Get** /invoice/{invoiceId}/download | Download Invoice
[**DownloadReceipt**](PaymentsApi.md#DownloadReceipt) | **Get** /receipt/{receiptId}/download | Download Receipt
[**GetInvoicesForCustomer**](PaymentsApi.md#GetInvoicesForCustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices
[**GetPaymentsForCustomer**](PaymentsApi.md#GetPaymentsForCustomer) | **Get** /customer/{customerId}/payment | List Customer Payments
[**ListPayment**](PaymentsApi.md#ListPayment) | **Get** /payment | List
[**RefundPayment**](PaymentsApi.md#RefundPayment) | **Post** /payment/{paymentId}/refund | Refund Payment

# **ChargeInvoice**
> InlineResponse20014 ChargeInvoice(ctx, invoiceId)
Charge Invoice

Attempts to charge a card that is on file for the invoice amount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The id of the invoice | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadInvoice**
> *os.File DownloadInvoice(ctx, invoiceId)
Download Invoice

Returns the pdf blob for the invoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The id of the invoice | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadReceipt**
> *os.File DownloadReceipt(ctx, receipt)
Download Receipt

Returns the pdf blob for the Receipt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **receipt** | **string**| The id of the receipt | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoicesForCustomer**
> InlineResponse2006 GetInvoicesForCustomer(ctx, customerId)
List Customer Invoices

List Customer Invoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentsForCustomer**
> InlineResponse2004 GetPaymentsForCustomer(ctx, customerId, optional)
List Customer Payments

List Customer Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
 **optional** | ***PaymentsApiGetPaymentsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiGetPaymentsForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **name** | **optional.String**| The name to search for | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayment**
> InlineResponse2009 ListPayment(ctx, optional)
List

List all payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiListPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiListPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **name** | **optional.String**| The name to search for | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundPayment**
> RefundPayment(ctx, body, paymentId)
Refund Payment

Issue a refund for payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueRefundPayment**](IssueRefundPayment.md)|  | 
  **paymentId** | **string**| The id of the payment | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

