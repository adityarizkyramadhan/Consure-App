package main

import (
	expertCtrl "Consure-App/controller/expert"
	reviewCtrl "Consure-App/controller/review"
	userCtrl "Consure-App/controller/user"
	"Consure-App/domain"
	generalRepo "Consure-App/repository/general/general_impl"
	userRepo "Consure-App/repository/user/user_impl"

	expertRepo "Consure-App/repository/expert/expert_impl"
	expertUc "Consure-App/usecase/expert/expert_impl"
	userUc "Consure-App/usecase/user/user_impl"

	reviewRepo "Consure-App/repository/review/review_impl"
	reviewUc "Consure-App/usecase/review/review_impl"
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
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong cd ke tiga")
	})

	repoGeneral := generalRepo.NewGeneralRepositoryImpl(db)

	//User
	repoUser := userRepo.NewUserRepositoryImpl(db)
	ucUser := userUc.NewUserUsecaseImpl(repoUser, repoGeneral)
	user := router.Group("user")
	userCtrl.NewUserController(ucUser, user)

	//Expert
	repoExpert := expertRepo.NewExpertRepository(db)
	ucExpert := expertUc.NewExpertUsecase(repoGeneral, repoExpert)
	expert := router.Group("expert")
	expertCtrl.NewExpertController(ucExpert, expert)

	//Review
	repoReview := reviewRepo.NewReviewRepository(db)
	ucReview := reviewUc.NewReviewUsecase(repoGeneral, repoReview)
	review := router.Group("review")
	reviewCtrl.NewController(ucReview, review)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err.Error())
	}

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
