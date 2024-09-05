package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occurred while loading .env")
	}
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := "postgres"
	pass := os.Getenv("POSTGRES_PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("Error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
	SetUpDatabase()
	dbname = os.Getenv("POSTGRES_DB_NAME")
	psqlSetup = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql = sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("Error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
	SetUpTables()
}

func SetUpDatabase() {
	_, err := Db.Exec("CREATE DATABASE " + os.Getenv("POSTGRES_DB_NAME"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table created successfully!")
	}
}

func SetUpTables() {
	_, err := Db.Exec("CREATE TABLE " + os.Getenv("POSTGRES_TABLE_NAME") + " (data VARCHAR)")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table created successfully!")
	}
}

type ReviewStatistics struct {
	Data string
}

func AddReviewStatistics(data string) {
	data_struct := ReviewStatistics{Data: data}

	_, err := Db.Exec("insert into reviewstatistics(data) values ($1)", data_struct.Data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data inserted successfully!")
	}

}
