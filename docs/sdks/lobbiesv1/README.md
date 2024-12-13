# LobbiesV1
(*LobbiesV1*)

## Overview

### Available Operations

* [~~CreatePrivateDeprecated~~](#createprivatedeprecated) - CreatePrivateLobbyDeprecated :warning: **Deprecated**
* [~~CreatePublic~~](#createpublic) - CreatePublicLobbyDeprecated :warning: **Deprecated**
* [~~ListActive~~](#listactive) - ListActivePublicLobbiesDeprecatedV1 :warning: **Deprecated**

## ~~CreatePrivateDeprecated~~

CreatePrivateLobbyDeprecated

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"os"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.LobbiesV1.CreatePrivateDeprecated(ctx, operations.CreatePrivateLobbyDeprecatedSecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `security`                                                                                                         | [operations.CreatePrivateLobbyDeprecatedSecurity](../../models/operations/createprivatelobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |                                                                                                                    |
| `appID`                                                                                                            | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                           |
| `region`                                                                                                           | [*components.Region](../../models/components/region.md)                                                            | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `local`                                                                                                            | **bool*                                                                                                            | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*string](../../.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreatePublic~~

CreatePublicLobbyDeprecated

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"os"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.LobbiesV1.CreatePublic(ctx, operations.CreatePublicLobbyDeprecatedSecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `security`                                                                                                       | [operations.CreatePublicLobbyDeprecatedSecurity](../../models/operations/createpubliclobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |                                                                                                                  |
| `appID`                                                                                                          | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | N/A                                                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                         |
| `region`                                                                                                         | [*components.Region](../../models/components/region.md)                                                          | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `local`                                                                                                          | **bool*                                                                                                          | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*string](../../.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~ListActive~~

ListActivePublicLobbiesDeprecatedV1

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.LobbiesV1.ListActive(ctx, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil, nil)
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
| `local`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `region`                                                 | [*components.Region](../../models/components/region.md)  | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[[]components.Lobby](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 422, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |