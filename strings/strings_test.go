package main
import "testing"

func TestStringy(t *testing.T) {
	got := Stringy("apple", "app")
	want := true

	if want != true{
		t.Errorf("the test failed because I got %t and the output is %t", want, got)
	}
}