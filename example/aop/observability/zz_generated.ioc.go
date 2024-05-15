//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package main

import (
	autowire "github.com/cc-cheunggg/ioc-golang/autowire"
	normal "github.com/cc-cheunggg/ioc-golang/autowire/normal"
	singleton "github.com/cc-cheunggg/ioc-golang/autowire/singleton"
	util "github.com/cc-cheunggg/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &app_{}
		},
	})
	appStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &App{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(appStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1_{}
		},
	})
	serviceImpl1StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl1{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serviceImpl1StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2_{}
		},
	})
	serviceImpl2StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl2{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serviceImpl2StructDescriptor)
}

type app_ struct {
	Run_ func()
}

func (a *app_) Run() {
	a.Run_()
}

type serviceImpl1_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl1_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceImpl2_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl2_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type AppIOCInterface interface {
	Run()
}

type ServiceImpl1IOCInterface interface {
	GetHelloString(name string) string
}

type ServiceImpl2IOCInterface interface {
	GetHelloString(name string) string
}

var _appSDID string

func GetAppSingleton() (*App, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImpl(_appSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*App)
	return impl, nil
}

func GetAppIOCInterfaceSingleton() (AppIOCInterface, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImplWithProxy(_appSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(AppIOCInterface)
	return impl, nil
}

type ThisApp struct {
}

func (t *ThisApp) This() AppIOCInterface {
	thisPtr, _ := GetAppIOCInterfaceSingleton()
	return thisPtr
}

var _serviceImpl1SDID string

func GetServiceImpl1Singleton() (*ServiceImpl1, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := singleton.GetImpl(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl1)
	return impl, nil
}

func GetServiceImpl1IOCInterfaceSingleton() (ServiceImpl1IOCInterface, error) {
	if _serviceImpl1SDID == "" {
		_serviceImpl1SDID = util.GetSDIDByStructPtr(new(ServiceImpl1))
	}
	i, err := singleton.GetImplWithProxy(_serviceImpl1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl1IOCInterface)
	return impl, nil
}

type ThisServiceImpl1 struct {
}

func (t *ThisServiceImpl1) This() ServiceImpl1IOCInterface {
	thisPtr, _ := GetServiceImpl1IOCInterfaceSingleton()
	return thisPtr
}

var _serviceImpl2SDID string

func GetServiceImpl2Singleton() (*ServiceImpl2, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := singleton.GetImpl(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl2)
	return impl, nil
}

func GetServiceImpl2IOCInterfaceSingleton() (ServiceImpl2IOCInterface, error) {
	if _serviceImpl2SDID == "" {
		_serviceImpl2SDID = util.GetSDIDByStructPtr(new(ServiceImpl2))
	}
	i, err := singleton.GetImplWithProxy(_serviceImpl2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl2IOCInterface)
	return impl, nil
}

type ThisServiceImpl2 struct {
}

func (t *ThisServiceImpl2) This() ServiceImpl2IOCInterface {
	thisPtr, _ := GetServiceImpl2IOCInterfaceSingleton()
	return thisPtr
}
