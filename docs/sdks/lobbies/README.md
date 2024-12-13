# Lobbies
(*Lobbies*)

## Overview

### Available Operations

* [~~CreatePrivate~~](#createprivate) - CreatePrivateLobby :warning: **Deprecated**
* [~~CreatePublic~~](#createpublic) - CreatePublicLobby :warning: **Deprecated**
* [~~CreateLocal~~](#createlocal) - CreateLocalLobby :warning: **Deprecated**
* [~~ListActivePublic~~](#listactivepublic) - ListActivePublicLobbiesDeprecatedV2 :warning: **Deprecated**
* [~~SetState~~](#setstate) - SetLobbyState :warning: **Deprecated**
* [Create](#create) - CreateLobby
* [ListPublic](#listpublic) - ListActivePublicLobbies
* [GetInfoByRoomID](#getinfobyroomid) - GetLobbyInfoByRoomId
* [GetByShortCode](#getbyshortcode) - GetLobbyInfoByShortCode

## ~~CreatePrivate~~

CreatePrivateLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"os"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.Lobbies.CreatePrivate(ctx, operations.CreatePrivateLobbySecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, operations.CreatePrivateLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionLondon,
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), cloudsdkgo.String("2swovpy1fnunu"))
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
| `security`                                                                                           | [operations.CreatePrivateLobbySecurity](../../models/operations/createprivatelobbysecurity.md)       | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `requestBody`                                                                                        | [operations.CreatePrivateLobbyRequestBody](../../models/operations/createprivatelobbyrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreatePublic~~

CreatePublicLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"os"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.Lobbies.CreatePublic(ctx, operations.CreatePublicLobbySecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, operations.CreatePublicLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionJohannesburg,
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), cloudsdkgo.String("2swovpy1fnunu"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `security`                                                                                         | [operations.CreatePublicLobbySecurity](../../models/operations/createpubliclobbysecurity.md)       | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |                                                                                                    |
| `requestBody`                                                                                      | [operations.CreatePublicLobbyRequestBody](../../models/operations/createpubliclobbyrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `appID`                                                                                            | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                           |
| `roomID`                                                                                           | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | 2swovpy1fnunu                                                                                      |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreateLocal~~

CreateLocalLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"os"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := cloudsdkgo.New(
        cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.Lobbies.CreateLocal(ctx, operations.CreateLocalLobbySecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, operations.CreateLocalLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionJohannesburg,
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), cloudsdkgo.String("2swovpy1fnunu"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `security`                                                                                       | [operations.CreateLocalLobbySecurity](../../models/operations/createlocallobbysecurity.md)       | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |                                                                                                  |
| `requestBody`                                                                                    | [operations.CreateLocalLobbyRequestBody](../../models/operations/createlocallobbyrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `appID`                                                                                          | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                         |
| `roomID`                                                                                         | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | 2swovpy1fnunu                                                                                    |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~ListActivePublic~~

Get all active lobbies for a an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

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

    res, err := s.Lobbies.ListActivePublic(ctx, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `appID`                                                                                 | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                |
| `region`                                                                                | [*components.Region](../../models/components/region.md)                                 | :heavy_minus_sign:                                                                      | Region to filter by. If omitted, active public lobbies in all regions will be returned. |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[[]components.Lobby](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~SetState~~

Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Lobbies.SetState(ctx, "2swovpy1fnunu", components.SetLobbyStateParams{
        State: "South Dakota",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `roomID`                                                                         | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | 2swovpy1fnunu                                                                    |
| `setLobbyStateParams`                                                            | [components.SetLobbyStateParams](../../models/components/setlobbystateparams.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `appID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                         |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## Create

Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.

### Example Usage

```go
package main

import(
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
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

    res, err := s.Lobbies.Create(ctx, operations.CreateLobbySecurity{
        PlayerAuth: os.Getenv("HATHORA_PLAYER_AUTH"),
    }, components.CreateLobbyV3Params{
        Visibility: components.LobbyVisibilityPrivate,
        RoomConfig: cloudsdkgo.String("{\"name\":\"my-room\"}"),
        Region: components.RegionSeattle,
    }, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), cloudsdkgo.String("LFG4"), cloudsdkgo.String("2swovpy1fnunu"))
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
| `security`                                                                       | [operations.CreateLobbySecurity](../../models/operations/createlobbysecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |                                                                                  |
| `createLobbyV3Params`                                                            | [components.CreateLobbyV3Params](../../models/components/createlobbyv3params.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `appID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                         |
| `shortCode`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | LFG4                                                                             |
| `roomID`                                                                         | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | 2swovpy1fnunu                                                                    |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ListPublic

Get all active lobbies for a given [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

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

    res, err := s.Lobbies.ListPublic(ctx, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `appID`                                                            | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                           |
| `region`                                                           | [*components.Region](../../models/components/region.md)            | :heavy_minus_sign:                                                 | If omitted, active public lobbies in all regions will be returned. |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[[]components.LobbyV3](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetInfoByRoomID

Get details for a lobby.

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

    res, err := s.Lobbies.GetInfoByRoomID(ctx, "2swovpy1fnunu", cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 422, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetByShortCode

Get details for a lobby. If 2 or more lobbies have the same `shortCode`, then the most recently created lobby will be returned.

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

    res, err := s.Lobbies.GetByShortCode(ctx, "LFG4", cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `shortCode`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | LFG4                                                     |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |