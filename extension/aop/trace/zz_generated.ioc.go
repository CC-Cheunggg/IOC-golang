//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package trace

import (
	"github.com/cc-cheunggg/ioc-golang/aop"
	autowire "github.com/cc-cheunggg/ioc-golang/autowire"
	normal "github.com/cc-cheunggg/ioc-golang/autowire/normal"
	singleton "github.com/cc-cheunggg/ioc-golang/autowire/singleton"
	util "github.com/cc-cheunggg/ioc-golang/autowire/util"
	aoptrace "github.com/cc-cheunggg/ioc-golang/extension/aop/trace/api/ioc_golang/aop/trace"
	"github.com/gin-gonic/gin"
	opentracing_go "github.com/opentracing/opentracing-go"
	"net/http"
)

func init() {
	debugServerTraceByMethodContextStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &debugServerTraceByMethodContext{}
		},
		ParamFactory: func() interface{} {
			var _ debugServerTraceByMethodContextParamInterface = &debugServerTraceByMethodContextParam{}
			return &debugServerTraceByMethodContextParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(debugServerTraceByMethodContextParamInterface)
			impl := i.(*debugServerTraceByMethodContext)
			return param.initDebugServerTraceByMethodContext(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	normal.RegisterStructDescriptor(debugServerTraceByMethodContextStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceInterceptor_{}
		},
	})
	traceInterceptorStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceInterceptor{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(traceInterceptorStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceGoRoutineInterceptorFacadeCtx_{}
		},
	})
	traceGoRoutineInterceptorFacadeCtxStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceGoRoutineInterceptorFacadeCtx{}
		},
		ParamFactory: func() interface{} {
			var _ traceGoRoutineInterceptorFacadeCtxParamInterface = &traceGoRoutineInterceptorFacadeCtxParam{}
			return &traceGoRoutineInterceptorFacadeCtxParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(traceGoRoutineInterceptorFacadeCtxParamInterface)
			impl := i.(*traceGoRoutineInterceptorFacadeCtx)
			return param.newTraceGoRoutineInterceptorFacadeCtx(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(traceGoRoutineInterceptorFacadeCtxStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &rpcInterceptor_{}
		},
	})
	rpcInterceptorStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &rpcInterceptor{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*rpcInterceptor)
			var constructFunc rpcInterceptorConstructFunc = initRPCInterceptor
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(rpcInterceptorStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceServiceImpl_{}
		},
	})
	traceServiceImplStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &traceServiceImpl{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(traceServiceImplStructDescriptor)
}

type debugServerTraceByMethodContextParamInterface interface {
	initDebugServerTraceByMethodContext(impl *debugServerTraceByMethodContext) (*debugServerTraceByMethodContext, error)
}
type traceGoRoutineInterceptorFacadeCtxParamInterface interface {
	newTraceGoRoutineInterceptorFacadeCtx(impl *traceGoRoutineInterceptorFacadeCtx) (*traceGoRoutineInterceptorFacadeCtx, error)
}
type rpcInterceptorConstructFunc func(impl *rpcInterceptor) (*rpcInterceptor, error)
type traceInterceptor_ struct {
	BeforeInvoke_       func(ctx *aop.InvocationContext)
	AfterInvoke_        func(ctx *aop.InvocationContext)
	StartTraceByMethod_ func(traceCtx *debugServerTraceByMethodContext)
	StopTraceByMethod_  func()
	GetCurrentSpan_     func() opentracing_go.Span
}

func (t *traceInterceptor_) BeforeInvoke(ctx *aop.InvocationContext) {
	t.BeforeInvoke_(ctx)
}

func (t *traceInterceptor_) AfterInvoke(ctx *aop.InvocationContext) {
	t.AfterInvoke_(ctx)
}

func (t *traceInterceptor_) StartTraceByMethod(traceCtx *debugServerTraceByMethodContext) {
	t.StartTraceByMethod_(traceCtx)
}

func (t *traceInterceptor_) StopTraceByMethod() {
	t.StopTraceByMethod_()
}

