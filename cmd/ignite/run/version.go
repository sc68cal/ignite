package run

import (
	"encoding/json"
	"fmt"
	"github.com/luxas/ignite/pkg/version"
	"github.com/spf13/cobra"
	"io"
	"sigs.k8s.io/yaml"
)

// VersionData provides the version information of kubeadm.
type VersionData struct {
	Ignite      version.Info `json:"igniteVersion"`
	Firecracker version.Info `json:"firecrackerVersion"`
}

// Version provides the version information of kubeadm in format depending on arguments
// specified in cobra.Command.
func Version(out io.Writer, cmd *cobra.Command) error {
	v := VersionData{
		Ignite:      version.GetIgnite(),
		Firecracker: version.GetFirecracker(),
	}

	of, _ := cmd.Flags().GetString("output")
	switch of {
	case "":
		fmt.Fprintf(out, "ignite version: %#v\n", v.Ignite)
	case "short":
		fmt.Fprintf(out, "%s\n", v.Ignite.GitVersion)
	case "yaml":
		y, err := yaml.Marshal(&v)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	case "json":
		y, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	default:
		return fmt.Errorf("invalid output format: %s", of)
	}
	return nil
}