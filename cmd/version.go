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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stonebirdjx/bluebird/pkg/consts"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the client and server version information for the current context.",
	Example: `  bluebird version`, // 保留两个空格
	Run:     versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println(consts.Version)
	os.Exit(consts.ExitCodeSuccess)
}
