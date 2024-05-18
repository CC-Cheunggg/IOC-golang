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

package ioc

import (
	"golang.org/x/sync/errgroup"

	"github.com/cc-cheunggg/ioc-golang/aop"
	"github.com/cc-cheunggg/ioc-golang/autowire"
	"github.com/cc-cheunggg/ioc-golang/config"
	"github.com/cc-cheunggg/ioc-golang/logger"

	_ "github.com/cc-cheunggg/ioc-golang/extension/imports/boot"
)

type LoadHook func() error

var (
	AfterLoad  LoadHook
	BeforeLoad LoadHook
	group      errgroup.Group
)

const Version = "1.0.3.0"

func Load(opts ...config.Option) error {

	if BeforeLoad != nil {
		err := BeforeLoad()
		if err != nil {
			logger.Error("[Boot] Start to call before load handler failure")
		}
	}

	printLogo()
	logger.Debug("Welcome to use ioc-golang %s!", Version)

	// 1. load config
	logger.Info("[Boot] Start to load ioc-golang config")
	if err := config.Load(opts...); err != nil {
		return err
	}

	logger.Info("[Boot] Start to init logger")
	if err := logger.Load(config.LoggerConfig()); err != nil {
		logger.Error("[Boot] Start to init logger unsuccessful")
	}

	logger.Debug("test logger.")

	// 2. load debug
	logger.Info("[Boot] Start to load debug")
	if err := aop.Load(); err != nil {
		return err
	}

	// 3. load autowire
	logger.Info("[Boot] Start to load autowire")
	err := autowire.Load()

	if err == nil && AfterLoad != nil {
		return AfterLoad()
	}

	return err
}
