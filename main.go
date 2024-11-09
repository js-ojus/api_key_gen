package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Regexp to remove non-alphanumeric characters from ClientID.
	re := regexp.MustCompile(`[-+/]`)

	// Generate the ClientID.
	buf1 := make([]byte, 20)
	_, err := rand.Read(buf1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var out1 bytes.Buffer
	enc := base64.NewEncoder(base64.RawStdEncoding, &out1)
	enc.Write(buf1)
	enc.Close()
	s1 := out1.String()
	s2 := strings.ToUpper(re.ReplaceAllString(s1, ""))[:20]
	fmt.Printf("%12s : %s\n", "ClientID", s2)

	// Generate the ClientSecret.
	buf2 := make([]byte, 30)
	_, err = rand.Read(buf2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var out2 bytes.Buffer
	enc = base64.NewEncoder(base64.RawStdEncoding, &out2)
	enc.Write(buf2)
	enc.Close()
	fmt.Printf("%12s : %s\n", "ClientSecret", out2.String())

	// Generate the hash of the ClientSecret.
	hash := sha256.Sum256(out2.Bytes())
	hashString := hex.EncodeToString(hash[:])
	fmt.Printf("%12s : %s\n", "SecretHash", hashString)
}
