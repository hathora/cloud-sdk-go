# DeploymentsV3
(*DeploymentsV3*)

## Overview

Operations that allow you configure and manage an application's [build](https://hathora.dev/docs/concepts/hathora-entities#build) at runtime.

### Available Operations

* [GetDeployments](#getdeployments) - GetDeployments
* [CreateDeployment](#createdeployment) - CreateDeployment
* [GetLatestDeployment](#getlatestdeployment) - GetLatestDeployment
* [GetDeployment](#getdeployment) - GetDeployment

## GetDeployments

Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application), optionally filtered by deploymentTag or buildTag.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetDeployments" method="get" path="/deployments/v3/apps/{appId}/deployments" -->
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

    res, err := s.DeploymentsV3.GetDeployments(ctx, hathoracloud.String("alpha"), hathoracloud.String("0.1.14-14c793"))
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
| `deploymentTag`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | alpha                                                    |
| `buildTag`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | 0.1.14-14c793                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeploymentsV3Page](../../models/components/deploymentsv3page.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## CreateDeployment

Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="CreateDeployment" method="post" path="/deployments/v3/apps/{appId}/deployments" -->
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

    res, err := s.DeploymentsV3.CreateDeployment(ctx, components.DeploymentConfigV3{
        DeploymentTag: hathoracloud.String("alpha"),
        IdleTimeoutEnabled: true,
        Env: []components.DeploymentConfigV3Env{
            components.DeploymentConfigV3Env{
                Value: "TRUE",
                Name: "EULA",
            },
            components.DeploymentConfigV3Env{
                Value: "TRUE",
                Name: "EULA",
            },
        },
        RoomsPerProcess: 3,
        AdditionalContainerPorts: []components.ContainerPort{
            components.ContainerPort{
                TransportType: components.TransportTypeUDP,
                Port: 8000,
                Name: "default",
            },
            components.ContainerPort{
                TransportType: components.TransportTypeUDP,
                Port: 8000,
                Name: "default",
            },
        },
        TransportType: components.TransportTypeUDP,
        ContainerPort: 4000,
        ExperimentalRequestedGPU: hathoracloud.Float64(1),
        RequestedMemoryMB: 1024,
        RequestedCPU: 0.5,
        BuildID: "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `deploymentConfigV3`                                                           | [components.DeploymentConfigV3](../../models/components/deploymentconfigv3.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `appID`                                                                        | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            | app-af469a92-5b45-4565-b3c4-b79878de67d2                                       |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*components.DeploymentV3](../../models/components/deploymentv3.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 400, 401, 404, 422, 429 | application/json        |
| errors.APIError         | 500                     | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## GetLatestDeployment

Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

### Example Usage

<!-- UsageSnippet language="go" operationID="GetLatestDeployment" method="get" path="/deployments/v3/apps/{appId}/deployments/latest" -->
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

    res, err := s.DeploymentsV3.GetLatestDeployment(ctx)
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

**[*components.DeploymentV3](../../models/components/deploymentv3.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetDeployment

Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment).

### Example Usage

<!-- UsageSnippet language="go" operationID="GetDeployment" method="get" path="/deployments/v3/apps/{appId}/deployments/{deploymentId}" -->
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

    res, err := s.DeploymentsV3.GetDeployment(ctx, "dep-6d4c6a71-2d75-4b42-94e1-f312f57f33c5")
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
| `deploymentID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | dep-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                 |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeploymentV3](../../models/components/deploymentv3.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |