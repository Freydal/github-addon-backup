package installations

import (
	"github-addon-backup/git"
	"log"
	"os"
	"path/filepath"
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
		if strings.HasPrefix(dir, "_") {
			//Addons
			addonsPath := filepath.Join(rootPath, dir, "interface", "addons")
			if err := git.CleanAndPullRepo(addonsPath, dir); err != nil {
				log.Println(err)
			}

			wtfAccountPath := filepath.Join(rootPath, dir, "WTF", "Account")
			wtfAccount, err := os.Open(wtfAccountPath)
			if err != nil {
				log.Fatal(err)
			}

			accounts, err := wtfAccount.Readdirnames(0)
			if err != nil {
				log.Fatal(err)
			}
			for _, account := range accounts {
				accountPath := filepath.Join(wtfAccountPath, account)

				if err := git.CleanAndPullRepo(accountPath, dir); err != nil {
					log.Println(err)
				}
			}
		}
	}

}
