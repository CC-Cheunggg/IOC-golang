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

package main

import (
	"fmt"
	"os"

	"github.com/cc-cheunggg/ioc-golang"
	configField "github.com/cc-cheunggg/ioc-golang/extension/config"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	ConfigValue        *configField.ConfigString `config:",config.app.config-value"`
	ConfigValueFromEnv *configField.ConfigString `config:",config.app.config-value-from-env"`
}

func (a *App) Run() {
	fmt.Printf("Load '%s' from config file\n", a.ConfigValue.Value())
	fmt.Printf("Load '%s' from env\n", a.ConfigValueFromEnv.Value())
}

func main() {
	// start
	os.Setenv("MY_CONFIG_ENV_KEY", "myEnvValue")
	if err := ioc.Load(); err != nil {
		panic(err)
	}

	app, err := GetAppIOCInterfaceSingleton()
	if err != nil {
		panic(err)
	}
	app.Run()
}
