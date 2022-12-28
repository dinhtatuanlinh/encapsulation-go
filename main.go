package main

import (
	"encapsulation-go/encapsulation"
	"fmt"
)

func main() {
	validation := encapsulation.NewValidation()
	fmt.Println(validation.GetPattern())
}
