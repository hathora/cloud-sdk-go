# AppsV2
(*AppsV2*)

## Overview

Operations that allow you manage your [applications](https://hathora.dev/docs/concepts/hathora-entities#application).

### Available Operations

* [GetApps](#getapps) - GetApps
* [CreateApp](#createapp) - CreateApp
* [GetApp](#getapp) - GetApp
* [UpdateApp](#updateapp) - UpdateApp
* [DeleteApp](#deleteapp) - DeleteApp
* [PatchApp](#patchapp) - PatchApp

## GetApps

Returns an unsorted list of your organization’s [applications](https://hathora.dev/docs/concepts/hathora-entities#application). An application is uniquely identified by an `appId`.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetApps" method="get" path="/apps/v2/apps" -->
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

    res, err := s.AppsV2.GetApps(ctx)
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

**[*components.ApplicationsPage](../../models/components/applicationspage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## CreateApp

Create a new [application](https://hathora.dev/docs/concepts/hathora-entities#application).

### Example Usage

<!-- UsageSnippet language="go" operationID="CreateApp" method="post" path="/apps/v2/apps" -->
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

    res, err := s.AppsV2.CreateApp(ctx, components.CreateAppConfig{
        AuthConfiguration: components.AuthConfiguration{},
        AppName: "minecraft",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `createAppConfig`                                                        | [components.CreateAppConfig](../../models/components/createappconfig.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `orgID`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                 |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*components.Application](../../models/components/application.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetApp

Get details for an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetApp" method="get" path="/apps/v2/apps/{appId}" -->
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

    res, err := s.AppsV2.GetApp(ctx)
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

**[*components.Application](../../models/components/application.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## UpdateApp

Set application config (will override all fields) for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.

### Example Usage

<!-- UsageSnippet language="go" operationID="UpdateApp" method="post" path="/apps/v2/apps/{appId}" -->
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

    res, err := s.AppsV2.UpdateApp(ctx, components.AppConfigWithServiceConfig{
        ServiceConfig: &components.ServiceConfigWrite{
            StaticProcessAllocation: []components.StaticProcessAllocationConfigWrite{
                components.StaticProcessAllocationConfigWrite{
                    MaxProcesses: hathoracloud.Pointer[int](3),
                    TargetProcesses: hathoracloud.Pointer[int](2),
                    MinProcesses: 1,
                    Region: components.RegionSaoPaulo,
                },
                components.StaticProcessAllocationConfigWrite{
                    MaxProcesses: hathoracloud.Pointer[int](3),
                    TargetProcesses: hathoracloud.Pointer[int](2),
                    MinProcesses: 1,
                    Region: components.RegionSaoPaulo,
                },
            },
        },
        AuthConfiguration: components.AuthConfiguration{},
        AppName: "minecraft",
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
| `appConfigWithServiceConfig`                                                                   | [components.AppConfigWithServiceConfig](../../models/components/appconfigwithserviceconfig.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `appID`                                                                                        | **string*                                                                                      | :heavy_minus_sign:                                                                             | N/A                                                                                            | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                       |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*components.Application](../../models/components/application.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## DeleteApp

Delete an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`. Your organization will lose access to this application.

### Example Usage

<!-- UsageSnippet language="go" operationID="DeleteApp" method="delete" path="/apps/v2/apps/{appId}" -->
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

    err := s.AppsV2.DeleteApp(ctx)
    if err != nil {
        log.Fatal(err)
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

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## PatchApp

Patch data for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.

### Example Usage

<!-- UsageSnippet language="go" operationID="PatchApp" method="patch" path="/apps/v2/apps/{appId}" -->
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

    res, err := s.AppsV2.PatchApp(ctx, components.PartialAppConfigWithServiceConfig{
        ServiceConfig: &components.ServiceConfigWrite{
            StaticProcessAllocation: []components.StaticProcessAllocationConfigWrite{
                components.StaticProcessAllocationConfigWrite{
                    MaxProcesses: hathoracloud.Pointer[int](3),
                    TargetProcesses: hathoracloud.Pointer[int](2),
                    MinProcesses: 1,
                    Region: components.RegionTokyo,
                },
                components.StaticProcessAllocationConfigWrite{
                    MaxProcesses: hathoracloud.Pointer[int](3),
                    TargetProcesses: hathoracloud.Pointer[int](2),
                    MinProcesses: 1,
                    Region: components.RegionTokyo,
                },
                components.StaticProcessAllocationConfigWrite{
                    MaxProcesses: hathoracloud.Pointer[int](3),
                    TargetProcesses: hathoracloud.Pointer[int](2),
                    MinProcesses: 1,
                    Region: components.RegionTokyo,
                },
            },
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `partialAppConfigWithServiceConfig`                                                                          | [components.PartialAppConfigWithServiceConfig](../../models/components/partialappconfigwithserviceconfig.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `appID`                                                                                                      | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | N/A                                                                                                          | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                     |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*components.Application](../../models/components/application.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |