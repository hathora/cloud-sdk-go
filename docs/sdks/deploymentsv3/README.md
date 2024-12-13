# DeploymentsV3
(*DeploymentsV3*)

## Overview

Operations that allow you configure and manage an application's [build](https://hathora.dev/docs/concepts/hathora-entities#build) at runtime.

### Available Operations

* [Create](#create) - CreateDeployment

## Create

Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected.

### Example Usage

```go
package main

import(
	"context"
	"os"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.DeploymentsV3.Create(ctx, components.DeploymentConfigV3{
        DeploymentTag: cloudsdkgo.String("alpha"),
        IdleTimeoutEnabled: false,
        Env: []components.DeploymentConfigV3Env{
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
        },
        TransportType: components.TransportTypeTCP,
        ContainerPort: 4000,
        RequestedMemoryMB: 1024,
        RequestedCPU: 0.5,
        BuildID: "bld-6d4c6a71-2d75-4b42-94e1-f312f57f33c5",
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 422, 429, 500 | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |