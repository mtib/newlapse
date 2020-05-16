package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mtib/newlapse/v2/capture"
	"github.com/mtib/newlapse/v2/convert"
	"github.com/mtib/newlapse/v2/xrandr"
)

var (
	capFlags  = flag.NewFlagSet("capture", flag.ExitOnError)
	convFlags = flag.NewFlagSet("convert", flag.ExitOnError)

	rate      = capFlags.Int("rate", 10, "Capture rate [in s]")
	all       = capFlags.Bool("all", false, "Use all recognized screens (overrides -screens)")
	screens   = capFlags.String("screens", "", "Comma separated list of display names (see xrandr)")
	videodevs = capFlags.String("cameras", "", "Comma separated list of video devices (e.g. /dev/video0)")

	fps = convFlags.Int("fps", 30, "Output fps")

	project = flag.String("project", "newlapse", "Project folder")

	captureScreens []string
	captureVideos  []string
	hasSettings    *bool = nil
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing subcommand: capture or convert")
		capFlags.PrintDefaults()
		convFlags.PrintDefaults()
		flag.PrintDefaults()
		os.Exit(1)
	}
	checkSettings()
	switch os.Args[1] {
	case "capture":
		capFlags.Parse(os.Args[2:])
		captureCli()
	case "convert":
		convFlags.Parse(os.Args[2:])
		convertCli()
	default:
		fmt.Println("Missing subcommand: capture or convert")
		capFlags.PrintDefaults()
		convFlags.PrintDefaults()
		flag.PrintDefaults()
	}
}

func captureCli() {
	if *all {
		onDisplays := xrandr.GetXrandr().On().Displays
		captureScreens = make([]string, 0)
		for _, k := range onDisplays {
			captureScreens = append(captureScreens, k.Name)
		}
	} else {
		captureScreens = []string{}
		if *screens != "" {
			captureScreens = strings.Split(*screens, ",")
		}
		for _, k := range captureScreens {
			if x, err := xrandr.GetDisplayFromName(k); err != nil || !x.Connected {
				log.Fatalf("Screen %s not connected", k)
			}
		}
	}
	for _, k := range captureScreens {
		err := os.MkdirAll(capturePath(k, "."), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	captureVideos = []string{}
	if *videodevs != "" {
		captureVideos = strings.Split(*videodevs, ",")
	}
	for _, k := range captureVideos {
		if _, err := os.Stat(k); os.IsNotExist(err) {
			log.Fatalf("Video device %s does not exist", k)
		}
		err := os.MkdirAll(capturePath(k, "."), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !*hasSettings {
		writeSettings()
	} else {
		settings, err := getSettings()
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
		if len(settings.UsedDisplays) != len(captureScreens) {
			log.Fatal("The number of screens doesn't match the project")
		}
		if len(settings.UsedVideoDevices) != len(captureVideos) {
			log.Fatal("The number of cameras doesn't match the project")
		}
	}
	captureLoop()
}

func capturePath(identifier string, timestamp string) string {
	cp := path.Join(*project, fmt.Sprintf("%x", sha256.Sum256([]byte(identifier)))[:12], timestamp)
	//log.Println(cp)
	return cp
}

func captureLoop() {
	t := time.NewTicker(time.Duration(*rate) * time.Second)
	for {
		select {
		case <-t.C:
			log.Println("Start capture")
			captureOnce()
			log.Println("Stop capture")
		}
	}
}

func captureOnce() {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	for _, screen := range captureScreens {
		xs, err := xrandr.GetDisplayFromName(screen)
		if err != nil {
			log.Printf("Screen %s has disappeared", screen)
			continue
		}
		capture.Scrot(capturePath(screen, timestamp), xs, capture.DefaultScrotSettings())
	}
	for _, devs := range captureVideos {
		err := capture.FsWebcamShot(devs, capturePath(devs, timestamp), capture.DefaultFsSettings())
		if err != nil {
			log.Println(err)
		}
	}
}

func convertCli() {
	list := convFlags.Args()
	folders := make([]string, 0)
	for _, l := range list {
		stat, err := os.Stat(l)
		if err != nil {
			log.Printf("Error while reading %s", l)
		}
		if stat.IsDir() {
			folders = append(folders, l)
		}
	}
	settings := convert.DefaultSettings()
	settings.Framerate = *fps
	for _, folder := range folders {
		convert.Render(folder, settings)
	}
}
