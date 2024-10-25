package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"test/wanikani"

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
	_, err := Db.Exec("CREATE TABLE reviewstatistics (ID VARCHAR, DataUpdatedAt VARCHAR, SubjectID VARCHAR, CreatedAt VARCHAR, SubjectType VARCHAR, MeaningCorrect VARCHAR, MeaningIncorrect VARCHAR, MeaningMaxStreak VARCHAR, MeaningCurrentStreak VARCHAR, ReadingCorrect VARCHAR, ReadingIncorrect VARCHAR, ReadingMaxStreak VARCHAR, ReadingCurrentStreak VARCHAR, PercentageCorrect VARCHAR, Hidden VARCHAR)")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table created successfully!")
	}
}

func SaveReviewStatisticsToDB(data string) {
	var response wanikani.WaniKaniResponse
	json.Unmarshal([]byte(data), &response)

	SetUpTables()

	for _, review_entry := range response.ReviewEntries {
		_, err := Db.Exec("insert into reviewstatistics(ID, DataUpdatedAt, SubjectID, CreatedAt, SubjectType, MeaningCorrect, MeaningIncorrect, MeaningMaxStreak, MeaningCurrentStreak, ReadingCorrect, ReadingIncorrect, ReadingMaxStreak, ReadingCurrentStreak, PercentageCorrect, Hidden) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", review_entry.ID, review_entry.DataUpdatedAt, review_entry.Data.SubjectID, review_entry.Data.CreatedAt, review_entry.Data.SubjectType, review_entry.Data.MeaningCorrect, review_entry.Data.MeaningIncorrect, review_entry.Data.MeaningMaxStreak, review_entry.Data.MeaningCurrentStreak, review_entry.Data.ReadingCorrect, review_entry.Data.ReadingIncorrect, review_entry.Data.ReadingMaxStreak, review_entry.Data.ReadingCurrentStreak, review_entry.Data.PercentageCorrect, review_entry.Data.Hidden)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Data inserted successfully!")
		}
	}

}
