package cmdcli

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

type versionCmd struct {
	out io.Writer
	version string
	buildTime string
	commitHash string
}

func (v *versionCmd) run() error {
	fmt.Fprintf(v.out, "helmvmgr %s\n", v.version)

	return nil
}

func newVersionCmd (out io.Writer, version string, buildTime string, commitHash string) *cobra.Command {
	v := &versionCmd{
		out: out,
		version: version,
		buildTime: buildTime,
		commitHash: commitHash,
	}
	cmd := &cobra.Command{
		Use: "version",
		Short: "short version",
		Long: "Long version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return v.run()
		},
	}

	return cmd
}