package capture

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mtib/newlapse/v2/xrandr"
)

type ScrotSettings struct {
	Quality   int
	Extension string
}

func Scrot(basename string, display xrandr.XrandrDisplay, settings ScrotSettings) {
	cmd := exec.Command("scrot",
		"-a", display.XrandrLine(),
		"-q", fmt.Sprintf("%d", settings.Quality),
		fmt.Sprintf("%s.%s", basename, settings.Extension))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func DefaultScrotSettings() ScrotSettings {
	return ScrotSettings{
		Quality:   80,
		Extension: "jpg",
	}
}
