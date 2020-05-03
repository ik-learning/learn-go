package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func climbingLeaderboard2(scores []int32, alice []int32) []int32 {
	v1 := map[int32]int32{}
	or := make([]int32,0)
	for i:=0;i<len(scores);i++{
		if _,e:= v1[scores[i]];e {
			v1[scores[i]] += 1
		} else {
			v1[scores[i]]=1
			or = append(or, scores[i])
		}
	}

	rank := make([]int32,len(alice))
	zi := 0
	for i:=len(alice)-1;i>=0;i--{
		for z:=zi;z<len(or);{
			if alice[i] >= or[z]  {
				zi = z
				break
			} else if z == len(or) - 1 {
				zi = z + 1
			    break
			} else {
				z++
			}
		}
		rank[i] = int32(zi)+1
	}
	fmt.Println(alice)
	return rank
}

func TestClimbingLeaderboard(t *testing.T) {
	fixtures := [...]struct {
		scores   []int32
		alice    []int32
		expected []int32
	}{
		{[]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5,25,50,120}, []int32{6,4,2,1}},
		{[]int32{100, 90,90,80,75,60}, []int32{50,65,77,90,102}, []int32{6,5,4,2,1}},
	}
	for _, e := range fixtures {
		result := climbingLeaderboard2(e.scores, e.alice)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}

