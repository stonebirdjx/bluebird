package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v1.0.0"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the client and server version information for the current context.",
	Example: `  bluebird version`, // 保留两个空格
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
