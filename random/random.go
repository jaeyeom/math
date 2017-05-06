package random

import (
	"crypto/rand"
	"math/big"
)

// Intn returns a random number in range [0, n). It uses crypto/rand package for
// better randomness.
func Intn(n int) int {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	return int(r.Int64())
}

// Shuffle shuffles, possibily a collection, using swap function. Parameter intn
// is a random number generator that returns integer in [0, size), such as
// rand.Intn or Intn in this package.
func Shuffle(size int, intn func(int) int, swap func(int, int)) {
	for i := 0; i < size; i++ {
		swap(i, i+intn(size-i))
	}
}

// Combination returns a random combination nCr, which is r-combinations from
// [0, n) range. 0 <= r <= n should hold.
func Combination(n, r int, intn func(int) int) map[int]struct{} {
	s := map[int]struct{}{}
	if r == n {
		for i := 0; i < n; i++ {
			s[i] = struct{}{}
		}
		return s
	}
	for i := n - r; i < n; i++ {
		x := intn(i)
		if _, picked := s[x]; !picked {
			s[x] = struct{}{}
		} else {
			s[i] = struct{}{}
		}
	}
	return s
}
