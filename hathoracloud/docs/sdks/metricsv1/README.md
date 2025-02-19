# MetricsV1
(*MetricsV1*)

## Overview

Deprecated. Use [ProcessesV3#GetProcessMetrics](https://hathora.dev/api#tag/ProcessesV3/operation/GetProcessMetrics) to fetch metrics about a specific process.

### Available Operations

* [~~GetMetricsDeprecated~~](#getmetricsdeprecated) - GetMetricsDeprecated :warning: **Deprecated**

## ~~GetMetricsDeprecated~~

Get metrics for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := hathoracloud.New(
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.MetricsV1.GetMetricsDeprecated(ctx, operations.GetMetricsDeprecatedRequest{
        AppID: hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        ProcessID: "cbfcddd2-0006-43ae-996c-995fff7bed2e",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetMetricsDeprecatedRequest](../../models/operations/getmetricsdeprecatedrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*components.DeprecatedProcessMetricsData](../../models/components/deprecatedprocessmetricsdata.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |