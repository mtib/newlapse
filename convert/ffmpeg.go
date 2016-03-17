package convert

import (
	"fmt"
	"os/exec"
	"regexp"
)

// Folder does it all
func Folder(fr int) {
	r := regexp.MustCompile("[1-9]{1,}[0-9]*s")
	odir, err := exec.Command("ls", "-l").Output()
	if err != nil {
		panic(err)
	}
	dirs := r.FindAllString(string(odir), -1)
	var converts chan *exec.Cmd
	converts = make(chan *exec.Cmd, len(dirs))
	for k, d := range dirs {
		fmt.Printf("start ffmpeg conversion #%d\n", k+1)
		converts <- ffmpegFolder(d, fr)
	}
	for {
		select {
		case c := <-converts:
			c.Wait()
		default:
			fmt.Println("completed conversion")
			return
		}
	}
}

func ffmpegFolder(folder string, framerate int) *exec.Cmd {
	//"ffmpeg -y -r {} -pattern_type glob -i '{}' -c:v libx264 {}"
	// cmd := exec.Command("/bin/sh", "-c", "ffmpeg", "-y", fmt.Sprintf("-r %d", framerate), "-pattern_type glob", fmt.Sprintf("-i '%s/%s'", folder, "*.png"), fmt.Sprintf("'%s_video.mp4'", folder))
	cmd := exec.Command("ffmpeg", "-y", "-r", fmt.Sprintf("%d", framerate), "-pattern_type", "glob", "-i", fmt.Sprintf("%s/*.png", folder), fmt.Sprintf("video_%c.mp4", []rune(folder)[0]))
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	return cmd
}

func ffmpegcall(folder string, fps int) string {
	return fmt.Sprintf("--y -r %d -pattern_type glob -i %s/*.png %s_video.mkv", fps, folder, folder)
}
