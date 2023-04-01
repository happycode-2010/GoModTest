package GoModTest

import (
	"fmt"
)

func PrintSome(s string) string {
	v := "Hello: " + s
	fmt.Println(v)
	return v
}
