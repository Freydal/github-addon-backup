package gh

import (
	"bytes"
	"github.com/cli/cli/v2/pkg/cmd/auth/login"
	"github.com/cli/cli/v2/pkg/cmd/auth/status"
	"github.com/cli/cli/v2/pkg/cmd/auth/token"
	"github.com/cli/cli/v2/pkg/cmd/factory"
	"github.com/cli/cli/v2/pkg/cmd/repo/create"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

func GetStatus() error {
	fct := factory.New("embed")

	cmd := status.NewCmdStatus(fct, nil)

	return cmd.RunE(nil, nil)
}

func Login() error {
	fct := factory.New("embed")

	cmd := login.NewCmdLogin(fct, nil)

	return cmd.RunE(&cobra.Command{}, nil)
}

func GetAuth() (string, error) {
	fct := factory.New("embed")

	r, w, _ := os.Pipe()
	fct.IOStreams.Out = w

	cmd := token.NewCmdToken(fct, nil)

	err := cmd.RunE(&cobra.Command{}, nil)
	if err != nil {
		return "", err
	}

	_ = w.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buf.String(), "\n"), nil
}

func RepoCreate(repoName string, absPath string) error {
	fct := factory.New("embed")

	cmd := create.NewCmdCreate(fct, nil)
	err := cmd.Flags().Set("public", "true")
	if err != nil {
		return err
	}

	err = cmd.Flags().Set("source", absPath)
	if err != nil {
		return err
	}

	return cmd.RunE(cmd, []string{repoName})
}
