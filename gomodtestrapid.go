package gomodtestrapid

import (
	"fmt"
)

func main() {
	Hello()
}

func Hello() {
	fmt.Println("Rapid Hello - v2")
}

func SuperHello(name string) {
	fmt.Println("Rapid Super Hello " + name)
}
