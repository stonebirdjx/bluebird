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

	"gopkg.in/yaml.v3"
)

type config struct {
	Port string `yaml:"port"`
}

var cfg = &config{}

func InitConfig(f string) error {
	content, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, cfg)
}

// GetPort 返回cfg的port信息。
func GetPort() string {
	return cfg.Port
}
