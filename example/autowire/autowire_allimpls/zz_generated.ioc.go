//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package main

import (
	autowire "github.com/cc-cheunggg/ioc-golang/autowire"
	normal "github.com/cc-cheunggg/ioc-golang/autowire/normal"
	singleton "github.com/cc-cheunggg/ioc-golang/autowire/singleton"
	"github.com/cc-cheunggg/ioc-golang/autowire/util"
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
}

type app_ struct {
	Run_ func()
}

func (a *app_) Run() {
	a.Run_()
}

type AppIOCInterface interface {
	Run()
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
