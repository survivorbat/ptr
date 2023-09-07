package ptr

import "fmt"

func ExamplePtr() {
	boolPtr := Ptr(true)
	intPtr := Ptr(234)
	stringPtr := Ptr("abc")

	fmt.Println(boolPtr, intPtr, stringPtr)
}
