package xrandr

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	displayRegex = regexp.MustCompile("(\\S+) ((dis|)connected)(.*?(\\d+)x(\\d+)\\+(\\d+)\\+(\\d+)|)")
)

type (
	XrandrQuery struct {
		Displays []XrandrDisplay
	}
	XrandrDisplay struct {
		Name           string
		CurrentMode    Mode
		SupportedModes []Mode
		Connected      bool
	}
	Dimension struct {
		Width, Height int
	}
	Mode struct {
		Dimension Dimension
		Offset    Dimension
		Rate      float32
	}
)

func GetXrandrText() string {
	cmd := exec.Command("xrandr")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func GetXrandr() XrandrQuery {
	var xrq XrandrQuery

	text := GetXrandrText()
	displays := displayRegex.FindAllStringSubmatch(text, -1)

	xrq.Displays = make([]XrandrDisplay, len(displays))

	for i, match := range displays {
		if match[2] == "disconnected" {
			xrq.Displays[i] = XrandrDisplay{
				Name:      match[1],
				Connected: false,
			}
			continue
		}

		width, _ := strconv.Atoi(match[5])
		height, _ := strconv.Atoi(match[6])
		offsetx, _ := strconv.Atoi(match[7])
		offsety, _ := strconv.Atoi(match[8])

		xrq.Displays[i] = XrandrDisplay{
			Name:      match[1],
			Connected: true,
			CurrentMode: Mode{
				Dimension: Dimension{
					Width:  width,
					Height: height,
				},
				Offset: Dimension{
					Width:  offsetx,
					Height: offsety,
				},
			},
		}
	}

	return xrq
}

// GetDisplayFromName returns the display struct for the display with same name,
// if a error is returned (display not found) the display struct contains garbage.
// Note that a disconnected display will not return an error.
func GetDisplayFromName(name string) (XrandrDisplay, error) {
	xr := GetXrandr() // TODO cahce GetXrandr for GetDisplayFromName
	for _, k := range xr.Displays {
		if k.Name == name {
			return k, nil
		}
	}
	return XrandrDisplay{}, errors.New("Display not found")
}

func (x XrandrQuery) On() XrandrQuery {
	var xrq XrandrQuery
	xrq.Displays = make([]XrandrDisplay, 0)
	for _, k := range x.Displays {
		if k.Connected {
			xrq.Displays = append(xrq.Displays, k)
		}
	}
	return xrq
}

func (d XrandrDisplay) XrandrLine() string {
	dim := d.CurrentMode.Dimension
	off := d.CurrentMode.Offset
	return fmt.Sprintf("%d,%d,%d,%d", off.Width, off.Height, dim.Width, dim.Height)
}
