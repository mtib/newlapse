package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtib/newlapse/capture"
	"github.com/mtib/newlapse/convert"
	"github.com/mtib/newlapse/crop"
)

var (
	taskCrop        = flag.Bool("crop", false, "tells newlapse to crop")
	taskCapture     = flag.Bool("capture", false, "tells newlapse to capture")
	taskConvert     = flag.Bool("convert", false, "tells newlapse to convert %%ds folders to videos")
	taskCCC         = flag.Bool("ccc", false, "equals '-capture -crop -convert'")
	captureInterval = flag.Int("rate", 10, "seconds to wait between scrots")
	folder          = flag.String("folder", "./capture", "which folder to do something with")
	fps             = flag.Int("fps", 20, "ffmpeg framerate for videos")
	config          = flag.String("config", "nil", "config to read screensetup from for cropping")
)

func main() {
	flag.Parse()
	if *taskCCC {
		taskCrop = taskCCC
		taskCapture = taskCCC
		taskConvert = taskCCC
	}

	if *taskCapture {
		if _, err := os.Stat(*folder); err != nil {
			os.MkdirAll(*folder, os.ModePerm)
		}
		fmt.Println(`capturing into folder:`, *folder)
		capture.Folder(*folder, *captureInterval)
	}
	if *taskCrop {
		fmt.Println(`cropping folder:`, *folder)
		if *config != "nil" {
			crop.ConfigFolder(*folder, crop.ReadConfig(*config))
		} else {
			crop.Folder(*folder)
		}
	}
	if *taskConvert {
		convert.Folder(*fps)
	}

}