func (t *traceInterceptor_) GetCurrentSpan() opentracing_go.Span {
	return t.GetCurrentSpan_()
}

type traceGoRoutineInterceptorFacadeCtx_ struct {
	BeforeInvoke_ func(ctx *aop.InvocationContext)
	AfterInvoke_  func(ctx *aop.InvocationContext)
	Type_         func() string
}

func (t *traceGoRoutineInterceptorFacadeCtx_) BeforeInvoke(ctx *aop.InvocationContext) {
	t.BeforeInvoke_(ctx)
}

func (t *traceGoRoutineInterceptorFacadeCtx_) AfterInvoke(ctx *aop.InvocationContext) {
	t.AfterInvoke_(ctx)
}

func (t *traceGoRoutineInterceptorFacadeCtx_) Type() string {
	return t.Type_()
}

type rpcInterceptor_ struct {
	BeforeServerInvoke_ func(c *gin.Context) error
	AfterServerInvoke_  func(ctx *gin.Context) error
	BeforeClientInvoke_ func(req *http.Request) error
	AfterClientInvoke_  func(response *http.Response) error
}

func (r *rpcInterceptor_) BeforeServerInvoke(c *gin.Context) error {
	return r.BeforeServerInvoke_(c)
}

func (r *rpcInterceptor_) AfterServerInvoke(ctx *gin.Context) error {
	return r.AfterServerInvoke_(ctx)
}

func (r *rpcInterceptor_) BeforeClientInvoke(req *http.Request) error {
	return r.BeforeClientInvoke_(req)
}

func (r *rpcInterceptor_) AfterClientInvoke(response *http.Response) error {
	return r.AfterClientInvoke_(response)
}

type traceServiceImpl_ struct {
	Trace_ func(req *aoptrace.TraceRequest, traceServer aoptrace.TraceService_TraceServer) error
}

func (t *traceServiceImpl_) Trace(req *aoptrace.TraceRequest, traceServer aoptrace.TraceService_TraceServer) error {
	return t.Trace_(req, traceServer)
}

type traceInterceptorIOCInterface interface {
	BeforeInvoke(ctx *aop.InvocationContext)
	AfterInvoke(ctx *aop.InvocationContext)
	StartTraceByMethod(traceCtx *debugServerTraceByMethodContext)
	StopTraceByMethod()
	GetCurrentSpan() opentracing_go.Span
}

type traceGoRoutineInterceptorFacadeCtxIOCInterface interface {
	BeforeInvoke(ctx *aop.InvocationContext)
	AfterInvoke(ctx *aop.InvocationContext)
	Type() string
}

type rpcInterceptorIOCInterface interface {
	BeforeServerInvoke(c *gin.Context) error
	AfterServerInvoke(ctx *gin.Context) error
	BeforeClientInvoke(req *http.Request) error
	AfterClientInvoke(response *http.Response) error
}

type traceServiceImplIOCInterface interface {
	Trace(req *aoptrace.TraceRequest, traceServer aoptrace.TraceService_TraceServer) error
}

var _debugServerTraceByMethodContextSDID string

func GetdebugServerTraceByMethodContext(p *debugServerTraceByMethodContextParam) (*debugServerTraceByMethodContext, error) {
	if _debugServerTraceByMethodContextSDID == "" {
		_debugServerTraceByMethodContextSDID = util.GetSDIDByStructPtr(new(debugServerTraceByMethodContext))
	}
	i, err := normal.GetImpl(_debugServerTraceByMethodContextSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*debugServerTraceByMethodContext)
	return impl, nil
}

var _traceInterceptorSDID string

func GettraceInterceptorSingleton() (*traceInterceptor, error) {
	if _traceInterceptorSDID == "" {
		_traceInterceptorSDID = util.GetSDIDByStructPtr(new(traceInterceptor))
	}
	i, err := singleton.GetImpl(_traceInterceptorSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*traceInterceptor)
	return impl, nil
}

