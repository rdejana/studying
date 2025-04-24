package main

import "testing"

func Test1(t *testing.T) {
	path := "/a//b////c/d//././/.."
	expected := "/a/b/c"
	result := SimplifyPath(path)

	if result != expected {
		t.Errorf("result path %s is not expected path %s", result, expected)
	}
}

func Test2(t *testing.T) {
	path := "/../"
	expected := "/"
	result := SimplifyPath(path)

	if result != expected {
		t.Errorf("result path %s is not expected path %s", result, expected)
	}
}

func Test3(t *testing.T) {
	path := "/home//foo/"
	expected := "/home/foo"
	result := SimplifyPath(path)

	if result != expected {
		t.Errorf("result path %s is not expected path %s", result, expected)
	}
}
