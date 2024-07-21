# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadReceipt**](ReceiptApi.md#DownloadReceipt) | **Get** /receipt/{receiptId}/download | Download Receipt

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

