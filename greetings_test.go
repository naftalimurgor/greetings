package greetings

import (
	"testing"
)

func TestGreeting(t *testing.T)  {
	// belongs in the same package as the test and is called directly as it is exported(starts with an upper case letter)
	greeting, err := Hello("Ruth")
	
	expect := "Hi Ruth. Welcome!"

	if greeting!= expect && err == nil {
		t.Fatalf("got %q expected %q", greeting, expect)
	}
}