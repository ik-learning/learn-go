package reverse

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestReverseToReturnReverseInputString(t *testing.T) {
	actualResult := reverse("Hello")

	assert.Equal(t, actualResult, "olleH", "they should be equal")
}