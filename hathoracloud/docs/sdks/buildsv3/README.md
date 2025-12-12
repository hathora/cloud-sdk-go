# BuildsV3

## Overview

Operations that allow you create and manage your [builds](https://hathora.dev/docs/concepts/hathora-entities#build).

### Available Operations

* [GetBuilds](#getbuilds) - GetBuilds
* [CreateBuild](#createbuild) - CreateBuild
* [GetBuild](#getbuild) - GetBuild
* [DeleteBuild](#deletebuild) - DeleteBuild
* [CreateBuildRegistry](#createbuildregistry) - CreateBuildRegistry
* [RunBuild](#runbuild) - RunBuild
* [RunBuildRegistry](#runbuildregistry) - RunBuildRegistry

## GetBuilds

Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

### Example Usage

<!-- UsageSnippet language="go" operationID="GetBuilds" method="get" path="/builds/v3/builds" -->
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

    res, err := s.BuildsV3.GetBuilds(ctx)
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

**[*components.BuildsV3Page](../../models/components/buildsv3page.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 408, 422, 429 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## CreateBuild

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) with optional `multipartUploadUrls` that can be used to upload larger builds in parts before calling `runBuild`. Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

### Example Usage

<!-- UsageSnippet language="go" operationID="CreateBuild" method="post" path="/builds/v3/builds" -->
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

    res, err := s.BuildsV3.CreateBuild(ctx, components.CreateMultipartBuildParams{
        BuildID: hathoracloud.Pointer("bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5"),
        BuildTag: hathoracloud.Pointer("0.1.14-14c793"),
        BuildSizeInBytes: 2645.24,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `createMultipartBuildParams`                                                                   | [components.CreateMultipartBuildParams](../../models/components/createmultipartbuildparams.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `orgID`                                                                                        | **string*                                                                                      | :heavy_minus_sign:                                                                             | N/A                                                                                            | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                                       |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*components.CreatedBuildV3WithMultipartUrls](../../models/components/createdbuildv3withmultiparturls.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 408, 422, 429 | application/json             |
| errors.APIError              | 500                          | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |

## GetBuild

Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).

### Example Usage

<!-- UsageSnippet language="go" operationID="GetBuild" method="get" path="/builds/v3/builds/{buildId}" -->
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

    res, err := s.BuildsV3.GetBuild(ctx, "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5")
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
| `buildID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                 |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.BuildV3](../../models/components/buildv3.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 408, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## DeleteBuild

Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted.
Be careful which builds you delete. This endpoint does not prevent you from deleting actively used builds.
Deleting a build that is actively build used by an app's deployment will cause failures when creating rooms.

### Example Usage

<!-- UsageSnippet language="go" operationID="DeleteBuild" method="delete" path="/builds/v3/builds/{buildId}" -->
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

    res, err := s.BuildsV3.DeleteBuild(ctx, "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5")
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
| `buildID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                 |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeletedBuild](../../models/components/deletedbuild.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 408, 422, 429 | application/json        |
| errors.APIError         | 500                     | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## CreateBuildRegistry

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) to be used with `runBuildRegistry`. Responds with a `buildId` that you must pass to [`RunBuildRegistry()`](https://hathora.dev/api#tag/BuildV3/operation/RunBuildRegistry) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

### Example Usage

<!-- UsageSnippet language="go" operationID="CreateBuildRegistry" method="post" path="/builds/v3/builds/registry" -->
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

    res, err := s.BuildsV3.CreateBuildRegistry(ctx, components.CreateBuildV3Params{
        BuildID: hathoracloud.Pointer("bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5"),
        BuildTag: hathoracloud.Pointer("0.1.14-14c793"),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `createBuildV3Params`                                                            | [components.CreateBuildV3Params](../../models/components/createbuildv3params.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `orgID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                         |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*components.BuildV3](../../models/components/buildv3.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 408, 422, 429 | application/json             |
| errors.APIError              | 500                          | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |

## RunBuild

Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

### Example Usage

<!-- UsageSnippet language="go" operationID="RunBuild" method="post" path="/builds/v3/builds/{buildId}/run" -->
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

    res, err := s.BuildsV3.RunBuild(ctx, "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5")
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
| `buildID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                 |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[io.ReadCloser](../../.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 408, 422, 429 | application/json             |
| errors.APIError              | 500                          | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |

## RunBuildRegistry

Builds a game server artifact from a public or private registry. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

### Example Usage

<!-- UsageSnippet language="go" operationID="RunBuildRegistry" method="post" path="/builds/v3/builds/{buildId}/runRegistry" -->
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

    res, err := s.BuildsV3.RunBuildRegistry(ctx, "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5", components.RegistryConfig{
        Image: "https://picsum.photos/seed/3gDPgtj/723/1525",
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `buildID`                                                              | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    | bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                               |
| `registryConfig`                                                       | [components.RegistryConfig](../../models/components/registryconfig.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `orgID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                               |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[io.ReadCloser](../../.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 408, 422, 429 | application/json             |
| errors.APIError              | 500                          | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |