package cmdcli

import (
	"github.com/spf13/cobra"
	"helmvmgr/internal"
)
var (
	Version string
	BuildTime string
	CommitHash string
	outputFormat string
)

func NewCliCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "helmvmgr",
		Short:        "helm versioning manager.",
		Long:         "helm versioning long usage",

	}
	out := cmd.OutOrStdout()

	cf := cmd.PersistentFlags()
	cf.StringVarP(&outputFormat, "output", "o", "", "output type")


	cf.Parse(args)

	cmd.AddCommand(
		newVersionCmd(utils.RenderOutput, out, Version, BuildTime, CommitHash),
	)

	return cmd
}

