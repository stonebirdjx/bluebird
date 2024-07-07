package model

import "fmt"

type Config struct {
	Port string  `yaml:"port"`
	Log  LogConf `yaml:"log"`
}

type LogConf struct {
	Dir        string `yaml:"dir"`
	Name       string `yaml:"name"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"max_size"`    // 一个文件最大可达多少M。
	MaxBackups int    `yaml:"max_backups"` // 最多同时保存多少个文件。
	MaxAge     int    `yaml:"max_age"`     // 一个文件最多可以保存多少天。
	Compress   bool   `yaml:"compress"`    // 用 gzip 压缩。
}

func (c *Config) Validate() error {
	if err := c.Log.Validate(); err != nil {
		return err
	}

	return nil
}

func (l *LogConf) Validate() error {
	if l.Dir == "" {
		return fmt.Errorf("log dir is empty")
	}

	if l.Level == "" {
		l.Level = "debug"
	}

	if l.Name == "" {
		l.Name = "bluebird.log"
	}

	if l.MaxSize == 0 {
		l.MaxSize = 100
	}

	if l.MaxBackups == 0 {
		l.MaxBackups = 5
	}

	if l.MaxAge == 0 {
		l.MaxAge = 7
	}

	return nil
}
