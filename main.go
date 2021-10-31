package main

import (
	"AltaStore/api"
	"AltaStore/api/middleware"
	"AltaStore/config"

	// Controller
	checkoutController "AltaStore/api/v1/checkout"
	paymentController "AltaStore/api/v1/checkoutpayment"
	shopController "AltaStore/api/v1/shopping"

	// Service
	checkoutService "AltaStore/business/checkout"
	paymentService "AltaStore/business/checkoutpayment"
	loggerService "AltaStore/business/logger"
	shopService "AltaStore/business/shopping"
	userService "AltaStore/business/user"

	// Repository
	checkoutRepository "AltaStore/modules/checkout"
	paymentRepository "AltaStore/modules/checkoutpayment"
	loggerRepo "AltaStore/modules/logger"
	shopRepository "AltaStore/modules/shopping"
	shopDetailRepository "AltaStore/modules/shoppingdetail"
	userRepository "AltaStore/modules/user"

	"context"
	"fmt"
	"time"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"AltaStore/modules/migration"
)

func newDatabaseConnection(cfg *config.ConfigApp) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName,
	)
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func newMongoDBConnection(cfg *config.ConfigApp) *mongo.Database {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoUsername, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort),
	)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	return client.Database(cfg.MongoDbName)
}

func main() {
	// retrieves application configuration and returns common values when there is a problem
	config := config.GetConfig()

	// Open mongodb logger
	mongoConnection := newMongoDBConnection(config)

	// Register repository
	logrRepo := loggerRepo.NewRepository(mongoConnection)

	// Register service
	logeService := loggerService.NewService(logrRepo)

	// Register logs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logeService)

	// open database server base session
	dbConnection := newDatabaseConnection(config)

	//initiate user repository
	user := userRepository.NewDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(user)

	// initiate shopping repository
	shopRepo := shopRepository.NewRepository(dbConnection)
	shopDetailRepo := shopDetailRepository.NewRepository(dbConnection)

	// initiate urchase Receiving service
	shopServc := shopService.NewService(shopRepo, shopDetailRepo)

	// initiate shopping controller
	shopHandler := shopController.NewController(shopServc)

	// initiate CheckOut Payment repository
	payment := paymentRepository.NewRepository(dbConnection)

	// initiate CheckOut Payment service
	paymentService := paymentService.NewService(userService, payment)

	// initiate CheckOut Payment controller
	paymentController := paymentController.NewController(paymentService)

	// initiate checkout repository shoping cart
	c_outeRepo := checkoutRepository.NewRepository(dbConnection)

	// initiate checkout service shopping cat
	c_outServc := checkoutService.NewService(paymentService, shopServc, c_outeRepo, shopDetailRepo)

	// initiate checkout controller shopingcart
	c_outController := checkoutController.NewController(c_outServc)

	// create echo http
	e := echo.New()

	// Register API Path and Controller
	api.RegisterPath(e, shopHandler, c_outController, paymentController)

	lock := make(chan error)

	go func(lock chan error) {
		address := fmt.Sprintf(":%d", config.AppPort)
		lock <- e.Start(address)
	}(lock)

	time.Sleep(1 * time.Millisecond)
	middleware.MakeLogEntry(nil).Info(fmt.Sprintf("Application Start In Port => ::%d", config.AppPort))

	err := <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("Shutdown Echo Service")
	}
}
