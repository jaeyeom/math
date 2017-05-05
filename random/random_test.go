package random

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestIntn(t *testing.T) {
	for i := 0; i < 100; i++ {
		rn := Intn(10)
		if 0 <= rn && rn < 10 {
			continue
		}
		t.Errorf("Intn should return in the range of [0, n)")
	}
}

func ExampleShuffle() {
	figures := sort.StringSlice{"rectangle", "hexagon", "square", "circle", "triangle", "pentagon"}
	deterministic := rand.New(rand.NewSource(1))
	Shuffle(figures.Len(), deterministic.Intn, figures.Swap)
	fmt.Println(figures)
	// Output:
	// [pentagon circle rectangle square hexagon triangle]
}
