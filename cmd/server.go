// Copyright 2024 The stonebirdjx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/cobra"
	"github.com/stonebirdjx/bluebird/biz/router"
)

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
	serverCmd.Flags().StringVarP(&scf.port, "port", "p", ":6789", "Port to listen on")
}

func serverRun(cmd *cobra.Command, args []string) {
	h := server.Default(
		server.WithHostPorts(scf.port),
	)

	router.CustomizedRegister(h)
	h.Spin()
}
