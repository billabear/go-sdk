# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChargeInvoice**](InvoicesApi.md#ChargeInvoice) | **Post** /invoice/{invoiceId}/charge | Charge Invoice
[**DownloadInvoice**](InvoicesApi.md#DownloadInvoice) | **Get** /invoice/{invoiceId}/download | Download Invoice
[**GetInvoicesForCustomer**](InvoicesApi.md#GetInvoicesForCustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices

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

