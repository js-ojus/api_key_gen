package main

import (
	"fmt"

	gen "api_key_gen"
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