func GettraceInterceptorIOCInterfaceSingleton() (traceInterceptorIOCInterface, error) {
	if _traceInterceptorSDID == "" {
		_traceInterceptorSDID = util.GetSDIDByStructPtr(new(traceInterceptor))
	}
	i, err := singleton.GetImplWithProxy(_traceInterceptorSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(traceInterceptorIOCInterface)
	return impl, nil
}

type ThistraceInterceptor struct {
}

func (t *ThistraceInterceptor) This() traceInterceptorIOCInterface {
	thisPtr, _ := GettraceInterceptorIOCInterfaceSingleton()
	return thisPtr
}

var _traceGoRoutineInterceptorFacadeCtxSDID string

func GettraceGoRoutineInterceptorFacadeCtx(p *traceGoRoutineInterceptorFacadeCtxParam) (*traceGoRoutineInterceptorFacadeCtx, error) {
	if _traceGoRoutineInterceptorFacadeCtxSDID == "" {
		_traceGoRoutineInterceptorFacadeCtxSDID = util.GetSDIDByStructPtr(new(traceGoRoutineInterceptorFacadeCtx))
	}
	i, err := normal.GetImpl(_traceGoRoutineInterceptorFacadeCtxSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*traceGoRoutineInterceptorFacadeCtx)
	return impl, nil
}

func GettraceGoRoutineInterceptorFacadeCtxIOCInterface(p *traceGoRoutineInterceptorFacadeCtxParam) (traceGoRoutineInterceptorFacadeCtxIOCInterface, error) {
	if _traceGoRoutineInterceptorFacadeCtxSDID == "" {
		_traceGoRoutineInterceptorFacadeCtxSDID = util.GetSDIDByStructPtr(new(traceGoRoutineInterceptorFacadeCtx))
	}
	i, err := normal.GetImplWithProxy(_traceGoRoutineInterceptorFacadeCtxSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(traceGoRoutineInterceptorFacadeCtxIOCInterface)
	return impl, nil
}

var _rpcInterceptorSDID string

func GetrpcInterceptorSingleton() (*rpcInterceptor, error) {
	if _rpcInterceptorSDID == "" {
		_rpcInterceptorSDID = util.GetSDIDByStructPtr(new(rpcInterceptor))
	}
	i, err := singleton.GetImpl(_rpcInterceptorSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*rpcInterceptor)
	return impl, nil
}

func GetrpcInterceptorIOCInterfaceSingleton() (rpcInterceptorIOCInterface, error) {
	if _rpcInterceptorSDID == "" {
		_rpcInterceptorSDID = util.GetSDIDByStructPtr(new(rpcInterceptor))
	}
	i, err := singleton.GetImplWithProxy(_rpcInterceptorSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(rpcInterceptorIOCInterface)
	return impl, nil
}

type ThisrpcInterceptor struct {
}

func (t *ThisrpcInterceptor) This() rpcInterceptorIOCInterface {
	thisPtr, _ := GetrpcInterceptorIOCInterfaceSingleton()
	return thisPtr
}

var _traceServiceImplSDID string

func GettraceServiceImplSingleton() (*traceServiceImpl, error) {
	if _traceServiceImplSDID == "" {
		_traceServiceImplSDID = util.GetSDIDByStructPtr(new(traceServiceImpl))
	}
	i, err := singleton.GetImpl(_traceServiceImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*traceServiceImpl)
	return impl, nil
}

func GettraceServiceImplIOCInterfaceSingleton() (traceServiceImplIOCInterface, error) {
	if _traceServiceImplSDID == "" {
		_traceServiceImplSDID = util.GetSDIDByStructPtr(new(traceServiceImpl))
	}
	i, err := singleton.GetImplWithProxy(_traceServiceImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(traceServiceImplIOCInterface)
	return impl, nil
}

type ThistraceServiceImpl struct {
}

func (t *ThistraceServiceImpl) This() traceServiceImplIOCInterface {
	thisPtr, _ := GettraceServiceImplIOCInterfaceSingleton()
	return thisPtr
}
