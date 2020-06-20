package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	bannerHandler "sieo_app/banner/delivery"
	bannerRepo "sieo_app/banner/repository"
	bannerService "sieo_app/banner/service"
	eoDelivery "sieo_app/eo/delivery"
	eoRepo "sieo_app/eo/repository"
	eoService "sieo_app/eo/service"
	eventHandler "sieo_app/event/delivery"
	eventRepo "sieo_app/event/repository"
	eventService "sieo_app/event/service"
	"sieo_app/models"
	userHandler "sieo_app/user/delivery"
	_userRepository "sieo_app/user/repository"
	_userService "sieo_app/user/service"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	port := viper.GetString("port.port")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	sqlConn, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
		return
	}

	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
		return
	}

	err = dbConn.DB().Ping()
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println("Success Connect DB")

	defer func() {
		err = dbConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	dbConn.Debug().AutoMigrate(
		&models.User{},
		&models.Eo{},
		&models.Event{},
		&models.Banner{},
	)

	route := mux.NewRouter().StrictSlash(true)

	bannerRepoUse := bannerRepo.CreateBannerRepoImpl(dbConn)
	bannerServiceUse := bannerService.CreateBannerService(bannerRepoUse)
	bannerHandler.CreateBannerHandler(route, bannerServiceUse)

	userRepository := _userRepository.CreateUserRepoImlp(dbConn)
	userService := _userService.CreateUserUsecaseImpl(userRepository)
	userHandler.CreateUserHandler(route, userService)

	eventRepoUse := eventRepo.CreateEventRepoImpl(dbConn, sqlConn)
	eventServiceUse := eventService.CreateEventService(eventRepoUse, bannerRepoUse)
	eventHandler.CreateEventHandler(route, eventServiceUse)

	eoRepo := eoRepo.CreateEoRepoImpl(dbConn)
	eoService := eoService.CreateEoService(eoRepo, userRepository)
	eoDelivery.CreateEoHandler(route, eoService)

	fmt.Println("Starting Web Server At Port : " + port)
	err = http.ListenAndServe(": "+port, route)
	if err != nil {
		log.Fatal(err)
	}

}
