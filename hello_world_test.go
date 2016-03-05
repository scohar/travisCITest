package main

import "testing"

func TestReturnHelloWorld(t *testing.T) {
    hwString := ReturnHelloWorld()
    expected := "hello world"
    if hwString != expected {
        t.Errorf("ReturnHelloWorld returned %q, expected $q", hwString, expected)
    }
}
