# Build

A build represents a game server artifact and its associated metadata.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `BuildTag`                                                                                                              | **string*                                                                                                               | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     | 0.1.14-14c793                                                                                                           |
| ~~`RegionalContainerTags`~~                                                                                             | [][components.RegionalContainerTags](../../models/components/regionalcontainertags.md)                                  | :heavy_check_mark:                                                                                                      | : warning: ** DEPRECATED **: This will be removed in a future release, please migrate away from it as soon as possible. |                                                                                                                         |
| `ImageSize`                                                                                                             | *int64*                                                                                                                 | :heavy_check_mark:                                                                                                      | The size (in bytes) of the Docker image built by Hathora.                                                               |                                                                                                                         |
| `Status`                                                                                                                | [components.BuildStatus](../../models/components/buildstatus.md)                                                        | :heavy_check_mark:                                                                                                      | N/A                                                                                                                     |                                                                                                                         |
| `DeletedAt`                                                                                                             | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | When the build was deleted.                                                                                             |                                                                                                                         |
| `FinishedAt`                                                                                                            | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) finished executing.                         |                                                                                                                         |
| `StartedAt`                                                                                                             | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) is called.                                  |                                                                                                                         |
| `CreatedAt`                                                                                                             | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | When [`CreateBuild()`](https://hathora.dev/api#tag/BuildV2/operation/CreateBuild) is called.                            |                                                                                                                         |
| `CreatedBy`                                                                                                             | *string*                                                                                                                | :heavy_check_mark:                                                                                                      | N/A                                                                                                                     | noreply@hathora.dev                                                                                                     |
| `BuildID`                                                                                                               | *int*                                                                                                                   | :heavy_check_mark:                                                                                                      | System generated id for a build. Increments by 1.                                                                       | 1                                                                                                                       |
| `AppID`                                                                                                                 | *string*                                                                                                                | :heavy_check_mark:                                                                                                      | System generated unique identifier for an application.                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                                |