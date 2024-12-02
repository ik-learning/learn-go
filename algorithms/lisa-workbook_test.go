package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func workbook(k int32, arr []int32) int32 {
	msg := false
	page := int32(0)
	special := int32(0)
	chapter := 0
	for i := 0; i < len(arr); i++ {
		chapter++
		var fullPages int32
		startPage := page + 1
		if arr[i] > k {
			if arr[i]%k > 0 {
				fullPages = arr[i]/k + 1
			} else {
				fullPages = arr[i] / k
			}
			issueStart := 1
			for e := startPage; e < startPage+fullPages+1; e++ {
				issueEnd := issueStart + int(k) - 1
				if issueEnd > int(arr[i]) {
					issueEnd = int(arr[i])
				}
				if e >= int32(issueStart) && e <= int32(issueEnd) {
					special += 1
				}
				issueStart = issueEnd + 1
			}
		} else {
			fullPages = 1
			currentPage := page + fullPages
			if currentPage <= arr[i] {
				special += 1
			}
		}
		page += fullPages
		if msg {
			fmt.Printf("Chapter > %d. Pages > %d. Start Page > %d. End Page > %d.\n", chapter, fullPages, startPage, page)
		}
	}
	return special
}

func TestWorkbook(t *testing.T) {
	fixtures := [...]struct {
		max      int32
		problems []int32
		expected int32
	}{
		{3, []int32{4, 2, 6, 1, 10}, 4},
		{5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}, 8},
		{5, []int32{1, 0, 6}, 2},
		{5, []int32{1, 0, 6, 1, 6}, 2},
	}
	for _, e := range fixtures {
		result := workbook(e.max, e.problems)
		Equal(t, e.expected, result, "they should be equal")
	}
}

// fmt.Printf("Page: %d. Issue Start: %d. Page End %d\n", startPage, issueStart, startPage+fullPages)
// if msg {
// fmt.Printf("Chapter %d, Found on Page: %d. Issues %d. \n", chapter, currentPage, arr[i])
// }
