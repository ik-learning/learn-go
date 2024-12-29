package main

import (
	b64 "encoding/base64"
	"fmt"

	"sigs.k8s.io/external-dns/endpoint"
)

func main() {
	keys := []string{
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=", // safe base64 url encoded 44 bytes and 32 when decoded
		// "01234567890123456789012345678901",             // plain txt 32 bytes
		// "passphrasewhichneedstobe32bytes!",             // plain txt 32 bytes
	}

	for _, k := range keys {
		key := []byte(k)
		if len(key) != 32 {
			var err error
			if key, err = b64.StdEncoding.DecodeString(string(key)); err != nil || len(key) != 32 {
				fmt.Errorf("the AES Encryption key must have a length of 32 byte")
			}
		}
		encrypted, _ := endpoint.EncryptText(
			"heritage=external-dns,external-dns/owner=example,external-dns/resource=ingress/default/example",
			key,
			nil,
		)
		fmt.Println("encrypted:", encrypted)

		decrypted, _, err := endpoint.DecryptText(encrypted, key)
		if err != nil {
			fmt.Println("Error decrypting:", err, "for key:", k)
		}
		fmt.Println("decrypted:", decrypted)
	}
}
