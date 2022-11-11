package git

import (
	gitClient "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"path/filepath"
)

func CleanAndPullRepo(path string, installation string) error {
	log.Println("Clean and pull", path, installation)

	addonsPath := filepath.Join(path, "interface", "addons")
	//Addons
	c, err := gitClient.PlainOpen(addonsPath)
	if err != nil {
		log.Fatal("Plain Open Failed", err)
	}

	log.Println("Opened", addonsPath)

	w, err := c.Worktree()
	if err != nil {
		log.Fatal("Worktree failed", err)
	}

	err = w.Clean(&gitClient.CleanOptions{})
	if err != nil {
		log.Fatal("Clean failed", err)
	}

	err = w.Checkout(&gitClient.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(installation),
	})
	if err != nil {
		log.Fatal("Checkout failed", err)
	}

	//sshAuth, err := gitClientSSH.DefaultAuthBuilder("")
	//log.Println(sshAuth.String())
	//log.Println(sshAuth.Name())
	//
	//fffaa, err := sshAuth.ClientConfig()

	err = w.Pull(&gitClient.PullOptions{})
	if err != nil {
		log.Fatal("Pull failed", err)
	}

	return nil
}
