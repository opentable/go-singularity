package client

import (
	"fmt"
	"testing"
)

func TestGoodpathRender(t *testing.T) {
	t.Parallel()
	path, err := pathRender("/user/{name}", urlParams{"name": "Joe"})
	if err != nil {
		t.Error("path render returned an error when it shouldn't")
	}

	if path != "/user/Joe" {
		t.Error("rendered path was %v when it should be /user/Joe")
	}
}

func TestBadpathRender(t *testing.T) {
	t.Parallel()
	path, err := pathRender("/user/{name}", urlParams{"grall": "Joe"})
	if err == nil {
		t.Error("path render didn't returned an error when it should")
	}

	if path == nil {
		t.Error("rendered path was %v when it should be <nil>")
	}
}
func ExamplepathRender() {
	path, err := pathRender("/user/{name}", urlParams{"name": "Joe"})
	fmt.Println(path)
	// Output: /user/Joe
}
