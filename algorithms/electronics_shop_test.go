package algs

import (
	"fmt"
	"testing"

	// "github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/assert"

	. "./utils"
)

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	change := b
	canBuy := false
	for k := 0; k < len(keyboards); k++ {
		for u := 0; u < len(drives); u++ {
			left := b - (keyboards[k] + drives[u])
			if left == 0 {
				return b
			} else if left > 0 && left < change {
				change = left
				canBuy = true
			}
		}
	}
	if canBuy {
		return b - change
	} else {
		return int32(-1)
	}
}

func TestShouldSpendMoney(t *testing.T) {
	values := ReadFixturesIntV3("./utils/fixtures/electronics-shop.json")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i].Input)
		input := values[i].Input
		result := getMoneySpent(input.Keyboard, input.Usb, input.Budget)
		fmt.Println(result)
		assert.Equal(t, result, values[i].Result, "they should be equal")
	}
}
