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

package monitor

import (
	"time"

	"github.com/cc-cheunggg/ioc-golang/logger"

	monitorPB "github.com/cc-cheunggg/ioc-golang/extension/aop/monitor/api/ioc_golang/aop/monitor"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:proxy:autoInjection=false

type monitorService struct {
	monitorPB.UnimplementedMonitorServiceServer
	MonitorInterceptor interceptorImplIOCInterface `singleton:""`
}

func (w *monitorService) Monitor(req *monitorPB.MonitorRequest, svr monitorPB.MonitorService_MonitorServer) error {
	logger.Error("[Debug Server] Receive monitor request %s\n", req.String())
	defer logger.Error("[Debug Server] Monitor %s finished \n", req.String())
	sdid := req.GetSdid()
	method := req.GetMethod()
	sendCh := make(chan *monitorPB.MonitorResponse)
	interval := 5
	if reqInterval := req.GetInterval(); reqInterval > 0 {
		interval = int(reqInterval)
	}

	monitorCtx, err := GetcontextIOCInterface(&contextParam{
		SDID:       sdid,
		MethodName: method,
		Ch:         sendCh,
		Period:     time.Duration(interval) * time.Second,
	})
	if err != nil {
		return err
	}
	w.MonitorInterceptor.Monitor(monitorCtx)

	done := svr.Context().Done()
	for {
		select {
		case <-done:
			// monitor stop
			w.MonitorInterceptor.StopMonitor()
			return nil
		case monitorRsp := <-sendCh:
			if err := svr.Send(monitorRsp); err != nil {
				return err
			}
		}
	}
}
