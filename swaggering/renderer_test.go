package swaggering

import "fmt"

import _ "testing"

func ExampleSnakeCase() {
	fmt.Println(snakeCase("HTTPtest"))
	fmt.Println(snakeCase("SomethingJSON"))
	// Output:
	// httptest
	// something_json
}
