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

package aop

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"

	"github.com/cc-cheunggg/ioc-golang"
	"github.com/cc-cheunggg/ioc-golang/test/iocli_command"
)

type mockWriter struct {
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func TestAOPConcurrent(t *testing.T) {
	closeCh := make(chan struct{})
	// set empty writer to make sure aop logs not be printed, to avoid timeout
	logrus.SetOutput(&mockWriter{})
	assert.Nil(t, ioc.Load())
	go func() {
		output, err := iocli_command.Run([]string{"monitor"}, time.Second*6)
		assert.Nil(t, err)
		t.Log(output)
		assert.True(t, strings.Contains(output, `github.com/cc-cheunggg/ioc-golang/test/stress/aop.NormalApp.RunTest()
Total: 10000, Success: 10000, Fail: 0, AvgRT: `))
		assert.True(t, strings.Contains(output, `us, FailRate: 0.00%
github.com/cc-cheunggg/ioc-golang/test/stress/aop.ServiceImpl1.GetHelloString()
Total: 10000, Success: 10000, Fail: 0, AvgRT: `))
		close(closeCh)
	}()
	time.Sleep(time.Second * 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				normalApp, err := GetNormalAppIOCInterface()
				assert.Nil(t, err)
				normalApp.RunTest(t)
			}
		}()
	}
	wg.Wait()
	<-closeCh
}

func TestAOPRecursive(t *testing.T) {
	logrus.SetOutput(&mockWriter{})
	assert.Nil(t, ioc.Load())
	closeCh := make(chan struct{})
	go func() {
		output, err := iocli_command.Run([]string{"monitor"}, time.Second*6)
		assert.Nil(t, err)
		t.Log(output)
		assert.True(t, strings.Contains(output, `github.com/cc-cheunggg/ioc-golang/test/stress/aop.RecursiveApp.RunTest()
Total: 901, Success: 901, Fail: 0, AvgRT: `))
		assert.True(t, strings.Contains(output, `us, FailRate: 0.00%
github.com/cc-cheunggg/ioc-golang/test/stress/aop.ServiceImpl1.GetHelloString()
Total: 2, Success: 2, Fail: 0, AvgRT: `))
		close(closeCh)
	}()
	time.Sleep(time.Second * 1)
	recApp, err := GetRecursiveAppIOCInterfaceSingleton()
	assert.Nil(t, err)
	recApp.RunTest(t)
	<-closeCh

	recApp.Reset()
	closeCh = make(chan struct{})
	go func() {
		output, err := iocli_command.Run([]string{"trace"}, time.Second*6)
		assert.Nil(t, err)
		assert.Equal(t, 901, strings.Count(output, ", OperationName: github.com/cc-cheunggg/ioc-golang/test/stress/aop.(*recursiveApp_).RunTest, StartTime: "))
		close(closeCh)
	}()
	time.Sleep(time.Second * 1)
	recApp.RunTest(t)
	<-closeCh
}
