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

package log

import (
	"os"
	"path"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/stonebirdjx/bluebird/biz/conf"
	"github.com/stonebirdjx/bluebird/pkg/consts"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() error {
	logInfo := conf.GetLogInfo()
	if err := os.MkdirAll(logInfo.Dir, 0777); err != nil {
		return err
	}

	fileName := path.Join(logInfo.Dir, logInfo.Name)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			return err
		}
	}

	logger := hertzzap.NewLogger(
		hertzzap.WithCoreEnc(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey: "msg",
				LevelKey:   "level",
				NameKey:    "name",
				TimeKey:    "ts",
				CallerKey:  "caller",
				//FunctionKey:    "func",
				StacktraceKey:  "stacktrace",
				LineEnding:     "\n",
				EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			})),
		hertzzap.WithZapOptions(
			zap.AddCaller(),
			zap.AddCallerSkip(3)),
		hertzzap.WithExtraKeys(
			[]hertzzap.ExtraKey{
				consts.LogIDExtraKey,
			}))
	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    logInfo.MaxSize,    // 一个文件最大可达多少M。
		MaxBackups: logInfo.MaxBackups, // 最多同时保存多少个文件。
		MaxAge:     logInfo.MaxAge,     // 一个文件最多可以保存多少天。
		Compress:   logInfo.Compress,   // 是否用 gzip 压缩。
	}

	logger.SetOutput(lumberjackLogger)
	switch logInfo.Level {
	case "debug":
		logger.SetLevel(hlog.LevelDebug)
	case "info":
		logger.SetLevel(hlog.LevelInfo)
	case "warn":
		logger.SetLevel(hlog.LevelWarn)
	case "error":
		logger.SetLevel(hlog.LevelError)
	case "fatal":
		logger.SetLevel(hlog.LevelFatal)
	}
	hlog.SetLogger(logger)

	return nil
}
