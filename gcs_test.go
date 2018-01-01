package gcsbucket_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/fishy/gcsbucket"
)

func Example() {
	ctx := context.Background()
	bkt := gcsbucket.Open("test-bucket")
	obj := "test/object"
	content := `Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

Excepteur sint occaecat cupidatat non proident,
sunt in culpa qui officia deserunt mollit anim id est laborum.`

	if err := bkt.Write(ctx, obj, strings.NewReader(content)); err != nil {
		// TODO: error handling
	}
	reader, err := bkt.Read(ctx, obj)
	if err != nil {
		// TODO: error handling
	}
	defer reader.Close()
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		// TODO: error handling
	}
	fmt.Println(string(buf)) // content
	if err := bkt.Delete(ctx, obj); err != nil {
		// TODO: error handling
	}
	fmt.Println(bkt.IsNotExist(bkt.Delete(ctx, obj))) // true
}

func TestEmpty(t *testing.T) {
	// Empty test to silence the no test warning of go test
}
