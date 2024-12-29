package main

import (
	"crypto/aes"
	"fmt"
)

// https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d

var (
	// We're using a 32 byte long secret key
	secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func main() {
	// This will successfully encrypt.
	ciphertext := encrypt("This is some sensitive information")
	fmt.Printf("Ciphertext: %x \n", ciphertext)

	// This will cause an error since the
	// plaintext is less than 16 bytes.
	ciphertext = encrypt("Hello")
}

func encrypt(plaintext string) string {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	// Make a buffer the same length as plaintext
	ciphertext := make([]byte, len(plaintext))
	aes.Encrypt(ciphertext, []byte(plaintext))

	return string(ciphertext)
}

// func newCipher(key []byte) (cipher.Block, error) {
// 	if !supportsAES {
// 		return newCipherGeneric(key)
// 	}
// 	// Note that under certain circumstances, we only return the inner aesCipherAsm.
// 	// This avoids an unnecessary allocation of the aesCipher struct.
// 	c := aesCipherGCM{aesCipherAsm{aesCipher{l: uint8(len(key) + 28)}}}
// 	var rounds int
// 	switch len(key) {
// 	case 128 / 8:
// 		rounds = 10
// 	case 192 / 8:
// 		rounds = 12
// 	case 256 / 8:
// 		rounds = 14
// 	default:
// 		return nil, KeySizeError(len(key))
// 	}
//
// 	expandKeyAsm(rounds, &key[0], &c.enc[0], &c.dec[0])
// 	if supportsAES && supportsGFMUL {
// 		return &c, nil
// 	}
// 	return &c.aesCipherAsm, nil
// }
