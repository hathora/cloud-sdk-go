# BuildsV1
(*BuildsV1*)

## Overview

### Available Operations

* [~~GetBuildsDeprecated~~](#getbuildsdeprecated) - GetBuildsDeprecated :warning: **Deprecated**
* [~~GetBuildInfoDeprecated~~](#getbuildinfodeprecated) - GetBuildInfoDeprecated :warning: **Deprecated**
* [~~CreateBuildDeprecated~~](#createbuilddeprecated) - CreateBuildDeprecated :warning: **Deprecated**
* [~~DeleteBuildDeprecated~~](#deletebuilddeprecated) - DeleteBuildDeprecated :warning: **Deprecated**
* [~~RunBuildDeprecated~~](#runbuilddeprecated) - RunBuildDeprecated :warning: **Deprecated**

## ~~GetBuildsDeprecated~~

Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetBuildsDeprecated" method="get" path="/builds/v1/{appId}/list" -->
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
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BuildsV1.GetBuildsDeprecated(ctx)
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
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[[]components.Build](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~GetBuildInfoDeprecated~~

Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetBuildInfoDeprecated" method="get" path="/builds/v1/{appId}/info/{buildId}" -->
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
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BuildsV1.GetBuildInfoDeprecated(ctx, 1)
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

## ~~CreateBuildDeprecated~~

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build). Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="CreateBuildDeprecated" method="post" path="/builds/v1/{appId}/create" -->
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
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.BuildsV1.CreateBuildDeprecated(ctx, components.CreateBuildParams{
        BuildTag: hathoracloud.String("0.1.14-14c793"),
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
| `createBuildParams`                                                          | [components.CreateBuildParams](../../models/components/createbuildparams.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `appID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | app-af469a92-5b45-4565-b3c4-b79878de67d2                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*components.Build](../../models/components/build.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~DeleteBuildDeprecated~~

Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="DeleteBuildDeprecated" method="delete" path="/builds/v1/{appId}/delete/{buildId}" -->
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
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    err := s.BuildsV1.DeleteBuildDeprecated(ctx, 1)
    if err != nil {
        log.Fatal(err)
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

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~RunBuildDeprecated~~

Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="RunBuildDeprecated" method="post" path="/builds/v1/{appId}/run/{buildId}" -->
```go
package main

import(
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"os"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := hathoracloud.New(
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.BuildsV1.RunBuildDeprecated(ctx, 1, operations.RunBuildDeprecatedRequestBody{
        File: operations.File{
            FileName: "example.file",
            Content: example,
        },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `buildID`                                                                                            | *int*                                                                                                | :heavy_check_mark:                                                                                   | N/A                                                                                                  | 1                                                                                                    |
| `requestBody`                                                                                        | [operations.RunBuildDeprecatedRequestBody](../../models/operations/runbuilddeprecatedrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*string](../../.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 400, 401, 404, 422, 429 | application/json        |
| errors.APIError         | 500                     | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |