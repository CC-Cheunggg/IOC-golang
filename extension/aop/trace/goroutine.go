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

package trace

import (
	"fmt"
	"sync"

	"github.com/petermattis/goid"

	"github.com/alibaba/ioc-golang/aop"
	"github.com/alibaba/ioc-golang/aop/common"
)

type goRoutineTraceInterceptor struct {
	tracingGrIDMap sync.Map // tracingGrIDMap stores goroutine-id -> goRoutineTracingContext
}

func (g *goRoutineTraceInterceptor) BeforeInvoke(ctx *aop.InvocationContext) {
	// 1. if current goroutine is watched?
	if val, ok := g.tracingGrIDMap.Load(ctx.GrID); ok {
		// this goRoutine is watched, add new child node
		val.(*goRoutineTracingContext).getTrace().addChildSpan(ctx.MethodFullName)
		return
	}
}

func (g *goRoutineTraceInterceptor) AfterInvoke(ctx *aop.InvocationContext) {
	// if current goroutine is watched?
	if val, ok := g.tracingGrIDMap.Load(ctx.GrID); ok {
		// this goRoutine is watched, return span
		traceCtx := val.(*goRoutineTracingContext)
		currentSpan := traceCtx.getTrace().currentSpan
		traceCtx.getTrace().returnSpan()

		// calculate level
		if common.TraceLevel(traceCtx.getTrace().entranceMethod) == 0 {
			// tracing finished
			currentSpan.span.Context().ForeachBaggageItem(func(k, v string) bool {
				fmt.Println(k, v)
				return true
			})
			g.tracingGrIDMap.Delete(ctx.GrID)
		}
	}
}

func (g *goRoutineTraceInterceptor) AddCurrentGRTracingContext(ctx *goRoutineTracingContext) {
	grID := goid.Get()
	g.tracingGrIDMap.Store(grID, ctx)
}

func (g *goRoutineTraceInterceptor) DeleteCurrentGRTracingContext() {
	grID := goid.Get()
	g.tracingGrIDMap.Delete(grID)
}

func (g *goRoutineTraceInterceptor) IsCurrentGRTracing() bool {
	grID := goid.Get()
	_, ok := g.tracingGrIDMap.Load(grID)
	return ok
}

func (g *goRoutineTraceInterceptor) GetCurrentGRTracingContext() *goRoutineTracingContext {
	grID := goid.Get()
	val, ok := g.tracingGrIDMap.Load(grID)
	if ok {
		return val.(*goRoutineTracingContext)
	}
	return nil
}

var goRoutineTraceInterceptorSingleton *goRoutineTraceInterceptor

func getGoRoutineTraceInterceptor() *goRoutineTraceInterceptor {
	if goRoutineTraceInterceptorSingleton == nil {
		goRoutineTraceInterceptorSingleton = &goRoutineTraceInterceptor{}
	}
	return goRoutineTraceInterceptorSingleton
}