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
		t.Error("rendered path was ", path, " when it should be /user/Joe")
	}
}

func TestBadpathRender(t *testing.T) {
	t.Parallel()
	path, err := pathRender("/user/{name}", urlParams{"grall": "Joe"})
	if err == nil {
		t.Error("path render didn't returned an error when it should")
	}

	if path != "" {
		t.Error("rendered path was ", path, " when it should be ''")
	}
}

func ExamplepathRender() {
	path, _ := pathRender("/user/{name}/{id}", urlParams{"name": "Joe", "id": 17})
	fmt.Println(path)
	// Output:
	// /user/Joe/17
}
