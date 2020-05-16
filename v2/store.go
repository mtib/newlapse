package main

import (
	"encoding/json"
	"os"
	"path"
)

type (
	ProjectJSON struct {
		UsedDisplays     []string `json:"displays"`
		UsedVideoDevices []string `json:"videos"`
		Rate             int      `json:"rate"`
	}
)

func settingsPath() string {
	return path.Join(*project, "newlapse.project")
}

func checkSettings() {
	_, err := os.Stat(settingsPath())
	val := err == nil
	hasSettings = &(val)
}

func writeSettings() {
	f, _ := os.Create(settingsPath())
	json.NewEncoder(f).Encode(ProjectJSON{
		UsedDisplays:     captureScreens,
		UsedVideoDevices: captureVideos,
		Rate:             *rate,
	})
}

func getSettings() (ProjectJSON, error) {
	f, err := os.Open(settingsPath())
	if err != nil {
		return ProjectJSON{}, err
	}
	var j ProjectJSON
	json.NewDecoder(f).Decode(&j)
	return j, nil
}
