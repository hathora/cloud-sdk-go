# OrganizationsV1

## Overview

### Available Operations

* [GetOrgs](#getorgs) - GetOrgs
* [GetUserPendingInvites](#getuserpendinginvites) - GetUserPendingInvites
* [GetOrgMembers](#getorgmembers) - GetOrgMembers
* [InviteUser](#inviteuser) - InviteUser
* [UpdateUserInvite](#updateuserinvite) - UpdateUserInvite
* [RescindInvite](#rescindinvite) - RescindInvite
* [GetOrgPendingInvites](#getorgpendinginvites) - GetOrgPendingInvites
* [AcceptInvite](#acceptinvite) - AcceptInvite
* [RejectInvite](#rejectinvite) - RejectInvite
* [GetUsageLimits](#getusagelimits) - GetUsageLimits

## GetOrgs

Returns an unsorted list of all organizations that you are a member of (an accepted membership invite). An organization is uniquely identified by an `orgId`.

### Example Usage

<!-- UsageSnippet language="go" operationID="GetOrgs" method="get" path="/orgs/v1" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.GetOrgs(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.OrgsPage](../../models/components/orgspage.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 408, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetUserPendingInvites

GetUserPendingInvites

### Example Usage

<!-- UsageSnippet language="go" operationID="GetUserPendingInvites" method="get" path="/orgs/v1/user/invites/pending" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.GetUserPendingInvites(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.PendingOrgInvitesPage](../../models/components/pendingorginvitespage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 408, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetOrgMembers

GetOrgMembers

### Example Usage

<!-- UsageSnippet language="go" operationID="GetOrgMembers" method="get" path="/orgs/v1/{orgId}/members" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.GetOrgMembers(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
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
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.OrgMembersPage](../../models/components/orgmemberspage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 408, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## InviteUser

InviteUser

### Example Usage

<!-- UsageSnippet language="go" operationID="InviteUser" method="put" path="/orgs/v1/{orgId}/invites" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.InviteUser(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39", components.CreateUserInvite{
        UserEmail: "noreply@hathora.dev",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `orgID`                                                                    | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                   |
| `createUserInvite`                                                         | [components.CreateUserInvite](../../models/components/createuserinvite.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*components.PendingOrgInvite](../../models/components/pendingorginvite.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 408, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## UpdateUserInvite

UpdateUserInvite

### Example Usage

<!-- UsageSnippet language="go" operationID="UpdateUserInvite" method="post" path="/orgs/v1/{orgId}/invites" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.UpdateUserInvite(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39", components.UpdateUserInvite{
        Scopes: components.CreateUpdateUserInviteScopesUserRole(
            components.UserRoleViewer,
        ),
        UserEmail: "noreply@hathora.dev",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `orgID`                                                                    | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                   |
| `updateUserInvite`                                                         | [components.UpdateUserInvite](../../models/components/updateuserinvite.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*bool](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 408, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## RescindInvite

RescindInvite

### Example Usage

<!-- UsageSnippet language="go" operationID="RescindInvite" method="post" path="/orgs/v1/{orgId}/invites/rescind" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    err := s.OrganizationsV1.RescindInvite(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39", components.RescindUserInvite{
        UserEmail: "noreply@hathora.dev",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `orgID`                                                                      | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                     |
| `rescindUserInvite`                                                          | [components.RescindUserInvite](../../models/components/rescinduserinvite.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 408, 422, 429 | application/json        |
| errors.APIError         | 500                     | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## GetOrgPendingInvites

GetOrgPendingInvites

### Example Usage

<!-- UsageSnippet language="go" operationID="GetOrgPendingInvites" method="get" path="/orgs/v1/{orgId}/invites/pending" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.OrganizationsV1.GetOrgPendingInvites(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
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
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.PendingOrgInvitesPage](../../models/components/pendingorginvitespage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 408, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## AcceptInvite

AcceptInvite

### Example Usage

<!-- UsageSnippet language="go" operationID="AcceptInvite" method="post" path="/orgs/v1/{orgId}/invites/accept" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    err := s.OrganizationsV1.AcceptInvite(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 408, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## RejectInvite

RejectInvite

### Example Usage

<!-- UsageSnippet language="go" operationID="RejectInvite" method="post" path="/orgs/v1/{orgId}/invites/reject" -->
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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    err := s.OrganizationsV1.RejectInvite(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 408, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetUsageLimits

GetUsageLimits

### Example Usage

<!-- UsageSnippet language="go" operationID="GetUsageLimits" method="get" path="/orgs/v1/metadata/usageLimits" -->
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

    res, err := s.OrganizationsV1.GetUsageLimits(ctx)
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

**[*components.UsageLimits](../../models/components/usagelimits.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 408, 429 | application/json   |
| errors.APIError    | 500                | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |