package cmdcli

import (
	"github.com/spf13/cobra"
	"helmvmgr/internal"
	"io"
)

var versionDesc = `Version:     %s
Commit Hash: %s
Build time:  %s
`
type versionCmd struct {
	out io.Writer
	render utils.RenderFunc
	Version string
	BuildTime string `yaml:"buildTime"`
	CommitHash string `yaml:"commitHash"`
	outputFormat string
}

var versionTemplate string = `Version:    {{ .Version }}
CommitHash: {{ .CommitHash }}
BuildTime:  {{ .BuildTime }}
`

func (v *versionCmd) run() error {
	return v.render(utils.RenderOutputParameters{
		Data:v,
		Out:v.out,
		Type:v.outputFormat,
		Template: versionTemplate,
	})
}

func newVersionCmd (render utils.RenderFunc, out io.Writer, version string, buildTime string, commitHash string) *cobra.Command {
	v := &versionCmd{
		out: out,
		render: render,
		Version: version,
		BuildTime: buildTime,
		CommitHash: commitHash,
		outputFormat: outputFormat,
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