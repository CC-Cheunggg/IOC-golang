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

package rpc_service

import (
	"github.com/cc-cheunggg/ioc-golang/autowire"
	"github.com/cc-cheunggg/ioc-golang/autowire/singleton"
	"github.com/cc-cheunggg/ioc-golang/extension/autowire/rpc/protocol/protocol_impl"
	"github.com/cc-cheunggg/ioc-golang/extension/autowire/rpc/proxy"
)

func init() {
	autowire.RegisterAutowire(func() autowire.Autowire {
		rpcAutowire := &Autowire{}
		// todo parse rpc param
		rpcAutowire.Autowire = singleton.NewSingletonAutowire(nil, nil, rpcAutowire)
		return rpcAutowire
	}())
}

const Name = "rpc-service"

type Autowire struct {
	autowire.Autowire
}

func (a *Autowire) TagKey() string {
	return Name
}

func (a *Autowire) CanBeEntrance() bool {
	return true
}

// GetAllStructDescriptors re-write SingletonAutowire
func (a *Autowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	return rpcStructDescriptorMap
}

var rpcStructDescriptorMap = make(map[string]*autowire.StructDescriptor)

func RegisterStructDescriptor(s *autowire.StructDescriptor) {
	sdID := s.ID()
	var originConstructFunc func(impl interface{}, _ interface{}) (interface{}, error)
	if s.ConstructFunc != nil {
		originConstructFunc = s.ConstructFunc
	}
	s.ConstructFunc = func(impl interface{}, _ interface{}) (interface{}, error) {
		if originConstructFunc != nil {
			var err error
			impl, err = originConstructFunc(impl, nil)
			if err != nil {
				return nil, err
			}
		}
		// if field is interface, try to inject proxy wrapped pointer
		impl = autowire.GetProxyFunction()(impl)

		// param not configured in server side, set default param
		iocProtocolInterface, err := protocol_impl.GetIOCProtocolIOCInterface(&protocol_impl.Param{
			ExportPort: defaultParam.ExportPort,
		})
		if err != nil {
			return nil, err
		}
		_, err = proxy.MetadataMap.Register(sdID, protocol_impl.IOCProtocolName, "", "", impl)
		if err != nil {
			panic(err)
		}

		defaultProxyInvoker := proxy.NewProxyInvoker(protocol_impl.IOCProtocolName, sdID, s.Alias)
		iocProtocolInterface.Export(defaultProxyInvoker)
		return impl, nil
	}
	rpcStructDescriptorMap[sdID] = s
	autowire.RegisterStructDescriptor(s)
}

func GetImpl(key string) (interface{}, error) {
	return autowire.Impl(Name, key, nil)
}

func GetImplWithProxy(key string) (interface{}, error) {
	return autowire.ImplWithProxy(Name, key, nil)
}
