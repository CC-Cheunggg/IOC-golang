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

package root

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cc-cheunggg/ioc-golang"
	"github.com/cc-cheunggg/ioc-golang/test/iocli_command"
)

func TestPrintVersion(t *testing.T) {
	output, err := iocli_command.Run([]string{"--version"}, time.Second)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(output, "iocli version "+ioc.Version))
}
