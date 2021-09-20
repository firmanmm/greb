package simple

import (
	"fmt"

	"github.com/firmanmm/greb"
)

func ExampleSimple() {
	result, err := greb.Generate("simple.greb")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	/* Output: Ex
	 */
}
