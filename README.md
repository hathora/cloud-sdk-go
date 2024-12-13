# github.com/hathora/cloud-sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/hathora/cloud-sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/hathora/cloud-sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/hathora/hathora). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Hathora Cloud API: Welcome to the Hathora Cloud API documentation! Learn how to use the Hathora Cloud APIs to build and scale your game servers globally.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/hathora/cloud-sdk-go](#githubcomhathoracloud-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/hathora/cloud-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name              | Type | Scheme      | Environment Variable        |
| ----------------- | ---- | ----------- | --------------------------- |
| `HathoraDevToken` | http | HTTP Bearer | `HATHORA_HATHORA_DEV_TOKEN` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/components"
	"github.com/hathora/cloud-sdk-go/models/operations"
	"log"
	"os"
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
		Region:     components.RegionSeattle,
	}, cloudsdkgo.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), cloudsdkgo.String("LFG4"), cloudsdkgo.String("2swovpy1fnunu"))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Apps](docs/sdks/apps/README.md)

* [~~List~~](docs/sdks/apps/README.md#list) - GetAppsV1Deprecated :warning: **Deprecated**
* [~~CreateDeprecated~~](docs/sdks/apps/README.md#createdeprecated) - CreateAppV1Deprecated :warning: **Deprecated**
* [~~Update~~](docs/sdks/apps/README.md#update) - UpdateAppV1Deprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/apps/README.md#getinfo) - GetAppInfoV1Deprecated :warning: **Deprecated**
* [GetAll](docs/sdks/apps/README.md#getall) - GetApps
* [Create](docs/sdks/apps/README.md#create) - CreateApp
* [Delete](docs/sdks/apps/README.md#delete) - DeleteApp

### [~~AppsV1~~](docs/sdks/appsv1/README.md)

* [~~Remove~~](docs/sdks/appsv1/README.md#remove) - DeleteAppV1Deprecated :warning: **Deprecated**

### [AppsV2](docs/sdks/appsv2/README.md)

* [Update](docs/sdks/appsv2/README.md#update) - UpdateApp
* [Get](docs/sdks/appsv2/README.md#get) - GetApp

### [Auth](docs/sdks/auth/README.md)

* [LoginAnonymous](docs/sdks/auth/README.md#loginanonymous) - LoginAnonymous
* [LoginWithNickname](docs/sdks/auth/README.md#loginwithnickname) - LoginNickname
* [LoginGoogle](docs/sdks/auth/README.md#logingoogle) - LoginGoogle

### [Billing](docs/sdks/billing/README.md)

* [Get](docs/sdks/billing/README.md#get) - GetBalance
* [GetUpcomingItems](docs/sdks/billing/README.md#getupcomingitems) - GetUpcomingInvoiceItems
* [GetUpcomingTotal](docs/sdks/billing/README.md#getupcomingtotal) - GetUpcomingInvoiceTotal
* [GetPaymentMethod](docs/sdks/billing/README.md#getpaymentmethod) - GetPaymentMethod
* [InitStripeCustomerPortalURL](docs/sdks/billing/README.md#initstripecustomerportalurl) - InitStripeCustomerPortalUrl
* [GetInvoices](docs/sdks/billing/README.md#getinvoices) - GetInvoices

### [Builds](docs/sdks/builds/README.md)

* [~~ListDeprecated~~](docs/sdks/builds/README.md#listdeprecated) - GetBuildsDeprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/builds/README.md#getinfo) - GetBuildInfoDeprecated :warning: **Deprecated**
* [~~DeleteDeprecated~~](docs/sdks/builds/README.md#deletedeprecated) - DeleteBuildDeprecated :warning: **Deprecated**
* [~~List~~](docs/sdks/builds/README.md#list) - GetBuildsV2Deprecated :warning: **Deprecated**
* [~~CreateMultipart~~](docs/sdks/builds/README.md#createmultipart) - CreateWithMultipartUploadsV2Deprecated :warning: **Deprecated**
* [~~Delete~~](docs/sdks/builds/README.md#delete) - DeleteBuildV2Deprecated :warning: **Deprecated**
* [Get](docs/sdks/builds/README.md#get) - GetBuild
* [~~Run~~](docs/sdks/builds/README.md#run) - RunBuildDeprecated :warning: **Deprecated**

#### [~~Builds.V2~~](docs/sdks/v2/README.md)

* [~~CreateDeprecated~~](docs/sdks/v2/README.md#createdeprecated) - CreateBuildV2Deprecated :warning: **Deprecated**
* [~~CreateWithUploadURLDeprecated~~](docs/sdks/v2/README.md#createwithuploadurldeprecated) - CreateBuildWithUploadUrlV2Deprecated :warning: **Deprecated**

#### [Builds.V3](docs/sdks/v3/README.md)

* [Run](docs/sdks/v3/README.md#run) - RunBuild

### [~~BuildsV1~~](docs/sdks/buildsv1/README.md)

* [~~CreateDeprecated~~](docs/sdks/buildsv1/README.md#createdeprecated) - CreateBuildDeprecated :warning: **Deprecated**

### [~~BuildsV2~~](docs/sdks/buildsv2/README.md)

* [~~GetInfo~~](docs/sdks/buildsv2/README.md#getinfo) - GetBuildInfoV2Deprecated :warning: **Deprecated**
* [~~RunDeprecated~~](docs/sdks/buildsv2/README.md#rundeprecated) - RunBuildV2Deprecated :warning: **Deprecated**

### [BuildsV3](docs/sdks/buildsv3/README.md)

* [List](docs/sdks/buildsv3/README.md#list) - GetBuilds
* [Create](docs/sdks/buildsv3/README.md#create) - CreateBuild
* [Remove](docs/sdks/buildsv3/README.md#remove) - DeleteBuild

### [Deployments](docs/sdks/deployments/README.md)

* [~~GetAll~~](docs/sdks/deployments/README.md#getall) - GetDeploymentsV1Deprecated :warning: **Deprecated**
* [~~Create~~](docs/sdks/deployments/README.md#create) - CreateDeploymentV1Deprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/deployments/README.md#getinfo) - GetDeploymentInfoV2Deprecated :warning: **Deprecated**
* [~~CreateDeprecated~~](docs/sdks/deployments/README.md#createdeprecated) - CreateDeploymentV2Deprecated :warning: **Deprecated**
* [List](docs/sdks/deployments/README.md#list) - GetDeployments
* [GetLatest](docs/sdks/deployments/README.md#getlatest) - GetLatestDeployment
* [Get](docs/sdks/deployments/README.md#get) - GetDeployment

### [~~DeploymentsV1~~](docs/sdks/deploymentsv1/README.md)

* [~~GetLatestDeprecated~~](docs/sdks/deploymentsv1/README.md#getlatestdeprecated) - GetLatestDeploymentV1Deprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/deploymentsv1/README.md#getinfo) - GetDeploymentInfoV1Deprecated :warning: **Deprecated**

### [~~DeploymentsV2~~](docs/sdks/deploymentsv2/README.md)

* [~~ListDeprecated~~](docs/sdks/deploymentsv2/README.md#listdeprecated) - GetDeploymentsV2Deprecated :warning: **Deprecated**
* [~~GetLatestDeprecated~~](docs/sdks/deploymentsv2/README.md#getlatestdeprecated) - GetLatestDeploymentV2Deprecated :warning: **Deprecated**

### [DeploymentsV3](docs/sdks/deploymentsv3/README.md)

* [Create](docs/sdks/deploymentsv3/README.md#create) - CreateDeployment

### [~~Discovery~~](docs/sdks/discovery/README.md)

* [~~GetPingEndpoints~~](docs/sdks/discovery/README.md#getpingendpoints) - GetPingServiceEndpointsDeprecated :warning: **Deprecated**

### [DiscoveryV2](docs/sdks/discoveryv2/README.md)

* [GetPingEndpoints](docs/sdks/discoveryv2/README.md#getpingendpoints) - GetPingServiceEndpoints

### [Fleets](docs/sdks/fleets/README.md)

* [GetAll](docs/sdks/fleets/README.md#getall) - GetFleets
* [GetRegion](docs/sdks/fleets/README.md#getregion) - GetFleetRegion
* [UpdateRegion](docs/sdks/fleets/README.md#updateregion) - UpdateFleetRegion
* [GetMetrics](docs/sdks/fleets/README.md#getmetrics) - GetFleetMetrics


### [Lobbies](docs/sdks/lobbies/README.md)

* [~~CreatePrivate~~](docs/sdks/lobbies/README.md#createprivate) - CreatePrivateLobby :warning: **Deprecated**
* [~~CreatePublic~~](docs/sdks/lobbies/README.md#createpublic) - CreatePublicLobby :warning: **Deprecated**
* [~~CreateLocal~~](docs/sdks/lobbies/README.md#createlocal) - CreateLocalLobby :warning: **Deprecated**
* [~~ListActivePublic~~](docs/sdks/lobbies/README.md#listactivepublic) - ListActivePublicLobbiesDeprecatedV2 :warning: **Deprecated**
* [~~SetState~~](docs/sdks/lobbies/README.md#setstate) - SetLobbyState :warning: **Deprecated**
* [Create](docs/sdks/lobbies/README.md#create) - CreateLobby
* [ListPublic](docs/sdks/lobbies/README.md#listpublic) - ListActivePublicLobbies
* [GetInfoByRoomID](docs/sdks/lobbies/README.md#getinfobyroomid) - GetLobbyInfoByRoomId
* [GetByShortCode](docs/sdks/lobbies/README.md#getbyshortcode) - GetLobbyInfoByShortCode

### [~~LobbiesV1~~](docs/sdks/lobbiesv1/README.md)

* [~~CreatePrivateDeprecated~~](docs/sdks/lobbiesv1/README.md#createprivatedeprecated) - CreatePrivateLobbyDeprecated :warning: **Deprecated**
* [~~CreatePublic~~](docs/sdks/lobbiesv1/README.md#createpublic) - CreatePublicLobbyDeprecated :warning: **Deprecated**
* [~~ListActive~~](docs/sdks/lobbiesv1/README.md#listactive) - ListActivePublicLobbiesDeprecatedV1 :warning: **Deprecated**

### [~~LobbiesV2~~](docs/sdks/lobbiesv2/README.md)

* [~~Create~~](docs/sdks/lobbiesv2/README.md#create) - CreateLobbyDeprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/lobbiesv2/README.md#getinfo) - GetLobbyInfo :warning: **Deprecated**

### [Logs](docs/sdks/logs/README.md)

* [GetForProcess](docs/sdks/logs/README.md#getforprocess) - GetLogsForProcess
* [Download](docs/sdks/logs/README.md#download) - DownloadLogForProcess

### [Management](docs/sdks/management/README.md)

* [SendVerificationEmail](docs/sdks/management/README.md#sendverificationemail) - SendVerificationEmail

### [~~Metrics~~](docs/sdks/metrics/README.md)

* [~~GetProcessMetrics~~](docs/sdks/metrics/README.md#getprocessmetrics) - GetMetricsDeprecated :warning: **Deprecated**

### [Organizations](docs/sdks/organizations/README.md)

* [Get](docs/sdks/organizations/README.md#get) - GetOrgs
* [GetUserPendingInvites](docs/sdks/organizations/README.md#getuserpendinginvites) - GetUserPendingInvites
* [GetMembers](docs/sdks/organizations/README.md#getmembers) - GetOrgMembers
* [InviteUser](docs/sdks/organizations/README.md#inviteuser) - InviteUser
* [RescindInvite](docs/sdks/organizations/README.md#rescindinvite) - RescindInvite
* [GetUsageLimits](docs/sdks/organizations/README.md#getusagelimits) - GetUsageLimits

#### [Organizations.Invites](docs/sdks/invites/README.md)

* [Update](docs/sdks/invites/README.md#update) - UpdateUserInvite
* [Reject](docs/sdks/invites/README.md#reject) - RejectInvite

### [OrganizationsV1](docs/sdks/organizationsv1/README.md)

* [GetPendingInvites](docs/sdks/organizationsv1/README.md#getpendinginvites) - GetOrgPendingInvites

### [Orgs](docs/sdks/orgs/README.md)

* [AcceptInvite](docs/sdks/orgs/README.md#acceptinvite) - AcceptInvite

### [Processes](docs/sdks/processes/README.md)

* [~~ListRunning~~](docs/sdks/processes/README.md#listrunning) - GetRunningProcesses :warning: **Deprecated**
* [~~GetStopped~~](docs/sdks/processes/README.md#getstopped) - GetStoppedProcesses :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/processes/README.md#getinfo) - GetProcessInfoV2Deprecated :warning: **Deprecated**
* [~~ListLatest~~](docs/sdks/processes/README.md#listlatest) - GetLatestProcessesV2Deprecated :warning: **Deprecated**
* [~~Stop~~](docs/sdks/processes/README.md#stop) - StopProcessV2Deprecated :warning: **Deprecated**
* [GetLatest](docs/sdks/processes/README.md#getlatest) - GetLatestProcesses
* [Count](docs/sdks/processes/README.md#count) - GetProcessesCountExperimental
* [Create](docs/sdks/processes/README.md#create) - CreateProcess
* [Get](docs/sdks/processes/README.md#get) - GetProcess

#### [Processes.Metrics](docs/sdks/hathoracloudmetrics/README.md)

* [Get](docs/sdks/hathoracloudmetrics/README.md#get) - GetProcessMetrics

### [~~ProcessesV1~~](docs/sdks/processesv1/README.md)

* [~~GetInfo~~](docs/sdks/processesv1/README.md#getinfo) - GetProcessInfoDeprecated :warning: **Deprecated**

### [~~ProcessesV2~~](docs/sdks/processesv2/README.md)

* [~~CountProcesses~~](docs/sdks/processesv2/README.md#countprocesses) - GetProcessesCountExperimentalV2Deprecated :warning: **Deprecated**
* [~~CreateProcess~~](docs/sdks/processesv2/README.md#createprocess) - CreateProcessV2Deprecated :warning: **Deprecated**

### [ProcessesV3](docs/sdks/processesv3/README.md)

* [Stop](docs/sdks/processesv3/README.md#stop) - StopProcess

### [Rooms](docs/sdks/rooms/README.md)

* [~~CreateDeprecated~~](docs/sdks/rooms/README.md#createdeprecated) - CreateRoomDeprecated :warning: **Deprecated**
* [~~GetInfo~~](docs/sdks/rooms/README.md#getinfo) - GetRoomInfoDeprecated :warning: **Deprecated**
* [~~ListInactiveForProcess~~](docs/sdks/rooms/README.md#listinactiveforprocess) - GetInactiveRoomsForProcessDeprecated :warning: **Deprecated**
* [~~Suspend~~](docs/sdks/rooms/README.md#suspend) - SuspendRoomDeprecated :warning: **Deprecated**
* [~~GetConnectionInfoDeprecated~~](docs/sdks/rooms/README.md#getconnectioninfodeprecated) - GetConnectionInfoDeprecated :warning: **Deprecated**
* [Create](docs/sdks/rooms/README.md#create) - CreateRoom
* [GetActiveForProcess](docs/sdks/rooms/README.md#getactiveforprocess) - GetActiveRoomsForProcess
* [Destroy](docs/sdks/rooms/README.md#destroy) - DestroyRoom
* [GetConnectionInfo](docs/sdks/rooms/README.md#getconnectioninfo) - GetConnectionInfo
* [UpdateConfig](docs/sdks/rooms/README.md#updateconfig) - UpdateRoomConfig

#### [Rooms.Inactive](docs/sdks/inactive/README.md)

* [ListForProcess](docs/sdks/inactive/README.md#listforprocess) - GetInactiveRoomsForProcess

### [~~RoomsV1~~](docs/sdks/roomsv1/README.md)

* [~~GetActiveForProcessDeprecated~~](docs/sdks/roomsv1/README.md#getactiveforprocessdeprecated) - GetActiveRoomsForProcessDeprecated :warning: **Deprecated**
* [~~DestroyDeprecated~~](docs/sdks/roomsv1/README.md#destroydeprecated) - DestroyRoomDeprecated :warning: **Deprecated**

### [RoomsV2](docs/sdks/roomsv2/README.md)

* [GetInfo](docs/sdks/roomsv2/README.md#getinfo) - GetRoomInfo
* [~~Suspend~~](docs/sdks/roomsv2/README.md#suspend) - SuspendRoomV2Deprecated :warning: **Deprecated**

### [Tokens](docs/sdks/tokens/README.md)

* [GetOrg](docs/sdks/tokens/README.md#getorg) - GetOrgTokens
* [Create](docs/sdks/tokens/README.md#create) - CreateOrgToken
* [RevokeOrg](docs/sdks/tokens/README.md#revokeorg) - RevokeOrgToken

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

Certain parameters are configured globally. These parameters may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, These global values will be used as defaults on the operations that use them. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `orgId` to `"org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"` at SDK initialization and then you do not have to pass the same value on calls to operations like `GetOrg`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameters are available.
Global parameters can also be set via environment variable.

| Name  | Type   | Description          | Environment    |
| ----- | ------ | -------------------- | -------------- |
| OrgID | string | The OrgID parameter. | HATHORA_ORG_ID |
| AppID | string | The AppID parameter. | HATHORA_APP_ID |

### Example

```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Global Parameters [global-parameters] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `errors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetOrg` function may return the following errors:

| Error Type      | Status Code   | Content Type     |
| --------------- | ------------- | ---------------- |
| errors.APIError | 401, 404, 429 | application/json |
| errors.SDKError | 4XX, 5XX      | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"github.com/hathora/cloud-sdk-go/models/errors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {

		var e *errors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *errors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                    |
| --- | ------------------------- |
| 0   | `https://api.hathora.dev` |
| 1   | `https:///`               |

#### Example

```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithServerIndex(1),
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	cloudsdkgo "github.com/hathora/cloud-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := cloudsdkgo.New(
		cloudsdkgo.WithServerURL("https://api.hathora.dev"),
		cloudsdkgo.WithSecurity(os.Getenv("HATHORA_HATHORA_DEV_TOKEN")),
		cloudsdkgo.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		cloudsdkgo.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.Tokens.GetOrg(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/hathora/cloud-sdk-go&utm_campaign=go)
