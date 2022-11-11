package config

import (
	"log"
	"runtime"
)

func DefaultPathAndMessage() (string, string) {
	baseMessage := "Base directory for World of Warcraft installation. -- default: "

	switch runtime.GOOS {
	case "windows":
		path := ""
		return path, baseMessage + path
	case "darwin":
		path := "/Applications/World of Warcraft"
		return path, baseMessage + path
	default:
		log.Fatalf("Runtime %s NYI", runtime.GOOS)
		return "", ""
	}
}
