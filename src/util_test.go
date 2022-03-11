package main

import "testing"

func TestReturnSomething(t *testing.T) {
	if returnSomething("aaa") != "aaaa" {
		t.Errorf("An error occured.")
	}
}
