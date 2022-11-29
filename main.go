package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

}

type DriverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func ReadEnvSupabase() (DriverSupabase, error) {
	envSupabase, err := godotenv.Read()
	if err != nil {
		return DriverSupabase{}, err
	}
	return DriverSupabase{
		User:     envSupabase["SUPABASE_USER"],
		Password: envSupabase["SUPABASE_PASSWORD"],
		Host:     envSupabase["SUPABASE_HOST"],
		Port:     envSupabase["SUPABASE_PORT"],
		DbName:   envSupabase["SUPABASE_DB_NAME"],
	}, nil
}

func MakeConnection(data DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(); err != nil {
		return nil, err
	}
	return db, nil
}
