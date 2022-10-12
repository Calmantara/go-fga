package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Calmantara/go-fga/config/postgres"
	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/gin-gonic/gin"

	engine "github.com/Calmantara/go-fga/config/gin"
	docs "github.com/Calmantara/go-fga/docs"
	userrepo "github.com/Calmantara/go-fga/pkg/repository/user"
	userhandler "github.com/Calmantara/go-fga/pkg/server/http/handler/user"
	v1 "github.com/Calmantara/go-fga/pkg/server/http/router/v1"
	userusecase "github.com/Calmantara/go-fga/pkg/usecase/user"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// comment dalam go
// untuk beberapa CODE GENERATOR -> tools yang digunakan untuk
// membuat code template di dalam project GO
// ex: swaggo, mockgen, dll
// untuk beberapa tools generator, tools akan membaca comment
// yang memiliki annotation

// @title UserOrder API
// @version 1.0
// @description This is api for creating user and user order
// @termOfService https://swagger.io/terms
// @contact.name FGA API Support
// @host localhost:8080
// @BasePath /
func main() {
	// generate postgres config and connect to postgres
	// this postgres client, will be used in repository layer
	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "postgresAdmin",
		DatabaseName: "postgres",
	})

	// gin engine
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	// setiap request yang datang ke API ini,
	// dia akan melalui gin.Recovery dan gin.Logger
	// .USE disini, adalah cara untuk memasukkan middleware juga
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger())

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		// secara default map jika di return dalam
		// response API, dia akan menjadi JSON
		respMap := map[string]any{
			"code":       0,
			"message":    "server up and running",
			"start_time": startTime,
		}

		// golang memiliki json package
		// json package bisa mentranslasikan
		// map menjadi suatu struct
		// nb: struct harus memiliki tag/annotation JSON
		var respStruct message.Response

		// marshal -> mengubah json/struct/map menjadi
		// array of byte atau bisa kita translatekan menjadi string
		// dengan format JSON
		resByte, err := json.Marshal(respMap)
		if err != nil {
			panic(err)
		}
		// unmarshal -> translasikan string/[]byte dengan format JSON
		// menjadi map/struct dengan tag/annotation json
		err = json.Unmarshal(resByte, &respStruct)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, respStruct)
	})

	docs.SwaggerInfo.BasePath = "/v1"
	ginEngine.GetGin().GET("/swagger/*any", ginswagger.
		WrapHandler(swaggerfiles.Handler))

	// generate user repository
	userRepo := userrepo.NewUserRepo(postgresCln)
	// initiate use case
	userUsecase := userusecase.NewUserUsecase(userRepo)
	// initiate handler
	useHandler := userhandler.NewUserHandler(userUsecase)
	// initiate router
	v1.NewUserRouter(ginEngine, useHandler).Routers()
	v1.NewLoginRouter(ginEngine).Routers()
	// ASSESSMENT
	// buat API
	// - get user
	// sebelum membuat order
	//	- table dengan relasi order -> user (FOREIGN KEY)
	// 			ref:https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-create-table/
	// 	- code base untuk repo, usecase, dll
	// - create order
	// - get order by user

	// running the service
	ginEngine.Serve()
}
