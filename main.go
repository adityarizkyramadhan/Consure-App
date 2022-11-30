package main

import (
	userCtrl "Consure-App/controller/user"
	"Consure-App/domain"
	generalRepo "Consure-App/repository/general/general_impl"
	userRepo "Consure-App/repository/user/user_impl"
	userUc "Consure-App/usecase/user/user_impl"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	driverDb, err := ReadEnvSupabase()
	if err != nil {
		panic(err.Error())
	}
	db, err := MakeConnection(driverDb)
	if err != nil {
		panic(err.Error())
	}

	//User
	repoGeneral := generalRepo.NewGeneralRepositoryImpl(db)
	repoUser := userRepo.NewUserRepositoryImpl(db)
	ucUser := userUc.NewUserUsecaseImpl(repoUser, repoGeneral)
	user := router.Group("user")
	userCtrl.NewUserController(ucUser, user)
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

type DriverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func ReadEnvSupabase() (DriverSupabase, error) {
	return DriverSupabase{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}

func MakeConnection(data DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(new(domain.User), new(domain.Expert), new(domain.Review)); err != nil {
		return nil, err
	}
	return db, nil
}
