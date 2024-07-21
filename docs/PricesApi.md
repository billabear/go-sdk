# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrice**](PricesApi.md#CreatePrice) | **Post** /product/{productId}/price | Create
[**ListPrice**](PricesApi.md#ListPrice) | **Get** /product/{productId}/price | List

# **CreatePrice**
> string CreatePrice(ctx, body, productId)
Create

Create a price

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Price**](Price.md)|  | 
  **productId** | **string**| The id of the product to retrieve | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPrice**
> InlineResponse2009 ListPrice(ctx, productId, optional)
List

List all prices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**| The id of the product to retrieve | 
 **optional** | ***PricesApiListPriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricesApiListPriceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

