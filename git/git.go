package git

import (
	gitClient "github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing"
	"log"
)

func CleanAndPullRepo(path string, installation string) error {
	c, err := gitClient.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	w, err := c.Worktree()
	if err != nil {
		log.Fatal(err)
	}

    err = w.Clean(&gitClient.CleanOptions{})
    if err != nil {
		log.Fatal(err)
	}

	err = w.Checkout(&gitClient.CheckoutOptions{
		Branch: plumbing.ReferenceName(installation),
	})

	err = w.Pull(&gitClient.PullOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
