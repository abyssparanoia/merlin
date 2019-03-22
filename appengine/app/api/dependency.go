package main

import (
	"os"

	"github.com/abyssparanoia/merlin/src/handler/api"
	"github.com/abyssparanoia/merlin/src/lib/cloudfirestore"
	"github.com/abyssparanoia/merlin/src/lib/firebaseauth"
	"github.com/abyssparanoia/merlin/src/lib/httpheader"
	"github.com/abyssparanoia/merlin/src/lib/jsonrpc2"
	"github.com/abyssparanoia/merlin/src/lib/mysql"
	"github.com/abyssparanoia/merlin/src/repository"
	"github.com/abyssparanoia/merlin/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	FirebaseAuth    *firebaseauth.Middleware
	HTTPHeader      *httpheader.Middleware
	JSONRPC2Handler *jsonrpc2.Handler
	SampleHandler   *api.SampleHandler
	UserHandler     *api.UserHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	crePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crePath == "" {
		panic("no config error: GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Client
	fCli, err := cloudfirestore.NewClient(crePath)
	if err != nil {
		panic(err.Error())
	}

	// Config
	dbCfg := mysql.GetSQLConfig()

	// Lib
	dbConn := mysql.NewSQLClient(dbCfg)
	dbConn.LogMode(true)

	// Repository
	repo := repository.NewSample(fCli)
	uRepo := repository.NewUser(dbConn)

	// Service
	faSvc := firebaseauth.NewService()
	hhSvc := httpheader.NewService()
	svc := service.NewSample(repo)
	uSvc := service.NewUser(uRepo)

	// Middleware
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.HTTPHeader = httpheader.NewMiddleware(hhSvc)

	// Handler
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
	d.SampleHandler = api.NewSampleHandler(svc)
	d.UserHandler = api.NewUserHandler(uSvc)
}
