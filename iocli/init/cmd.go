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

package init

import (
	"github.com/spf13/cobra"

	"github.com/cc-cheunggg/ioc-golang/iocli/root"
)

var (
	path      string
	name      string
	modPrefix string
	modName   string
)

func init() {
	root.Cmd.AddCommand(initCMD)
	initCMD.PersistentFlags().StringVarP(&path, "dir", "d", ".", "Specify the directory of the project")
	initCMD.PersistentFlags().StringVarP(&name, "name", "n", "helloioc", "Specify the name of the project")
	initCMD.PersistentFlags().StringVarP(&modPrefix, "modPrefix", "p", "github.com/alibaba", "Specify the mod prefix of the project")
	initCMD.PersistentFlags().StringVarP(&modName, "modName", "m", "", "Specify the mod name of the project")
}

// initCMD iocli init cmd.
//
// install pkger:
//
// $ go get github.com/markbates/pkger/cmd/pkger
//
// $ pkger -h
//
// e.g.:
//
// $ iocli init -d hello -n helloiocgo -m github.com/photowey/helloiocgo
//
// $ iocli init -d hello -n helloiocgo -p github.com/photowey
var initCMD = &cobra.Command{
	Use:   "init",
	Short: "Create IOC-golang scaffold template project.",
	Long:  "Create IOC-golang scaffold template project.",
	Example: `Create IOC-golang scaffold template project.

$ iocli init -d [project.path]       // default: .
$ iocli init -n [project.name]       // default: helloioc
$ iocli init -p [project.mod.prefix] // default search: .../go.mod
$ iocli init -m [project.mod.name]   // default search: .../go.mod

e.g.:
$ iocli init -d hello -n helloiocgo
$ iocli init -d hello -n helloiocgo -p github.com/alibaba
$ iocli init -d hello -n helloiocgo -m github.com/alibaba/helloiocgo
$ iocli init -d $GOPATH/src/github.com/alibaba -n helloiocgo -m github.com/alibaba/helloiocgo
`,
	Run: func(c *cobra.Command, args []string) {
		Run(
			WithPath(path),
			WithName(name),
			WithModPrefix(modPrefix),
			WithModName(modName),
		)
	},
}
