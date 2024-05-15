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

package common

type Config struct {
	Engine        string `yaml:"engine"`         // 日志引擎
	Level         string `yaml:"level"`          // 级别
	Prefix        string `yaml:"prefix"`         // 日志前缀
	Format        string `yaml:"format"`         // 输出
	Director      string `yaml:"director"`       // 日志文件夹
	Encoder       string `yaml:"encoder"`        // 编码级
	StacktraceKey string `yaml:"stacktrace-key"` // 栈名
	MaxAge        int64  `yaml:"max-age"`        // 日志留存时间
	ShowLine      bool   `yaml:"show-line"`      // 显示行
	Std           bool   `yaml:"std"`            // 是否控制台打印
}
