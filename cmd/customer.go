package cmd

import (
	cfg "CustomerAPI/configs"
	"CustomerAPI/docs"
	customerHandler "CustomerAPI/internal/customer/handlers"
	customerRepository "CustomerAPI/internal/customer/storages/mongo"
	customMiddleware "CustomerAPI/pkg/middlewares"
	"CustomerAPI/pkg/utils"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func Execute(env string) {
	cfg.SetConfigs(env)
	docs.SwaggerInfo.Host = utils.GetSwagHostEnv()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(customMiddleware.PanicExceptionHandling())

	opts := options.Client().ApplyURI(cfg.GetConfigs().MongoConnectionURI)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var mc, err = mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mc.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	//TODO: include logging as well.
	//logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.ANSIC})

	customerRepo := customerRepository.NewRepository(mc.Database(cfg.GetConfigs().CustomerDBName).Collection(cfg.GetConfigs().CustomerCollectionName))
	customerHandler.NewHandler(e, customerRepo)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//TODO: include healthcheck as well.
	//e.GET("/health", health_check.HealthCheck)

	log.Fatal(e.Start(":1323"))
}
