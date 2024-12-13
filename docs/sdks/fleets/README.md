# Fleets
(*Fleets*)

## Overview

### Available Operations

* [GetAll](#getall) - GetFleets
* [GetRegion](#getregion) - GetFleetRegion
* [UpdateRegion](#updateregion) - UpdateFleetRegion
* [GetMetrics](#getmetrics) - GetFleetMetrics

## GetAll

Returns an array of [fleets](https://hathora.dev/docs/concepts/hathora-entities#fleet).

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

    res, err := s.Fleets.GetAll(ctx, cloudsdkgo.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
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

**[*components.FleetsPage](../../models/components/fleetspage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetRegion

Gets the configuration for a given [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

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

    res, err := s.Fleets.GetRegion(ctx, "<id>", components.RegionSaoPaulo, cloudsdkgo.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
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
| `fleetID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `region`                                                 | [components.Region](../../models/components/region.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.FleetRegion](../../models/components/fleetregion.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## UpdateRegion

Updates the configuration for a given [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

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

    err := s.Fleets.UpdateRegion(ctx, "<id>", components.RegionSingapore, components.FleetRegionConfig{
        CloudMinVcpus: 511402,
    }, cloudsdkgo.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `fleetID`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `region`                                                                     | [components.Region](../../models/components/region.md)                       | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `fleetRegionConfig`                                                          | [components.FleetRegionConfig](../../models/components/fleetregionconfig.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `orgID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 422, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## GetMetrics

Gets metrics for a [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

### Example Usage

```go
package main

import(
	"context"
	"os"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
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

    res, err := s.Fleets.GetMetrics(ctx, operations.GetFleetMetricsRequest{
        FleetID: "<id>",
        Region: components.RegionLondon,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetFleetMetricsRequest](../../models/operations/getfleetmetricsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.FleetMetricsData](../../models/components/fleetmetricsdata.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 422, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |