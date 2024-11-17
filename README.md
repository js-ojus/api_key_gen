## `api_key_gen`
`api_key_gen` is a really tiny API credentials generator. It is both a library and a CLI program.

## Installation
Install it as usual with:

```bash
go get github.com/js-ojus/api_key_gen
```

Or, add it to your `go.mod`.

## Usage
### CLI
You can either compile the program in `cmd` or simply run it from source as:

```
go run cmd/main.go
```

### Library
The CLI program's `main.go` is a simple illustration of how the credentials can be generated using the library.

```go
package main

import (
	"fmt"

	gen "github.com/js-ojus/api_key_gen"
)

func main() {
	// Generate ClientID, ClientSecret and SecretHash.
	id, secret, hash, err := gen.Generate()
	if err != nil {
		fmt.Printf("error : %s", err.Error())
		return
	}

	fmt.Printf("%12s : %s\n", "ClientID", id)
	fmt.Printf("%12s : %s\n", "ClientSecret", secret)
	fmt.Printf("%12s : %s\n", "SecretHash", hash)
}
```
