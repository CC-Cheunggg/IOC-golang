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
	"time"

	"github.com/cc-cheunggg/ioc-golang"
	"github.com/cc-cheunggg/ioc-golang/autowire/util"
	"github.com/cc-cheunggg/ioc-golang/example/autowire/autowire_allimpls/service"
	_ "github.com/cc-cheunggg/ioc-golang/example/autowire/autowire_allimpls/service/impl"
	"github.com/cc-cheunggg/ioc-golang/extension/autowire/allimpls"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	ServiceImpls []service.Service `allimpls:""`
}

func (a *App) Run() {
	for {
		time.Sleep(time.Second * 3)
		// get all impls by field injection
		for _, s := range a.ServiceImpls {
			fmt.Println(s.GetHelloString("laurence"))
		}

		// get all impls by API
		allServiceImpls, err := allimpls.GetImpl(util.GetSDIDByStructPtr(new(service.Service)))
		if err != nil {
			panic(err)
		}
		for _, s := range allServiceImpls.([]service.Service) {
			fmt.Println(s.GetHelloString("laurence"))
		}
	}
}

func main() {
	// start
	if err := ioc.Load(); err != nil {
		panic(err)
	}

	app, err := GetAppSingleton()
	if err != nil {
		panic(err)
	}
	app.Run()
}
