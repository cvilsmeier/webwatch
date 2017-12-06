package main

import (
	"testing"
)

func assertNil(t *testing.T, act interface{}) {
	t.Helper()
	if act != nil {
		t.Fatalf("expected nil but was %v", act)
	}
}

func assertNotNil(t *testing.T, act interface{}) {
	t.Helper()
	if act == nil {
		t.Fatalf("expected not nil but was %v", act)
	}
}

func assertEqStr(t *testing.T, exp, act string) {
	t.Helper()
	if exp != act {
		t.Fatalf("expected %q but was %q", exp, act)
	}
}

func assertEqInt(t *testing.T, exp, act int) {
	t.Helper()
	if exp != act {
		t.Fatalf("expected %d but was %d", exp, act)
	}
}

func assertTrue(t *testing.T, cond bool) {
	t.Helper()
	if !cond {
		t.Fatalf("expected true but was false")
	}
}
