# {{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSeatsSubscriptions**](SubscriptionsApi.md#AddSeatsSubscriptions) | **Post** /subscription/{subscriptionId}/seats/add | Add Seats
[**CancelSubscription**](SubscriptionsApi.md#CancelSubscription) | **Post** /subscription/{subscriptionId}/cancel | Cancel Subscription
[**ChangeSubscriptionPrice**](SubscriptionsApi.md#ChangeSubscriptionPrice) | **Post** /subscription/{subscriptionId}/price | Change Price
[**CreateSubscription**](SubscriptionsApi.md#CreateSubscription) | **Post** /customer/{customerId}/subscription/start | Create Subscription
[**CustomerChangeSubscriptionPlan**](SubscriptionsApi.md#CustomerChangeSubscriptionPlan) | **Post** /subscription/{subscriptionId}/plan | Change Subscription Plan
[**ExtendTrial**](SubscriptionsApi.md#ExtendTrial) | **Post** /subscription/{subscriptionId}/extend | Extend Trial Subscription
[**GetActiveForCustomer**](SubscriptionsApi.md#GetActiveForCustomer) | **Get** /customer/{customerId}/subscription/active | List Customer Active Subscriptions
[**GetForCustomer**](SubscriptionsApi.md#GetForCustomer) | **Get** /customer/{customerId}/subscription | List Customer Subscriptions
[**ListSubscriptionPlans**](SubscriptionsApi.md#ListSubscriptionPlans) | **Get** /subscription/plans | List Subscription Plans
[**ListSubscriptions**](SubscriptionsApi.md#ListSubscriptions) | **Get** /subscription | List
[**RemoveSeatsSubscriptions**](SubscriptionsApi.md#RemoveSeatsSubscriptions) | **Post** /subscription/{subscriptionId}/seats/remove | Remove Seats
[**ShowSubscriptionById**](SubscriptionsApi.md#ShowSubscriptionById) | **Get** /subscription/{subscriptionId} | Detail
[**StartTrial**](SubscriptionsApi.md#StartTrial) | **Post** /customer/{customerId}/subscription/trial | Start Trial Subscription For Customer

# **AddSeatsSubscriptions**
> InlineResponse20013 AddSeatsSubscriptions(ctx, body, subscriptionId)
Add Seats

Adds seats to a per seat subscription<br><br><strong>Since 1.1.4</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SeatsAddBody**](SeatsAddBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSubscription**
> CancelSubscription(ctx, body, subscriptionId)
Cancel Subscription

Info for a specific subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionIdCancelBody**](SubscriptionIdCancelBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeSubscriptionPrice**
> InlineResponse20013 ChangeSubscriptionPrice(ctx, body, subscriptionId)
Change Price

Changes the price being used for a price. Useful for changing pricing schedule or just price.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionIdPriceBody**](SubscriptionIdPriceBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscription**
> Subscription CreateSubscription(ctx, body, customerId)
Create Subscription

Create subscription for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionStartBody**](SubscriptionStartBody.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerChangeSubscriptionPlan**
> Subscription CustomerChangeSubscriptionPlan(ctx, body, subscriptionId)
Change Subscription Plan

Change the subscription plan for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionIdPlanBody**](SubscriptionIdPlanBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendTrial**
> Subscription ExtendTrial(ctx, body, subscriptionId)
Extend Trial Subscription

Extend a trial subscription so it's converted from a trial to a normal subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionIdExtendBody**](SubscriptionIdExtendBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveForCustomer**
> InlineResponse2008 GetActiveForCustomer(ctx, customerId)
List Customer Active Subscriptions

List all Active customer subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForCustomer**
> InlineResponse2008 GetForCustomer(ctx, customerId)
List Customer Subscriptions

List all customer subscriptions<br><br><strong>Since 1.1</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubscriptionPlans**
> InlineResponse20012 ListSubscriptionPlans(ctx, optional)
List Subscription Plans

List all subscriptions plans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsApiListSubscriptionPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiListSubscriptionPlansOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubscriptions**
> InlineResponse2008 ListSubscriptions(ctx, optional)
List

List all subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsApiListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiListSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSeatsSubscriptions**
> InlineResponse20013 RemoveSeatsSubscriptions(ctx, body, subscriptionId)
Remove Seats

Remove seats to a per seat subscription<br><br><strong>Since 1.1.4</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SeatsRemoveBody**](SeatsRemoveBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSubscriptionById**
> Subscription ShowSubscriptionById(ctx, subscriptionId)
Detail

Info for a specific subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartTrial**
> Subscription StartTrial(ctx, body, customerId)
Start Trial Subscription For Customer

Start subscription for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionTrialBody**](SubscriptionTrialBody.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

