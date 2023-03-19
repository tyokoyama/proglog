package log

import (
	"os"
	"testing"

	api "github.com/travisjeffery/proglog/api/v1"
)

func TestSegment(t *testing.T) {
	dir, _ := os.MkdirTemp("", "segment-test")
	defer os.RemoveAll(dir)

	want := &api.Record{Value: []byte("hello world")}

}