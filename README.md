<p align="center">
  <img width="450px" src="https://ha-static-data.s3.eu-central-1.amazonaws.com/github-readme-logo-v2.png">
</p>

<p align="center">
  <h1 style="text-align: center">BillaBear GO SDK</h1>
</p>

## Installation

```shell
go get github.com/billabear/go-sdk
```

## Documentation for API Endpoints

All URIs are relative to *https://{customerId}.billabear.cloud/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CheckoutApi* | [**CreateCheckout**](docs/CheckoutApi.md#createcheckout) | **Post** /checkout | Create Checkout
*CustomersApi* | [**AddSeatsSubscriptions**](docs/CustomersApi.md#addseatssubscriptions) | **Post** /subscription/{subscriptionId}/seats/add | Add Seats
*CustomersApi* | [**ApplyVoucherToCustomer**](docs/CustomersApi.md#applyvouchertocustomer) | **Post** /customer/{customerId}/voucher | Apply voucher
*CustomersApi* | [**CreateCustomer**](docs/CustomersApi.md#createcustomer) | **Post** /customer | Create
*CustomersApi* | [**DisableCustomer**](docs/CustomersApi.md#disablecustomer) | **Post** /customer/{customerId}/disable | Disable Customer
*CustomersApi* | [**EnableCustomer**](docs/CustomersApi.md#enablecustomer) | **Post** /customer/{customerId}/enable | Enable Customer
*CustomersApi* | [**GetActiveForCustomer**](docs/CustomersApi.md#getactiveforcustomer) | **Get** /customer/{customerId}/subscription/active | List Customer Active Subscriptions
*CustomersApi* | [**GetAllCustomers**](docs/CustomersApi.md#getallcustomers) | **Get** /customer | List
*CustomersApi* | [**GetCustomerById**](docs/CustomersApi.md#getcustomerbyid) | **Get** /customer/{customerId} | Detail
*CustomersApi* | [**GetCustomerLimitsById**](docs/CustomersApi.md#getcustomerlimitsbyid) | **Get** /customer/{customerId}/limits | Fetch Customer Limits
*CustomersApi* | [**GetForCustomer**](docs/CustomersApi.md#getforcustomer) | **Get** /customer/{customerId}/subscription | List Customer Subscriptions
*CustomersApi* | [**GetInvoicesForCustomer**](docs/CustomersApi.md#getinvoicesforcustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices
*CustomersApi* | [**GetPaymentsForCustomer**](docs/CustomersApi.md#getpaymentsforcustomer) | **Get** /customer/{customerId}/payment | List Customer Payments
*CustomersApi* | [**GetRefundsForCustomer**](docs/CustomersApi.md#getrefundsforcustomer) | **Get** /customer/{customerId}/refund | List Customer Refunds
*CustomersApi* | [**ListPaymentDetails**](docs/CustomersApi.md#listpaymentdetails) | **Get** /customer/{customerId}/payment-methods | List Customer&#x27;s Payment Details
*CustomersApi* | [**RemoveSeatsSubscriptions**](docs/CustomersApi.md#removeseatssubscriptions) | **Post** /subscription/{subscriptionId}/seats/remove | Remove Seats
*CustomersApi* | [**UpdateCustomer**](docs/CustomersApi.md#updatecustomer) | **Put** /customer/{customerId} | Update
*InvoicesApi* | [**ChargeInvoice**](docs/InvoicesApi.md#chargeinvoice) | **Post** /invoice/{invoiceId}/charge | Charge Invoice
*InvoicesApi* | [**DownloadInvoice**](docs/InvoicesApi.md#downloadinvoice) | **Get** /invoice/{invoiceId}/download | Download Invoice
*InvoicesApi* | [**GetInvoicesForCustomer**](docs/InvoicesApi.md#getinvoicesforcustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices
*PaymentDetailsApi* | [**CompleteFrontendPaymentDetails**](docs/PaymentDetailsApi.md#completefrontendpaymentdetails) | **Post** /customer/{customerId}/payment-methods/frontend-payment-token | Complete Frontend Detail Collection
*PaymentDetailsApi* | [**DeletePaymentDetails**](docs/PaymentDetailsApi.md#deletepaymentdetails) | **Delete** /payment-methods/{paymentDetailsId} | Delete
*PaymentDetailsApi* | [**DeletePaymentDetailsCustomer**](docs/PaymentDetailsApi.md#deletepaymentdetailscustomer) | **Delete** /customer/{customerId}/payment-methods/{paymentDetailsId} | Delete With Customer
*PaymentDetailsApi* | [**GetPaymentDetails**](docs/PaymentDetailsApi.md#getpaymentdetails) | **Get** /payment-methods/{paymentDetailsId} | Fetch
*PaymentDetailsApi* | [**ListPaymentDetails**](docs/PaymentDetailsApi.md#listpaymentdetails) | **Get** /customer/{customerId}/payment-methods | List Customer&#x27;s Payment Details
*PaymentDetailsApi* | [**MakeDefaultPaymentDetails**](docs/PaymentDetailsApi.md#makedefaultpaymentdetails) | **Post** /payment-methods/{paymentDetailsId}/default | Make Default
*PaymentDetailsApi* | [**MakeDefaultPaymentDetailsCustomer**](docs/PaymentDetailsApi.md#makedefaultpaymentdetailscustomer) | **Post** /customer/{customerId}/payment-methods/{paymentDetailsId}/default | Make Default With Customer
*PaymentDetailsApi* | [**StartFrontendPaymentDetails**](docs/PaymentDetailsApi.md#startfrontendpaymentdetails) | **Get** /customer/{customerId}/payment-methods/frontend-payment-token | Start Frontend Detail Collection
*PaymentsApi* | [**ChargeInvoice**](docs/PaymentsApi.md#chargeinvoice) | **Post** /invoice/{invoiceId}/charge | Charge Invoice
*PaymentsApi* | [**DownloadInvoice**](docs/PaymentsApi.md#downloadinvoice) | **Get** /invoice/{invoiceId}/download | Download Invoice
*PaymentsApi* | [**DownloadReceipt**](docs/PaymentsApi.md#downloadreceipt) | **Get** /receipt/{receiptId}/download | Download Receipt
*PaymentsApi* | [**GetInvoicesForCustomer**](docs/PaymentsApi.md#getinvoicesforcustomer) | **Get** /customer/{customerId}/invoices | List Customer Invoices
*PaymentsApi* | [**GetPaymentsForCustomer**](docs/PaymentsApi.md#getpaymentsforcustomer) | **Get** /customer/{customerId}/payment | List Customer Payments
*PaymentsApi* | [**ListPayment**](docs/PaymentsApi.md#listpayment) | **Get** /payment | List
*PaymentsApi* | [**RefundPayment**](docs/PaymentsApi.md#refundpayment) | **Post** /payment/{paymentId}/refund | Refund Payment
*PricesApi* | [**CreatePrice**](docs/PricesApi.md#createprice) | **Post** /product/{productId}/price | Create
*PricesApi* | [**ListPrice**](docs/PricesApi.md#listprice) | **Get** /product/{productId}/price | List
*ProductsApi* | [**CreateProduct**](docs/ProductsApi.md#createproduct) | **Post** /product | Create
*ProductsApi* | [**ListProduct**](docs/ProductsApi.md#listproduct) | **Get** /product | List
*ProductsApi* | [**ShowProductById**](docs/ProductsApi.md#showproductbyid) | **Get** /product/{productId} | Detail
*ProductsApi* | [**UpdateProduct**](docs/ProductsApi.md#updateproduct) | **Put** /product/{productId} | Update
*ReceiptApi* | [**DownloadReceipt**](docs/ReceiptApi.md#downloadreceipt) | **Get** /receipt/{receiptId}/download | Download Receipt
*RefundsApi* | [**GetRefundsForCustomer**](docs/RefundsApi.md#getrefundsforcustomer) | **Get** /customer/{customerId}/refund | List Customer Refunds
*RefundsApi* | [**ListRefund**](docs/RefundsApi.md#listrefund) | **Get** /refund | List
*RefundsApi* | [**ShowRefundById**](docs/RefundsApi.md#showrefundbyid) | **Get** /refund/{refundId} | Detail
*SubscriptionsApi* | [**AddSeatsSubscriptions**](docs/SubscriptionsApi.md#addseatssubscriptions) | **Post** /subscription/{subscriptionId}/seats/add | Add Seats
*SubscriptionsApi* | [**CancelSubscription**](docs/SubscriptionsApi.md#cancelsubscription) | **Post** /subscription/{subscriptionId}/cancel | Cancel Subscription
*SubscriptionsApi* | [**ChangeSubscriptionPrice**](docs/SubscriptionsApi.md#changesubscriptionprice) | **Post** /subscription/{subscriptionId}/price | Change Price
*SubscriptionsApi* | [**CreateSubscription**](docs/SubscriptionsApi.md#createsubscription) | **Post** /customer/{customerId}/subscription/start | Create Subscription
*SubscriptionsApi* | [**CustomerChangeSubscriptionPlan**](docs/SubscriptionsApi.md#customerchangesubscriptionplan) | **Post** /subscription/{subscriptionId}/plan | Change Subscription Plan
*SubscriptionsApi* | [**ExtendTrial**](docs/SubscriptionsApi.md#extendtrial) | **Post** /subscription/{subscriptionId}/extend | Extend Trial Subscription
*SubscriptionsApi* | [**GetActiveForCustomer**](docs/SubscriptionsApi.md#getactiveforcustomer) | **Get** /customer/{customerId}/subscription/active | List Customer Active Subscriptions
*SubscriptionsApi* | [**GetForCustomer**](docs/SubscriptionsApi.md#getforcustomer) | **Get** /customer/{customerId}/subscription | List Customer Subscriptions
*SubscriptionsApi* | [**ListSubscriptionPlans**](docs/SubscriptionsApi.md#listsubscriptionplans) | **Get** /subscription/plans | List Subscription Plans
*SubscriptionsApi* | [**ListSubscriptions**](docs/SubscriptionsApi.md#listsubscriptions) | **Get** /subscription | List
*SubscriptionsApi* | [**RemoveSeatsSubscriptions**](docs/SubscriptionsApi.md#removeseatssubscriptions) | **Post** /subscription/{subscriptionId}/seats/remove | Remove Seats
*SubscriptionsApi* | [**ShowSubscriptionById**](docs/SubscriptionsApi.md#showsubscriptionbyid) | **Get** /subscription/{subscriptionId} | Detail
*SubscriptionsApi* | [**StartTrial**](docs/SubscriptionsApi.md#starttrial) | **Post** /customer/{customerId}/subscription/trial | Start Trial Subscription For Customer

## Documentation For Models

 - [Address](docs/Address.md)
 - [BillingAdmin](docs/BillingAdmin.md)
 - [CheckoutBody](docs/CheckoutBody.md)
 - [CheckoutItems](docs/CheckoutItems.md)
 - [CheckoutSubscriptions](docs/CheckoutSubscriptions.md)
 - [Customer](docs/Customer.md)
 - [Feature](docs/Feature.md)
 - [FrontendToken](docs/FrontendToken.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2007Data](docs/InlineResponse2007Data.md)
 - [InlineResponse2007Receipts](docs/InlineResponse2007Receipts.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse201Lines](docs/InlineResponse201Lines.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [Invoice](docs/Invoice.md)
 - [InvoiceLines](docs/InvoiceLines.md)
 - [IssueRefundPayment](docs/IssueRefundPayment.md)
 - [Limit](docs/Limit.md)
 - [ModelError](docs/ModelError.md)
 - [PaymentDetails](docs/PaymentDetails.md)
 - [Price](docs/Price.md)
 - [Product](docs/Product.md)
 - [ProductTaxType](docs/ProductTaxType.md)
 - [Refund](docs/Refund.md)
 - [SeatsAddBody](docs/SeatsAddBody.md)
 - [SeatsRemoveBody](docs/SeatsRemoveBody.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionIdCancelBody](docs/SubscriptionIdCancelBody.md)
 - [SubscriptionIdExtendBody](docs/SubscriptionIdExtendBody.md)
 - [SubscriptionIdPlanBody](docs/SubscriptionIdPlanBody.md)
 - [SubscriptionIdPriceBody](docs/SubscriptionIdPriceBody.md)
 - [SubscriptionPlan](docs/SubscriptionPlan.md)
 - [SubscriptionStartBody](docs/SubscriptionStartBody.md)
 - [SubscriptionTrialBody](docs/SubscriptionTrialBody.md)
 - [VoucherCode](docs/VoucherCode.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@billabear.com
