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

package call

import (
	"github.com/sirupsen/logrus"

	"github.com/alibaba/ioc-golang/logger"
)

type LogConfig struct {
	Level                  string                 `yaml:"level"`
	GlobalLoggerReadOnly   bool                   `yaml:"global-logger-read-only"`
	InvocationAOPLogConfig InvocationAOPLogConfig `yaml:"invocation-aop-log"`
}

type InvocationAOPLogConfig struct {
	Disable              bool   `yaml:"disable"`
	Level                string `yaml:"level"`       // default info
	PrintLevel           string `yaml:"print-level"` // default debug
	DisablePrintParams   bool   `yaml:"disable-print-params"`
	PrintParamsMaxDepth  int    `yaml:"print-params-max-depth"`
	PrintParamsMaxLength int    `yaml:"print-params-max-length"`
}

func (l *LogConfig) fillDefaultConfig() {
	defaultLevel := "info"
	defaultInvocationCtxPrintLevel := "debug"

	// fill logs level
	if l.Level == "" {
		logger.Info("[AOP Log] log config level is using default '%s'", defaultLevel)
		l.Level = defaultLevel
	}
	_, err := logrus.ParseLevel(l.Level)
	if err != nil {
		logger.Error("[AOP Log] parse log config level %s failed with error = %s, using default '%s'", l.Level, err, defaultLevel)
		l.Level = defaultLevel
	}

	// fill InvocationAOPLogConfig
	if l.InvocationAOPLogConfig.PrintParamsMaxDepth <= 0 {
		logger.Info("[AOP Log] log config print-params-max-depth is set to default 3")
		l.InvocationAOPLogConfig.PrintParamsMaxDepth = 3
	}

	// fill invocation aop ctx logs level
	if l.InvocationAOPLogConfig.Level == "" {
		logger.Info("[AOP Log] log config invocation ctx logs level is using default '%s'", defaultLevel)
		l.InvocationAOPLogConfig.Level = defaultLevel
	}
	_, err = logrus.ParseLevel(l.InvocationAOPLogConfig.Level)
	if err != nil {
		logger.Error("[AOP Log] parse log config invocation ctx logs level %s failed with error = %s, using default '%s'", l.InvocationAOPLogConfig.Level, err, defaultLevel)
		l.InvocationAOPLogConfig.Level = defaultLevel
	}

	// fill invocation aop ctx logs print level
	if l.InvocationAOPLogConfig.PrintLevel == "" {
		logger.Info("[AOP Log] log config invocation ctx logs print level is using default '%s'", defaultInvocationCtxPrintLevel)
		l.InvocationAOPLogConfig.PrintLevel = defaultInvocationCtxPrintLevel
	}
	_, err = logrus.ParseLevel(l.InvocationAOPLogConfig.PrintLevel)
	if err != nil {
		logger.Error("[AOP Log] parse log config invocation ctx logs print level %s failed with error = %s, using default '%s'", l.InvocationAOPLogConfig.PrintLevel, err, defaultInvocationCtxPrintLevel)
		l.InvocationAOPLogConfig.PrintLevel = defaultInvocationCtxPrintLevel
	}

	if l.InvocationAOPLogConfig.PrintParamsMaxLength == 0 {
		l.InvocationAOPLogConfig.PrintParamsMaxLength = 1000
	}
}
