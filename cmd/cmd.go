// Copyright 2024 The stonebirdjx. All rights reserved.
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

// Package cmd provides a way to start a program with command line parameters
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bluebird",
	Short: "Bluebird is a very fast, efficient and comprehensive framework",
	Long: `Bluebird is a service with speed limit, authentication, logs, metric, and trace.
				  love by stonbird in Go.
				  Complete documentation is available at https://github.com/stonebirdjx/bluebird`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bluebird is a very fast, efficient and comprehensive framework")
	},
}

func init() {
	rootCmd.AddCommand(
		serverCmd,
		versionCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
