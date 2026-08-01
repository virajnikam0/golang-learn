package o2

import (
	"matt_learning/02_hello_again/o3"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello Test"
	got := o3.SayHello("Test 1")

	if want != got {
		t.Errorf("Wanted %s got %s", want, got)
	}
}
