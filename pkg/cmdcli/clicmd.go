package cmdcli

import (
	"fmt"
	"github.com/spf13/cobra"
)
var (
	Version string
	BuildTime string
	CommitHash string
)

func NewCliCmd(args []string) *cobra.Command {
	fmt.Printf("v3: %s\n", Version)
	cmd := &cobra.Command{
		Use:          "helmvmgr",
		Short:        "helm versioning manager.",
		Long:         "helm versioning long usage",

	}
	out := cmd.OutOrStdout()

	cmd.AddCommand(
		newVersionCmd(out, Version, BuildTime, CommitHash),
	)
	return cmd
}

