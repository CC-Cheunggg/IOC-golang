/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	disableLogs  = false
	logrusLogger *logrus.Logger
	zapLogger    *zap.SugaredLogger
)

func Info(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	if logrusLogger != nil {
		logrusLogger.Infof(format, a...)
	}
	if zapLogger != nil {
		zapLogger.Infof(format, a...)
	}
	color.Blue(format, a...)
}
func Debug(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	if logrusLogger != nil {
		logrusLogger.Debugf(format, a...)
	}
	if zapLogger != nil {
		zapLogger.Debugf(format, a...)
	}
	color.Cyan(format, a...)
}

func Error(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	if logrusLogger != nil {
		logrusLogger.Errorf(format, a...)
	}
	if zapLogger != nil {
		zapLogger.Errorf(format, a...)
	}
	color.Red(format, a...)
}

func Disable() {
	disableLogs = true
}

func SetLogrus(logger *logrus.Logger) {
	logrusLogger = logger
}

func SetZap(logger *zap.Logger) {
	zapLogger = logger.Sugar()
}
