# CreateRoomParams


## Fields

| Field                                                                                                                                                                                                         | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   | Example                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DeploymentID`                                                                                                                                                                                                | **string*                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                            | System generated id for a deployment.                                                                                                                                                                         | dep-6d4c6a71-2d75-4b42-94e1-f312f57f33c5                                                                                                                                                                      |
| `ClientIPs`                                                                                                                                                                                                   | []*string*                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                            | N/A                                                                                                                                                                                                           |                                                                                                                                                                                                               |
| `RoomConfig`                                                                                                                                                                                                  | **string*                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                            | Optional configuration parameters for the room. Can be any string including stringified JSON. It is accessible from the room via [`GetRoomInfo()`](https://hathora.dev/api#tag/RoomV2/operation/GetRoomInfo). | {"name":"my-room"}                                                                                                                                                                                            |
| `Region`                                                                                                                                                                                                      | [components.Region](../../models/components/region.md)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                            | N/A                                                                                                                                                                                                           |                                                                                                                                                                                                               |