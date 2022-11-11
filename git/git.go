package git

import (
	"fmt"
	gitClient "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"time"
)

func CleanAndPullRepo(path string, branch string) error {
	c, err := gitClient.PlainOpen(path)
	if err != nil {
		if err == gitClient.ErrRepositoryNotExists {
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
		if err == gitClient.NoErrAlreadyUpToDate {
			log.Println("Up to date!")
		} else {
			return err
		}
	}

	return nil
}

func DailyPush(path string, branch string) error {
	c, err := gitClient.PlainOpen(path)
	if err != nil {
		if err == gitClient.ErrRepositoryNotExists {
			return nil
		}
		return err
	}
	log.Println(fmt.Sprintf("Pushing %s (%s)", path, branch))

	w, err := c.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&gitClient.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: true,
		Keep:   true,
	})
	if err != nil {
		return err
	}

	a, err := w.Status()
	if a.IsClean() {
		return nil
	}

	err = w.AddWithOptions(&gitClient.AddOptions{All: true})
	if err != nil {
		return err
	}

	hash, err := w.Commit(time.Now().Format("_1-_2-06"), &gitClient.CommitOptions{})
	if err != nil {
		return err
	}

	log.Println(hash)

	return nil
}
