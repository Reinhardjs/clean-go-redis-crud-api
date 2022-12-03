package configs

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var redisClient redis.Conn

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// Init PostgreDB connection
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	sslMode := os.Getenv("sslmode")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName, sslMode, password)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(dbUri)
	db = conn

	// Init redis connection
	redisConnection, err := redis.Dial("tcp", os.Getenv("redis_host")+":6379")
	redisClient = redisConnection

	if err != nil {
		fmt.Print(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func GetRedis() redis.Conn {
	return redisClient
}
