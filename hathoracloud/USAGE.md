<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/hathora/cloud-sdk-go/hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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