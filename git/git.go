package git

import (
	"fmt"
	gitClient "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
)

func CleanAndPullRepo(path string, branch string) error {
	c, err := gitClient.PlainOpen(path)
	if err != nil {
		if err.Error() == "repository does not exist" {
			return nil
		}
		return err
	}

	log.Println(fmt.Sprintf("Clean and pull %s (%s)", path, branch))

	w, err := c.Worktree()
	if err != nil {
		return err
	}

	err = w.Clean(&gitClient.CleanOptions{})
	if err != nil {
		return err
	}

	err = w.Checkout(&gitClient.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return err
	}

	err = w.Pull(&gitClient.PullOptions{})
	if err != nil {
		if err.Error() == "already up-to-date" {
			log.Println("Up to date!")
		} else {
			return err
		}
	}

	return nil
}
