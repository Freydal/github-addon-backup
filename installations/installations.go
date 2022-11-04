package installations

import (
	"fmt"
	"github-addon-backup/git"
	"log"
	"os"
	"strings"
)

func KeepUpToDate(rootPath string) {
	rootFile, err := os.Open(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := rootFile.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		log.Println(dir)

		if strings.HasPrefix(dir, "__") {
			go git.CleanAndPullRepo(fmt.Sprintf("%s%s%s", rootPath, string(os.PathSeparator), dir), dir)
		}
	}

}
