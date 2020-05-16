package convert

import (
	"fmt"
	"os/exec"
	"path"
	"strconv"
)

type Settings struct {
	CRF         int
	Venc        string
	InFiletype  string
	OutFiletype string
	Framerate   int
}

func Render(folder string, settings Settings) error {
	cmd := exec.Command("ffmpeg",
		"-y",
		"-pattern_type", "glob",
		"-i", fmt.Sprintf("%s/*.%s", folder, settings.InFiletype),
		"-r", strconv.Itoa(settings.Framerate),
		"-crf", strconv.Itoa(settings.CRF),
		"-vcodec", settings.Venc,
		fmt.Sprintf("%s.%s", path.Base(folder), settings.OutFiletype))
	return cmd.Run()
}

func DefaultSettings() Settings {
	return Settings{
		CRF:         30,
		Venc:        "libx265",
		InFiletype:  "jpg",
		OutFiletype: "mp4",
		Framerate:   30,
	}
}
