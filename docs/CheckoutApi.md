# {{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckout**](CheckoutApi.md#CreateCheckout) | **Post** /checkout | Create Checkout

# **CreateCheckout**
> InlineResponse201 CreateCheckout(ctx, body)
Create Checkout

Create checkout<br><br><strong>Since 2024.01</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckoutBody**](CheckoutBody.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

