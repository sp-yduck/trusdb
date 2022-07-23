package main

import (
	"os"
	"reflect"
	"testing"
)

func TestFileReadWritePage(t *testing.T) {
	fm := NewFileManager("newfile")
	defer os.Remove("newfile")
	defer fm.ShutDown()

	data := make([]byte, PageSize)
	buffer := make([]byte, PageSize)
	copy(data, "this is a test string")
	fm.WritePage(PageID(0), data)

	fm.ReadPage(PageID(0), buffer)

	if !reflect.DeepEqual(data, buffer) {
		t.Errorf("expected page is %v, not %v", string(data), string(buffer))
	}
}
