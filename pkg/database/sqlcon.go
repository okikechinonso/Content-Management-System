package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func MySqlCon() (db *sql.DB, err error) {

	//host := os.Getenv("HOST")
	//port := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", "postgres://megqyyfdthyzrz:0fdfa9b60a54af403c68817f80eef83f5999248222e00dd1a3b160d3760d3ec6@ec2-54-209-165-105.compute-1.amazonaws.com:5432/da92eirikddah1")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, nil
}
