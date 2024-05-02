package cmd

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Example: `  bluebird server `,
	Run:     serverRun,
}

// serverFlags
type serverCmdFlags struct {
	configFile string
	port       string
}

var scf serverCmdFlags

func init() {
	serverCmd.Flags().StringVarP(&scf.configFile, "config", "c", "config.yaml", "Config file to use")
	serverCmd.Flags().StringVarP(&scf.port, "port", "p", "6789", "Port to listen on")
}

func serverRun(cmd *cobra.Command, args []string) {
	// TODO: Add your server logic here

}
