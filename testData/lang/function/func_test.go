package function_test

import (
	"fmt"

	"github.com/peczenyj/gomarkdoc/testData/lang/function"
)

func ExampleStandalone() {
	res, _ := function.Standalone(2, "abc")
	fmt.Println(res)
	// Output: 2
}

func ExampleStandalone_zero() {
	res, _ := function.Standalone(0, "def")
	fmt.Println(res)
	// Output: 0
}

func ExampleReceiver() {
	r := &function.Receiver{}
	fmt.Println(r)
}

func ExampleReceiver_subTest() {
	var r function.Receiver
	r.WithReceiver()
}

func ExampleGeneric_WithGenericReceiver() {
	r := function.Generic[int]{}
	r.WithGenericReceiver()
}
