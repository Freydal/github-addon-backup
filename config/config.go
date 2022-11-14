package config

import (
	"log"
	"runtime"
)

func DefaultPathAndMessage() string {

	switch runtime.GOOS {
	case "windows":
		path := "D:\\Program Files (x86)\\World of Warcraft"
		return path
	case "darwin":
		path := "/Applications/World of Warcraft"
		return path
	default:
		log.Fatalf("Runtime %s NYI", runtime.GOOS)
		return ""
	}
}
