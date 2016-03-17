package crop

import (
	"fmt"
	"testing"
)

func TestXrandr(t *testing.T) {
	s, e := getXrandr()
	if e != nil {
		t.FailNow()
	}
	fmt.Println(string(s))
}

func TestRegXrand(t *testing.T) {
	sa, e := regXrandr()
	if e != nil {
		t.FailNow()
	}
	fmt.Printf("%q\n", sa)
}

func TestScreenSetup(t *testing.T) {
	ss, e := ScreenSetup()
	if e != nil {
		t.FailNow()
	}
	fmt.Println(ss)
}
