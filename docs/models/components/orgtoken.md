# OrgToken


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Scopes`                                                               | [*components.Scopes](../../models/components/scopes.md)                | :heavy_minus_sign:                                                     | If not defined, the token has Admin access.                            |                                                                        |
| `CreatedAt`                                                            | [time.Time](https://pkg.go.dev/time#Time)                              | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `CreatedBy`                                                            | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    | noreply@hathora.dev                                                    |
| `LastFourCharsOfKey`                                                   | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `Status`                                                               | [components.OrgTokenStatus](../../models/components/orgtokenstatus.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `Name`                                                                 | *string*                                                               | :heavy_check_mark:                                                     | Readable name for a token. Must be unique within an organization.      | ci-token                                                               |
| `OrgID`                                                                | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `OrgTokenID`                                                           | *string*                                                               | :heavy_check_mark:                                                     | System generated unique identifier for an organization token.          | org-token-af469a92-5b45-4565-b3c4-b79878de67d2                         |