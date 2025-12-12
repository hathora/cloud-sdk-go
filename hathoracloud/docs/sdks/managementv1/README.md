# ManagementV1

## Overview

 

### Available Operations

* [SendVerificationEmail](#sendverificationemail) - SendVerificationEmail

## SendVerificationEmail

SendVerificationEmail

### Example Usage

<!-- UsageSnippet language="go" operationID="SendVerificationEmail" method="post" path="/management/v1/sendverificationemail" -->
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

    s := hathoracloud.New()

    res, err := s.ManagementV1.SendVerificationEmail(ctx, components.VerificationEmailRequest{
        UserID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.VerificationEmailRequest](../../models/components/verificationemailrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.VerificationEmailSuccess](../../models/components/verificationemailsuccess.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 408, 429    | application/json |
| errors.APIError  | 500              | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |