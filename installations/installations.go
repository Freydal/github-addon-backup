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

func UpdateDaily(rootWowDirPath string) {
	rootWowDirFile, err := os.Open(rootWowDirPath)
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := rootWowDirFile.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		if strings.HasPrefix(dir, "_") {
			//Addons
			addonsPath := filepath.Join(rootWowDirPath, dir, "interface", "addons")
			if err := git.DailyPush(addonsPath, dir); err != nil {
				log.Println(err)
			}

			wtfAccountPath := filepath.Join(rootWowDirPath, dir, "WTF", "Account")
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
				if err := git.DailyPush(accountPath, dir); err != nil {
					log.Println(err)
				}
			}
		}
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
