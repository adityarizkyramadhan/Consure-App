package main

import (
	expertCtrl "Consure-App/controller/expert"
	reviewCtrl "Consure-App/controller/review"
	trxCtrl "Consure-App/controller/transaction"
	userCtrl "Consure-App/controller/user"
	"Consure-App/domain"
	generalRepo "Consure-App/repository/general/general_impl"
	userRepo "Consure-App/repository/user/user_impl"

	expertRepo "Consure-App/repository/expert/expert_impl"
	expertUc "Consure-App/usecase/expert/expert_impl"
	userUc "Consure-App/usecase/user/user_impl"

	transactionRepo "Consure-App/repository/transaction/transaction_impl"
	transactionUc "Consure-App/usecase/transaction/transaction_impl"

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
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	driverDb, err := ReadEnvDatabase()
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

	//Transaksi
	repoTrc := transactionRepo.NewTransactionRepository(db)
	ucTrc := transactionUc.NewTransactionUsecase(repoGeneral, repoTrc)
	transaction := router.Group("transaction")
	trxCtrl.NewTransactionController(ucTrc, transaction)

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

func ReadEnvDatabase() (DriverSupabase, error) {
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
	if err := db.AutoMigrate(new(domain.User), new(domain.Expert), new(domain.Review), new(domain.Transaction)); err != nil {
		return nil, err
	}
	return db, nil
}

// func main() {
// 	//2022-10-25 15:00:00.000
// 	dateString := "2022-10-25 15:00:00.000"
// 	date, error := time.Parse("2006-01-02 15:04:05.000", dateString)

// 	if error != nil {
// 		fmt.Println(error)
// 		return
// 	}

// 	fmt.Printf("Type of dateString: %T\n", dateString)
// 	fmt.Printf("Type of date: %d\n", date.Day())
// 	fmt.Println()
// 	fmt.Println(date.Hour())
// 	fmt.Printf("Value of dateString: %v\n", dateString)
// 	fmt.Printf("Value of date: %v", date)
// }
