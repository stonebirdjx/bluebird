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
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/cobra"
	"github.com/stonebirdjx/bluebird/biz/conf"
	"github.com/stonebirdjx/bluebird/biz/log"
	"github.com/stonebirdjx/bluebird/biz/router"
	"github.com/stonebirdjx/bluebird/biz/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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
	serverCmd.Flags().StringVarP(&scf.port, "port", "p", "6789", "Port to listen on")
}

func serverRun(cmd *cobra.Command, args []string) {
	// 系统初始化
	if err := serverInit(); err != nil {
		panic(err)
	}
	tracer :=  otel.Tracer("bluebird")
	_,span :=tracer.Start(context.Background(),"test-span")
	span.SetStatus(codes.Ok,"test ok")
	span.End()
	port := getPort()
	h := server.Default(
		server.WithHostPorts(port))

	router.CustomizedRegister(h)
	h.Spin()
}

func serverInit() error {
	if err := conf.Init(scf.configFile); err != nil {
		return err
	}
	if err := log.Init(); err != nil {
		return err
	}
	if err := trace.Init(); err != nil {
		return err
	}
	return nil
}

func getPort() string {
	port := conf.GetPort()
	if port != "" {
		return fmt.Sprint(":" + port)
	}
	return fmt.Sprint(":" + scf.port)
}
