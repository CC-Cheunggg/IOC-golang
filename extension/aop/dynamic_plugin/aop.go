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

package dynamic_plugin

import (
	"google.golang.org/grpc"

	"github.com/cc-cheunggg/ioc-golang/aop"
	"github.com/cc-cheunggg/ioc-golang/extension/aop/dynamic_plugin/api/ioc_golang/aop/dynamic_plugin"
)

const Name = "dynamic_plugin"

func init() {
	aop.RegisterAOP(aop.AOP{
		Name: Name,
		GRPCServiceRegister: func(server *grpc.Server) {
			listServiceImplSingleton, _ := GetdynamicPluginServiceImplSingleton()
			dynamic_plugin.RegisterDynamicPluginServiceServer(server, listServiceImplSingleton)
		},
	})
}
