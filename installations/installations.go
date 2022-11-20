package installations

import (
	"github-addon-backup/git"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Init(rootPath string) {
	err := wowDirectoryWalk(rootPath, git.InitIfMissing)
	if err != nil {
		log.Fatal(err)
	}
}

func KeepUpToDate(rootPath string) {
	err := wowDirectoryWalk(rootPath, git.CleanAndPullRepo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateDaily(rootWowDirPath string) {
	err := wowDirectoryWalk(rootWowDirPath, git.DailyPush)
	if err != nil {
		log.Fatal(err)
	}
}

func wowDirectoryWalk(rootPath string, action git.Action) error {
	rootFile, err := os.Open(rootPath)
	if err != nil {
		return err
	}

	dirs, err := rootFile.Readdirnames(0)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if strings.HasPrefix(dir, "_") {
			//Addons
			addonsPath := filepath.Join(rootPath, dir, "interface", "addons")
			if err := action(addonsPath, dir); err != nil {
				log.Println(err)
			}

			wtfAccountPath := filepath.Join(rootPath, dir, "WTF", "Account")
			wtfAccount, err := os.Open(wtfAccountPath)
			if err != nil {
				return err
			}

			accounts, err := wtfAccount.Readdirnames(0)
			if err != nil {
				return err
			}
			for _, account := range accounts {
				if account == "SavedVariables" {
					continue
				}

				accountPath := filepath.Join(wtfAccountPath, account)

				if err := action(accountPath, dir); err != nil {
					log.Println(err)
				}
			}
		}
	}
	return nil
}
