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

package cli

import (
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/cc-cheunggg/ioc-golang/extension/autowire/allimpls"
	"github.com/cc-cheunggg/ioc-golang/iocli/gen/generator/plugin"
)

const allimplsAutowireTypeAnnotation = "ioc:autowire:allimpls:autowireType"

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/cc-cheunggg/ioc-golang/iocli/gen/generator/plugin.CodeGeneratorPluginForOneStruct
// +ioc:autowire:allimpls:autowireType=normal
// +ioc:autowire:constructFunc=create

type allImplsCodeGenerationPlugin struct {
	allimplsAutowireType string
}

func create(t *allImplsCodeGenerationPlugin) (*allImplsCodeGenerationPlugin, error) {
	return t, nil
}

func (t *allImplsCodeGenerationPlugin) Name() string {
	return allimpls.Name
}

func (t *allImplsCodeGenerationPlugin) Type() plugin.Type {
	return plugin.Autowire
}

func (t *allImplsCodeGenerationPlugin) Init(info markers.TypeInfo) {
	allimplsAutowireType := ""
	if allimplsAutowireTypeValues := info.Markers[allimplsAutowireTypeAnnotation]; len(allimplsAutowireTypeValues) > 0 {
		allimplsAutowireType = allimplsAutowireTypeValues[0].(string)
	}
	t.allimplsAutowireType = allimplsAutowireType
}

func (t *allImplsCodeGenerationPlugin) GenerateSDMetadataForOneStruct(root *loader.Package, c plugin.CodeWriter) {
	if t.allimplsAutowireType != "" {
		c.Linef(`"%s": map[string]interface{}{`, allimpls.Name)
		c.Linef(`"%s":"%s",`, allimpls.AutowireTypeMetadataKey, t.allimplsAutowireType)
		c.Line(`},`)
	}
}

func (t *allImplsCodeGenerationPlugin) GenerateInFileForOneStruct(root *loader.Package, c plugin.CodeWriter) {
}
