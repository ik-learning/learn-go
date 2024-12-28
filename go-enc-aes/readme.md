# Cypher Golang

- [Cypher Blog](https://bitfieldconsulting.com/posts/aes-encryptio)
- [Aes Library](https://pkg.go.dev/crypto/aes)
- [Go Encrypt and Decrypt](https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/)
- [Sources](https://github.com/golang/go/blob/master/src/crypto/aes/aes_test.go)

In ASCII, each character is represented by a single byte. Thus, a single ASCII character can be stored in Go as a byte data type.
In contrast, Unicode characters outside the simple ASCII set are comprised of multiple bytes up to a maximum of 4 bytes.

f byte is an 8 bit integer, does it mean we can represent any raw data using just 8 bit integer numbers?

```go
package aes
// https://github.com/golang/go/blob/master/src/crypto/aes/aes.go
// NewCipher creates and returns a new [cipher.Block].
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewCipher(key []byte) (cipher.Block, error) {
	k := len(key)
	switch k {
	default:
		return nil, KeySizeError(k)
	case 16, 24, 32:
		break
	}
	if boring.Enabled {
		return boring.NewAESCipher(key)
	}
	return newCipher(key)
}
```