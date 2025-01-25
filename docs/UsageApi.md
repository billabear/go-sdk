# {{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerUsageLimit**](UsageApi.md#CreateCustomerUsageLimit) | **Post** /customer/{customerId}/uasge-limit | Create Usage Limit
[**CreateEvent**](UsageApi.md#CreateEvent) | **Post** /events | Create Event
[**CustomerCustomerIdUasgeLimitLimitIdDelete**](UsageApi.md#CustomerCustomerIdUasgeLimitLimitIdDelete) | **Delete** /customer/{customerId}/uasge-limit/{limitId} | Delete Usage Limit
[**GetCustomerCosts**](UsageApi.md#GetCustomerCosts) | **Get** /customer/{customerId}/costs | Usage Cost Estimate
[**GetCustomerUsageLimitsById**](UsageApi.md#GetCustomerUsageLimitsById) | **Get** /customer/{customerId}/uasge-limit | Fetch Customer Usage Limits

# **CreateCustomerUsageLimit**
> UsageLimit CreateCustomerUsageLimit(ctx, body, customerId)
Create Usage Limit

Create Usage Limit for the custoemr

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerIdUasgelimitBody**](CustomerIdUasgelimitBody.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**UsageLimit**](UsageLimit.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEvent**
> CreateEvent(ctx, body)
Create Event

Creates an event that is used for usage billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Event**](Event.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerCustomerIdUasgeLimitLimitIdDelete**
> CustomerCustomerIdUasgeLimitLimitIdDelete(ctx, customerId, usageLimitId)
Delete Usage Limit

Delete Usage Limit for the custoemr

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
  **usageLimitId** | **string**| The id of the usage limit | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerCosts**
> InlineResponse2001 GetCustomerCosts(ctx, customerId)
Usage Cost Estimate

The estimated costs from usage based billing for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerUsageLimitsById**
> InlineResponse2005 GetCustomerUsageLimitsById(ctx, customerId)
Fetch Customer Usage Limits

Usage Limits for a specific customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

