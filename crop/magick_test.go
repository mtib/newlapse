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

func TestRelfile(t *testing.T) {
	type combi struct{ folder, file string }
	data := map[string]combi{
		"folder/file":  combi{"folder", "file"},
		"folder/file2": combi{"folder/", "file2"},
		"folder/file3": combi{"folder", "/file3"},
		"folder/file4": combi{"folder/", "/file4"},
		"123/456":      combi{"123", "456"},
	}
	for k, v := range data {
		ans := relfile(v.folder, v.file)
		if ans != k {
			fmt.Print(" --> ")
			t.Fail()
		}
		fmt.Printf("%s + %s --> %s (%s)\n", v.folder, v.file, ans, k)
	}
}

func TestFolder(t *testing.T) {
	ImageForScreens(nil, "./")
}

func TestXrandError(t *testing.T) {
	fmt.Println(XrandrError("Testing Error!"))
}
