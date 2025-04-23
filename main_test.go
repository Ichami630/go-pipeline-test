package main

import "testing"

func TestSayHello(t *testing.T) {
	if got := sayHello(); got != "Hello, World!" {
		t.Errorf("sayHello() = %v, want %v", got, "Hello, World!")
	}
}
