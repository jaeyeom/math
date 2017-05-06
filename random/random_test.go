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
	// For test purpose. For better randomness, you may use random.Intn.
	deterministic := rand.New(rand.NewSource(1))
	Shuffle(figures.Len(), deterministic.Intn, figures.Swap)
	fmt.Println(figures)
	// Output:
	// [pentagon circle rectangle square hexagon triangle]
}

func ExampleShuffle_closure() {
	figures := []string{"rectangle", "hexagon", "square", "circle", "triangle", "pentagon"}
	// For test purpose. For better randomness, you may use random.Intn.
	deterministic := rand.New(rand.NewSource(1))
	Shuffle(len(figures), deterministic.Intn, func(i, j int) {
		figures[i], figures[j] = figures[j], figures[i]
	})
	fmt.Println(figures)
	// Output:
	// [pentagon circle rectangle square hexagon triangle]

}

func TestCombination(t *testing.T) {
	for n := 1; n <= 10; n++ {
		for r := 0; r <= n; r++ {
			c := Combination(n, r, Intn)
			if len(c) != r {
				t.Errorf("size want %d actual %d", r, len(c))
			}
			for e := range c {
				if e < 0 || e >= n {
					t.Errorf("%d should be in [0, %d)", e, n)
				}
			}
		}
	}
}

func ExampleCombination() {
	const r = 3
	deterministic := rand.New(rand.NewSource(1))
	c := Combination(10, r, deterministic.Intn)
	l := make([]int, 0, r)
	for n := range c {
		l = append(l, n)
	}
	sort.Ints(l)
	fmt.Println(l)
	// Output: [2 6 7]
}
