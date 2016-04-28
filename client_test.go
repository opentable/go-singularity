package singularity

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

func ExampleComplexPathRender() {
	path, _ := pathRender("/api/history/request/{requestId}/deploy/{deployId}", urlParams{"requestId": "192.168.99.100:5000hellolabels:latest", "deployId": "4990399829e84ce2be5a523b252e1afd"})
	fmt.Println(path)
	// Output:
	// /api/history/request/192.168.99.100:5000hellolabels:latest/deploy/4990399829e84ce2be5a523b252e1afd
}
