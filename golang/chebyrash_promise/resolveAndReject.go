package main

import (
	"errors"
	"fmt"

	"github.com/chebyrash/promise"
)

func main() {

	var p1 = promise.Resolve("Hello, World")
	result, _ := p1.Await()
	fmt.Println(result)

	var p2 = promise.Reject(errors.New("bad error"))
	_, err := p2.Await()
	fmt.Println(err)
}
