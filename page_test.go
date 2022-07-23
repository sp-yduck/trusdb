package main

import (
	"reflect"
	"testing"
)

func TestPageIDIsValid(t *testing.T) {
	// valid page id
	pid := PageID(0)
	if !pid.IsValid() {
		t.Errorf("page id %d is valid, not invalid", pid)
	}
	pid = PageID(-2)
	if !pid.IsValid() {
		t.Errorf("page id %d is valid, not invalid", pid)
	}

	// invalid page id
	pid = PageID(-1)
	if pid.IsValid() {
		t.Errorf("page id %d is invalid, not valid", pid)
	}
	pid = InvalidPageID
	if pid.IsValid() {
		t.Errorf("page id %d is invalid, not valid", pid)
	}
}

func TestPageIDSerialize(t *testing.T) {
	pid := PageID(0)
	bytes := pid.Serialize()
	expBytes := []byte{0, 0, 0, 0}
	if !reflect.DeepEqual(expBytes, bytes) {
		t.Errorf("serialized %d should be %v not %v", pid, expBytes, bytes)
	}

	pid = PageID(1)
	bytes = pid.Serialize()
	expBytes = []byte{1, 0, 0, 0}
	if !reflect.DeepEqual(expBytes, bytes) {
		t.Errorf("serialized %d should be %v not %v", pid, expBytes, bytes)
	}

	pid = PageID(-1)
	bytes = pid.Serialize()
	expBytes = []byte{255, 255, 255, 255}
	if !reflect.DeepEqual(expBytes, bytes) {
		t.Errorf("serialized %d should be %v not %v", pid, expBytes, bytes)
	}

	pid = PageID(-2)
	bytes = pid.Serialize()
	expBytes = []byte{254, 255, 255, 255}
	if !reflect.DeepEqual(expBytes, bytes) {
		t.Errorf("serialized %d should be %v not %v", pid, expBytes, bytes)
	}
}
