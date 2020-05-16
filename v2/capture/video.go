package capture

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type FsSettings struct {
	Quality   int
	Width     int
	Height    int
	Extension string
}

func FsWebcamShot(device string, path string, settings FsSettings) error {
	file := fmt.Sprintf("%s.%s", path, settings.Extension)
	cmd := exec.Command("fswebcam",
		"-d", device,
		"-r", fmt.Sprintf("%dx%d", settings.Width, settings.Height),
		"--jpeg", strconv.Itoa(settings.Quality),
		file)
	err := cmd.Run()
	if err != nil {
		return err
	}
	if _, err = os.Stat(file); os.IsNotExist(err) {
		return errors.New("File expected to be created not found")
	}
	return nil
}

func DefaultFsSettings() FsSettings {
	return FsSettings{
		Quality:   80,
		Width:     1280,
		Height:    720,
		Extension: "jpg",
	}
}
