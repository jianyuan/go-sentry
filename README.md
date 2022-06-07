# go-sentry [![Go Reference](https://pkg.go.dev/badge/github.com/jianyuan/go-sentry/v2/sentry.svg)](https://pkg.go.dev/github.com/jianyuan/go-sentry/v2/sentry)

Go library for accessing the [Sentry Web API](https://docs.sentry.io/api/).

## Installation
go-sentry is compatible with modern Go releases in module mode, with Go installed:

```sh
go get github.com/jianyuan/go-sentry/v2/sentry
```

## Usage

```go
import "github.com/jianyuan/go-sentry/v2/sentry"
```

Create a new Sentry client. Then, use the various services on the client to access different parts of the
Sentry Web API. For example:

```go
client := sentry.NewClient(nil)

// List all organizations
orgs, _, err := client.Organizations.List(ctx, nil)
```

### Authentication

The library does not directly handle authentication. When creating a new client, pass an
`http.Client` that can handle authentication for you. We recommend the [oauth2](https://pkg.go.dev/golang.org/x/oauth2)
library. For example:

```go
package main

import (
	"github.com/jianyuan/go-sentry/v2/sentry"
	"golang.org/x/oauth2"
)

func main() {
    ctx := context.Background()
    tokenSrc := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "YOUR-API-KEY"},
    )
    httpClient := oauth2.NewClient(ctx, tokenSrc)
    
    client := sentry.NewClient(httpClient)
    
    // List all organizations
    orgs, _, err := client.Organizations.List(ctx, nil)
}
```

## Code structure
The code structure was inspired by [google/go-github](https://github.com/google/go-github).

## License
This library is distributed under the [MIT License](LICENSE).
