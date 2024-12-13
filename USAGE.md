<!-- Start SDK Example Usage [usage] -->
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

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->