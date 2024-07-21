# BillaBear{{classname}}

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSeatsSubscriptions**](CustomersApi.md#AddSeatsSubscriptions) | **Post** /subscription/{subscriptionId}/seats/add | Add Seats
[**ApplyVoucherToCustomer**](CustomersApi.md#ApplyVoucherToCustomer) | **Post** /customer/{customerId}/voucher | Apply voucher
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /customer | Create
[**DisableCustomer**](CustomersApi.md#DisableCustomer) | **Post** /customer/{customerId}/disable | Disable Customer
[**EnableCustomer**](CustomersApi.md#EnableCustomer) | **Post** /customer/{customerId}/enable | Enable Customer
[**GetActiveForCustomer**](CustomersApi.md#GetActiveForCustomer) | **Get** /customer/{customerId}/subscription/active | List Customer Active Subscriptions
[**GetAllCustomers**](CustomersApi.md#GetAllCustomers) | **Get** /customer | List
[**GetCustomerById**](CustomersApi.md#GetCustomerById) | **Get** /customer/{customerId} | Detail
[**GetCustomerLimitsById**](CustomersApi.md#GetCustomerLimitsById) | **Get** /customer/{customerId}/limits | Fetch Customer Limits
[**GetForCustomer**](CustomersApi.md#GetForCustomer) | **Get** /customer/{customerId}/subscription | List Customer Subscriptions
[**GetInvoicesForCustomer**](CustomersApi.md#GetInvoicesForCustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices
[**GetPaymentsForCustomer**](CustomersApi.md#GetPaymentsForCustomer) | **Get** /customer/{customerId}/payment | List Customer Payments
[**GetRefundsForCustomer**](CustomersApi.md#GetRefundsForCustomer) | **Get** /customer/{customerId}/refund | List Customer Refunds
[**ListPaymentDetails**](CustomersApi.md#ListPaymentDetails) | **Get** /customer/{customerId}/payment-methods | List Customer&#x27;s Payment Details
[**RemoveSeatsSubscriptions**](CustomersApi.md#RemoveSeatsSubscriptions) | **Post** /subscription/{subscriptionId}/seats/remove | Remove Seats
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Put** /customer/{customerId} | Update

# **AddSeatsSubscriptions**
> InlineResponse20011 AddSeatsSubscriptions(ctx, body, subscriptionId)
Add Seats

Adds seats to a per seat subscription<br><br><strong>Since 1.1.4</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SeatsAddBody**](SeatsAddBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyVoucherToCustomer**
> string ApplyVoucherToCustomer(ctx, body, customerId)
Apply voucher

Apply Voucher to Customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VoucherCode**](VoucherCode.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomer**
> Customer CreateCustomer(ctx, body)
Create

Create a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Customer**](Customer.md)|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableCustomer**
> string DisableCustomer(ctx, customerId)
Disable Customer

Disable customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableCustomer**
> string EnableCustomer(ctx, customerId)
Enable Customer

Enable a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveForCustomer**
> InlineResponse2006 GetActiveForCustomer(ctx, customerId)
List Customer Active Subscriptions

List all Active customer subscriptions

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

# **GetAllCustomers**
> InlineResponse200 GetAllCustomers(ctx, optional)
List

List all customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersApiGetAllCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiGetAllCustomersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **email** | **optional.String**| The email to search for | 
 **country** | **optional.String**| The country code to search for | 
 **reference** | **optional.String**| The reference to search for | 
 **externalReference** | **optional.String**| The external reference to search for | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerById**
> Customer GetCustomerById(ctx, customerId)
Detail

Info for a specific customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**Customer**](Customer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerLimitsById**
> InlineResponse2001 GetCustomerLimitsById(ctx, customerId)
Fetch Customer Limits

Limits for a specific customer

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

# **GetForCustomer**
> InlineResponse2006 GetForCustomer(ctx, customerId)
List Customer Subscriptions

List all customer subscriptions<br><br><strong>Since 1.1</strong>

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

# **GetInvoicesForCustomer**
> InlineResponse2004 GetInvoicesForCustomer(ctx, customerId)
List Customer Invoices

List Customer Invoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentsForCustomer**
> InlineResponse2003 GetPaymentsForCustomer(ctx, customerId, optional)
List Customer Payments

List Customer Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
 **optional** | ***CustomersApiGetPaymentsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiGetPaymentsForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| How many items to return at one time (max 100) | 
 **lastKey** | **optional.String**| The key to be used in pagination to say what the last key of the previous page was | 
 **name** | **optional.String**| The name to search for | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRefundsForCustomer**
> InlineResponse2002 GetRefundsForCustomer(ctx, customerId, optional)
List Customer Refunds

List Customer Refund

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The id of the customer to retrieve | 
 **optional** | ***CustomersApiGetRefundsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiGetRefundsForCustomerOpts struct
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

# **ListPaymentDetails**
> InlineResponse2005 ListPaymentDetails(ctx, customerId)
List Customer's Payment Details

List all customers <br><br>Added in version 1.1

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

# **RemoveSeatsSubscriptions**
> InlineResponse20011 RemoveSeatsSubscriptions(ctx, body, subscriptionId)
Remove Seats

Remove seats to a per seat subscription<br><br><strong>Since 1.1.4</strong>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SeatsRemoveBody**](SeatsRemoveBody.md)|  | 
  **subscriptionId** | **string**| The id of the subscription to retrieve | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> Customer UpdateCustomer(ctx, body, customerId)
Update

Update a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Customer**](Customer.md)|  | 
  **customerId** | **string**| The id of the customer to retrieve | 

### Return type

[**Customer**](Customer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

