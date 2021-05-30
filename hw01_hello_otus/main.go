package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.
	s := stringutil.Reverse("Hello, OTUS!")
	fmt.Print(s)
}
