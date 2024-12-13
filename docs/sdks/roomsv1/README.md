# RoomsV1
(*RoomsV1*)

## Overview

### Available Operations

* [~~GetActiveForProcessDeprecated~~](#getactiveforprocessdeprecated) - GetActiveRoomsForProcessDeprecated :warning: **Deprecated**
* [~~DestroyDeprecated~~](#destroydeprecated) - DestroyRoomDeprecated :warning: **Deprecated**

## ~~GetActiveForProcessDeprecated~~

GetActiveRoomsForProcessDeprecated

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"os"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.RoomsV1.GetActiveForProcessDeprecated(ctx, "cbfcddd2-0006-43ae-996c-995fff7bed2e", cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `processID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | cbfcddd2-0006-43ae-996c-995fff7bed2e                     |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[[]components.RoomWithoutAllocations](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~DestroyDeprecated~~

DestroyRoomDeprecated

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"os"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    err := s.RoomsV1.DestroyDeprecated(ctx, "2swovpy1fnunu", cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 429, 500 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |