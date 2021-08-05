package main

import "fmt"
import "rsc.io/quote"
func main() {
	/*
		Command: "go mod tidy"
		Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module.
	*/

	fmt.Println(quote.Go())
}