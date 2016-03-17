package crop

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	// DisplayRegex is regex string
	DisplayRegex = "([0-9]{1,})x([0-9]{1,})\\+([0-9]{1,})\\+([0-9]{1,})"
)

var (
	// ScreenMatching is regex
	ScreenMatching = regexp.MustCompile(DisplayRegex)
)

// Screens is array of screen setups
type Screens []Screen

// Screen is single screens setup
type Screen struct {
	Width, Height, Padx, Pady int
}

// XrandrError is to be thrown when no information can be queried
type XrandrError string

func (x XrandrError) Error() string {
	return fmt.Sprintf("xrandr error: %s", string(x))
}

func (s Screen) String() string {
	return fmt.Sprintf("%dx%d+%d+%d", s.Width, s.Height, s.Padx, s.Pady)
}

func (s Screens) String() string {
	a := fmt.Sprintf("Screensetup with %d Screen", len(s))
	if len(s) > 1 {
		a += "s"
	}
	a += "\n"
	for k := range s {
		a += fmt.Sprintf("%d.: %v\n", k+1, s[k])
	}
	return a
}

// ScreenSetup returns current Screen Setup for all screens
func ScreenSetup() (Screens, error) {
	sa, err := regXrandr()
	if err != nil {
		return nil, err
	}
	screens := make(Screens, len(sa))
	for k, v := range sa {
		w, _ := strconv.ParseInt(v[1], 10, 32)
		h, _ := strconv.ParseInt(v[2], 10, 32)
		px, _ := strconv.ParseInt(v[3], 10, 32)
		py, _ := strconv.ParseInt(v[4], 10, 32)
		screens[k] = Screen{int(w), int(h), int(px), int(py)}
	}
	return screens, nil
}

func regXrandr() ([][]string, error) {
	scr, err := getXrandr()
	if err != nil {
		return nil, XrandrError("no xrandr")
	}
	return ScreenMatching.FindAllStringSubmatch(string(scr), -1), nil
}

func getXrandr() ([]byte, error) {
	c := exec.Command("xrandr")
	return c.Output()
}

// Folder does everything
func Folder(folder string) error {
	sc, err := ScreenSetup()
	if err != nil {
		panic(err)
	}
	return ImageForScreens(sc, folder)
}

// ImageForScreens crops folder and creates a new folder per screen
func ImageForScreens(s Screens, folder string) error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for k := range s {
		os.Mkdir(fmt.Sprintf("%ds", k+1), os.ModePerm)
	}
	var wg sync.WaitGroup
	fmt.Printf("start cropping (~%d files)\n", len(files))
	for _, f := range files {
		if n := f.Name(); strings.HasSuffix(n, ".png") || strings.HasSuffix(n, ".jpg") {
			wg.Add(len(s))
			for k, cs := range s {
				go func(cs Screen, folder string, n string, k int) {
					cropcall(cs, relfile(folder, n), fmt.Sprintf("%ds/%s", k+1, n))
					wg.Done()
				}(cs, folder, n, k)
			}
			wg.Wait()
		}
	}
	fmt.Println("completed cropping")
	return nil
}

func relfile(folder, name string) string {
	ds := regexp.MustCompile(`/.?/`)
	merge := folder + "/" + name
	for len(ds.FindAllString(merge, -1)) > 0 {
		merge = ds.ReplaceAllString(merge, "/")
	}
	return merge
}

func cropcall(croparea Screen, oldimage, newimage string) {
	c := exec.Command("convert", oldimage, "-crop", fmt.Sprintf("%s", croparea), "+repage", newimage)
	c.Start()
}
