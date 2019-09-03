// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package logger

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/entrypoint"
)

// Injectors from injector.go:

func Build() (Logger, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return zap, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (Logger, func(), error) {
	context, cleanup, err := entrypoint.ContextProviderTest()
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := ProviderCfgTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mock, cleanup3, err := ProviderTest(context, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return mock, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}