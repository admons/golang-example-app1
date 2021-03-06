// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package users_router

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/db"
	"github.com/aristat/golang-example-app/app/db/repo"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/grpc"
	"github.com/aristat/golang-example-app/app/http_routers/oauth-router"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/session"
	"github.com/aristat/golang-example-app/app/tracing"
)

// Injectors from injector.go:

func Build() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sessionConfig, cleanup5, err := session.Cfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup6, err := session.Provider(context, sessionConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup7, err := db.Cfg(viper)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup8, err := db.ProviderGORM(context, zap, dbConfig)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbManager, cleanup9, err := db.Provider(context, zap, dbConfig, gormDB)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauth_routerConfig, cleanup10, err := oauth_router.Cfg(viper)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup11, err := oauth_router.TokenStore(oauth_routerConfig)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientStore, cleanup12, err := oauth_router.ClientStore(oauth_routerConfig)
	if err != nil {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauth_routerManager, cleanup13, err := oauth_router.Provider(context, zap, tokenStore, manager, clientStore)
	if err != nil {
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	usersRepo, cleanup14, err := repo.NewUsersRepo(gormDB)
	if err != nil {
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repoRepo, cleanup15, err := repo.Provider(usersRepo)
	if err != nil {
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup16, err := tracing.ProviderCfg(viper)
	if err != nil {
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup17, err := tracing.Provider(context, configuration, zap)
	if err != nil {
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcConfig, cleanup18, err := grpc.Cfg(viper)
	if err != nil {
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	poolManager, cleanup19, err := grpc.Provider(context, tracer, zap, grpcConfig)
	if err != nil {
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceManagers := ServiceManagers{
		Session:     manager,
		DB:          dbManager,
		Oauth:       oauth_routerManager,
		Repo:        repoRepo,
		PoolManager: poolManager,
	}
	users_routerManager, cleanup20, err := Provider(context, zap, serviceManagers)
	if err != nil {
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return users_routerManager, func() {
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProviderTest()
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfgTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mock, cleanup3, err := logger.ProviderTest(context, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup4, err := session.ProviderTest()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup5, err := db.CfgTest()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup6, err := db.ProviderGORMTest()
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbManager, cleanup7, err := db.Provider(context, mock, dbConfig, gormDB)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup8, err := oauth_router.TokenStoreTest()
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientStore, cleanup9, err := oauth_router.ClientStoreTest()
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauth_routerManager, cleanup10, err := oauth_router.Provider(context, mock, tokenStore, manager, clientStore)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	usersRepo, cleanup11, err := repo.NewUsersRepo(gormDB)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repoRepo, cleanup12, err := repo.Provider(usersRepo)
	if err != nil {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup13, err := tracing.ProviderTest()
	if err != nil {
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcConfig, cleanup14, err := grpc.CfgTest()
	if err != nil {
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	poolManager, cleanup15, err := grpc.Provider(context, tracer, mock, grpcConfig)
	if err != nil {
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceManagers := ServiceManagers{
		Session:     manager,
		DB:          dbManager,
		Oauth:       oauth_routerManager,
		Repo:        repoRepo,
		PoolManager: poolManager,
	}
	users_routerManager, cleanup16, err := Provider(context, mock, serviceManagers)
	if err != nil {
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return users_routerManager, func() {
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
