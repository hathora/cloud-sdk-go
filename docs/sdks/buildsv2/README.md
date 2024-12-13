# BuildsV2
(*BuildsV2*)

## Overview

### Available Operations

* [~~GetInfo~~](#getinfo) - GetBuildInfoV2Deprecated :warning: **Deprecated**
* [~~RunDeprecated~~](#rundeprecated) - RunBuildV2Deprecated :warning: **Deprecated**

## ~~GetInfo~~

Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).

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

    res, err := s.BuildsV2.GetInfo(ctx, 1, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `buildID`                                                | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.Build](../../models/components/build.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~RunDeprecated~~

Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"os"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }


    res, err := s.BuildsV2.RunDeprecated(ctx, 1, operations.RunBuildV2DeprecatedRequestBody{}, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `buildID`                                                                                                | *int*                                                                                                    | :heavy_check_mark:                                                                                       | N/A                                                                                                      | 1                                                                                                        |
| `requestBody`                                                                                            | [operations.RunBuildV2DeprecatedRequestBody](../../models/operations/runbuildv2deprecatedrequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `appID`                                                                                                  | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                 |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[io.ReadCloser](../../.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 400, 401, 404, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |