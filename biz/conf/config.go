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

package conf

import (
	"os"

	"github.com/stonebirdjx/bluebird/pkg/model"
	"gopkg.in/yaml.v3"
)

var cfg = &model.Config{}

func Init(f string) error {
	content, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(content, cfg); err != nil {
		return nil
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	return nil
}

// GetPort 返回cfg的port信息。
func GetPort() string {
	return cfg.Port
}

func GetLogInfo() model.LogConf {
	return cfg.Log
}

func GetOtelInfo() model.OtelConf {
	return cfg.Otel
}
