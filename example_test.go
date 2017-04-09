package humanize_test

import (
	"fmt"

	"github.com/olivere/humanize"
)

func ExampleSize() {
	fmt.Println(humanize.Size(2048))
	// Output:
	// 2 kB
}
