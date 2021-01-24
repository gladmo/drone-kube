package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "v1.0.0"
)

var cmd = &cobra.Command{
	Use:   "drone-kube",
	Short: "drone-kube cli",
}

func init() {
	cmd.AddCommand(versionCmd, deliveryCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of drone-kube",
	Long:  `All software has versions. This is drone-kube's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("drone-kube Command", Version)
	},
}

// Execute cmd entrance
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
