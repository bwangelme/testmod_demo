package main

import (
	"fmt"
	"github.com/bwangelme/testmod/v2"
)

func main() {
	greet, err := testmod.Hi("xff", "cn")
	if err != nil {
		panic(err)
	}
	fmt.Println(greet)
}
