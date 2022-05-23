package main

import (
	"fmt"

	v1 "github.com/jdmuntacb/gomodtestrapid/cmd/v1"
	v2 "github.com/jdmuntacb/gomodtestrapid/cmd/v2"
)

func main() {
	Hello()
	v1.Hello()
	v2.Hello()
}

func Hello() {
	fmt.Println("Rapid Hello - main")
}
