package hl7

import (
	"fmt"
)

func ExampleQuery_GetString() {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	msh9_1, _ := ParseQuery("MSH-9-1")
	msh9_2, _ := ParseQuery("MSH-9-2")

	fmt.Printf("%s_%s", msh9_1.GetString(m), msh9_2.GetString(m))
	// Output: ORU_R01
}
