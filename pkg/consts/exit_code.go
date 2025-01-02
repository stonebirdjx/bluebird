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

package consts

// ExitCode
const (
	ExitCodeSuccess                     = 0   // 成功
	ExitCodeGeneralError                = 1   // 通用错误（General error），一般性错误，通常表示命令执行失败。
	ExitCodeMisuseOfShellBuiltins       = 2   // 误用命令（Misuse of shell builtins），命令的用法错误，例如，缺少必要的参数。
	ExitCodeCommandInvokedCannotExecute = 126 // 命令不可执行（Command invoked cannot execute）， 尝试运行一个不可执行的命令，例如权限不足。
	ExitCodeCommandNotFound             = 127 // 命令未找到（Command not found），尝试运行一个不存在的命令或脚本。
	ExitCodeInvalidExitArgument         = 128 // 无效的退出参数（Invalid exit argument），使用了无效的退出状态码。
	ExitCodeScriptTerminatedByControlC  = 130 // 进程被 Ctrl+C 终止（Script terminated by Control-C），进程接收到 SIGINT 信号（中断信号）。
	ExitCodeScriptTerminatedBySIGKILL   = 137 // 进程被 kill 命令终止（Script terminated by SIGKILL），进程接收到 SIGKILL 信号，通常是使用 kill -9 终止的。
	ExitCodeSegmentationFault           = 139 // 段错误（Segmentation fault），进程因段错误（非法内存访问）而终止。
	ExitCodeScriptTerminatedBySIGTERM   = 143 // 进程被 SIGTERM 终止（Script terminated by SIGTERM），进程接收到 SIGTERM 信号，通常是使用 kill 终止的。
	ExitCodeExitStatusOutOfRange        = 255 // 退出状态码超出范围（Exit status out of range），一般表示程序返回了超出范围的退出状态码。
)
