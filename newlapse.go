package main

import (
	"flag"
	"fmt"
	"github.com/mtib/newlapse/capture"
	"github.com/mtib/newlapse/convert"
	"github.com/mtib/newlapse/crop"
	"os"
)

var (
	taskCrop        = flag.Bool("crop", false, "tells newlapse to crop")
	taskCapture     = flag.Bool("capture", false, "tells newlapse to capture")
	taskConvert     = flag.Bool("convert", false, "tells newlapse to convert %%ds folders to videos")
	taskCCC         = flag.Bool("ccc", false, "equals '-capture -crop -convert'")
	captureInterval = flag.Int("rate", 10, "seconds to wait between scrots")
	folder          = flag.String("folder", "./capture", "which folder to do something with")
	fps             = flag.Int("fps", 20, "ffmpeg framerate for videos")
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
		crop.Folder(*folder)
	}
	if *taskConvert {
		convert.Folder(*fps)
	}

}
