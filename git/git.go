package git

import (
	"fmt"
	"github-addon-backup/gh"
	"github.com/go-git/go-billy/v5/osfs"
	gitClient "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type Action = func(string, string) error

func InitIfMissing(path string, branch string) error {
	c, err := gitClient.PlainOpen(path)

	if c != nil && err != gitClient.ErrRepositoryNotExists {
		return nil
	}

	log.Println("Creating repo: %s (%s)", path, branch)

	wt := osfs.New(path)
	dotGit, err := wt.Chroot(".git")
	if err != nil {
		return err
	}
	fs := filesystem.NewStorage(dotGit, cache.NewObjectLRUDefault())

	c, err = gitClient.Init(fs, wt)
	if err != nil {
		return err
	}

	worktree, err := c.Worktree()
	if err != nil {
		return err
	}

	branchRef := plumbing.NewBranchReferenceName(branch)
	branchConfig := config.Branch{
		Name:   branch,
		Remote: "origin",
		Merge:  branchRef,
	}
	err = c.CreateBranch(&branchConfig)
	if err != nil {
		return err
	}

	err = c.Storer.SetReference(plumbing.NewSymbolicReference("HEAD", branchRef))
	if err != nil {
		return err
	}

	_, err = worktree.Add(".")
	if err != nil {
		return err
	}

	_, err = worktree.Commit(
		fmt.Sprintf("Initial commit: %s", time.Now().Format("1-_2-06")),
		&gitClient.CommitOptions{
			All: true,
		})
	if err != nil {
		return err
	}

	name := filepath.Base(path)
	name = strings.ReplaceAll(name, "#", "-hash-")

	err = gh.RepoCreate(name, path)
	if err != nil {
		return err
	}

	err = c.Push(&gitClient.PushOptions{})
	if err != nil {
		return err
	}

	return nil
}

func CleanAndPullRepo(path string, branch string) error {
	c, err := gitClient.PlainOpen(path)
	if err != nil {
		if err == gitClient.ErrRepositoryNotExists {
			return nil
		}
		return err
	}

	w, err := c.Worktree()
	if err != nil {
		return err
	}

	head, err := c.Head()
	if err != nil {
		return err
	}

	status, err := w.Status()
	if err != nil {
		return err
	}

	if !status.IsClean() {
		log.Println(fmt.Sprintf("Cleaning %s(%s)", path, head.Name()))
		err = w.Clean(&gitClient.CleanOptions{
			Dir: true,
		})
		if err != nil {
			log.Println("Failed to clean")
			return err
		}

		err = w.Checkout(&gitClient.CheckoutOptions{
			Branch: head.Name(),
			Force:  true,
		})
		if err != nil {
			log.Println("Failed to checkout to reset status")
			return err
		}
	}

	branchRefName := plumbing.NewBranchReferenceName(branch)

	if head.Name() != branchRefName {
		log.Println(fmt.Sprintf("Checking out branch %s(%s -> %s)", path, head.Name(), branchRefName))
		err = w.Checkout(&gitClient.CheckoutOptions{
			Branch: branchRefName,
			Force:  true,
		})

		if err != nil {
			if err == plumbing.ErrReferenceNotFound {
				err = c.Fetch(&gitClient.FetchOptions{})
				if err != nil && err != gitClient.NoErrAlreadyUpToDate {
					return err
				}

				cfg, err := c.Config()
				if err != nil {
					return err
				}

				configBranch, ok := cfg.Branches[branch]

				if !ok {
					configBranch = &config.Branch{
						Name:   branch,
						Remote: "origin",
						Merge:  branchRefName,
					}
					err = c.CreateBranch(configBranch)

					if err != nil {
						return err
					}
				}

				err = w.Checkout(&gitClient.CheckoutOptions{
					Branch: branchRefName,
					Create: true,
					Force:  true,
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	err = w.Pull(&gitClient.PullOptions{})
	if err != nil {
		if err == gitClient.NoErrAlreadyUpToDate {
			log.Println(fmt.Sprintf("Up to date! %s(%s)", path, head.Name()))
			return nil
		} else if err == gitClient.ErrNonFastForwardUpdate {

			reference, err := c.Storer.Reference(plumbing.NewRemoteReferenceName("origin", branch))
			if err != nil {
				return err
			}

			err = w.Reset(&gitClient.ResetOptions{
				Mode:   gitClient.HardReset,
				Commit: reference.Hash(),
			})
			if err != nil {
				return err
			}
			log.Println(fmt.Sprintf("Hard Reset to origin! %s(%s)", path, head.Name()))
			log.Println(fmt.Sprintf("Cleaning after reset %s(%s)", path, head.Name()))
			err = w.Clean(&gitClient.CleanOptions{
				Dir: true,
			})
			if err != nil {
				log.Println("Failed to clean")
				return err
			}
		} else {
			return err
		}
	}

	log.Println(fmt.Sprintf("Pulled %s(%s)", path, head.Name()))

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
