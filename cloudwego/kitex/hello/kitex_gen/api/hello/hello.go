// Code generated by Kitex v0.11.3. DO NOT EDIT.

package hello

import (
	"context"
	"errors"
	api "github.com/ahang7/go-hello/cloudwego/kitex/hello/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo": kitex.NewMethodInfo(
		echoHandler,
		newHelloEchoArgs,
		newHelloEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"add": kitex.NewMethodInfo(
		addHandler,
		newHelloAddArgs,
		newHelloAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	helloServiceInfo                = NewServiceInfo()
	helloServiceInfoForClient       = NewServiceInfoForClient()
	helloServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return helloServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return helloServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*api.Hello)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloEchoArgs)
	realResult := result.(*api.HelloEchoResult)
	success, err := handler.(api.Hello).Echo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloEchoArgs() interface{} {
	return api.NewHelloEchoArgs()
}

func newHelloEchoResult() interface{} {
	return api.NewHelloEchoResult()
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloAddArgs)
	realResult := result.(*api.HelloAddResult)
	success, err := handler.(api.Hello).Add(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloAddArgs() interface{} {
	return api.NewHelloAddArgs()
}

func newHelloAddResult() interface{} {
	return api.NewHelloAddResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, request *api.Request) (r *api.Response, err error) {
	var _args api.HelloEchoArgs
	_args.Request = request
	var _result api.HelloEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Add(ctx context.Context, request *api.AddRequest) (r *api.AddResponse, err error) {
	var _args api.HelloAddArgs
	_args.Request = request
	var _result api.HelloAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
