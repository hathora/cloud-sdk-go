# BillingV1
(*BillingV1*)

## Overview

 

### Available Operations

* [~~GetBalance~~](#getbalance) - GetBalance :warning: **Deprecated**
* [GetUpcomingInvoiceItems](#getupcominginvoiceitems) - GetUpcomingInvoiceItems
* [GetUpcomingInvoiceTotal](#getupcominginvoicetotal) - GetUpcomingInvoiceTotal
* [GetPaymentMethod](#getpaymentmethod) - GetPaymentMethod
* [InitStripeCustomerPortalURL](#initstripecustomerportalurl) - InitStripeCustomerPortalUrl
* [GetInvoices](#getinvoices) - GetInvoices

## ~~GetBalance~~

GetBalance

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetBalance" method="get" path="/billing/v1/balance" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.GetBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*float64](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetUpcomingInvoiceItems

GetUpcomingInvoiceItems

### Example Usage

<!-- UsageSnippet language="go" operationID="GetUpcomingInvoiceItems" method="get" path="/billing/v1/upcoming/items" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.GetUpcomingInvoiceItems(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.InvoiceItemPage](../../models/components/invoiceitempage.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetUpcomingInvoiceTotal

GetUpcomingInvoiceTotal

### Example Usage

<!-- UsageSnippet language="go" operationID="GetUpcomingInvoiceTotal" method="get" path="/billing/v1/upcoming/total" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.GetUpcomingInvoiceTotal(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUpcomingInvoiceTotalResponseBody](../../models/operations/getupcominginvoicetotalresponsebody.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetPaymentMethod

GetPaymentMethod

### Example Usage

<!-- UsageSnippet language="go" operationID="GetPaymentMethod" method="get" path="/billing/v1/paymentmethod" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.GetPaymentMethod(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.PaymentMethod](../../models/components/paymentmethod.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## InitStripeCustomerPortalURL

InitStripeCustomerPortalUrl

### Example Usage

<!-- UsageSnippet language="go" operationID="InitStripeCustomerPortalUrl" method="post" path="/billing/v1/customerportalurl" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.InitStripeCustomerPortalURL(ctx, components.CustomerPortalURL{
        ReturnURL: "https://sunny-majority.com/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `customerPortalURL`                                                          | [components.CustomerPortalURL](../../models/components/customerportalurl.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `orgID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*string](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetInvoices

GetInvoices

### Example Usage

<!-- UsageSnippet language="go" operationID="GetInvoices" method="get" path="/billing/v1/invoices" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BillingV1.GetInvoices(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[[]components.Invoice](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |