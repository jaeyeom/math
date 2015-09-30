package vector

import "fmt"

func ExampleNew() {
	fmt.Println(New(3))
	// Output: [0 0 0]
}

func ExampleVector_Len() {
	fmt.Println(New(2).Len())
	// Output: 2
}

func ExampleVector_Norm() {
	fmt.Printf("%.2f\n", Vector{3, 4}.Norm())
	// Output: 5.00
}

func ExampleDot() {
	fmt.Println(Dot(Vector{1, 2, 3}, Vector{0, 4, 5}))
	// Output: 23
}
