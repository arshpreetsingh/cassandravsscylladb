package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// var DB *sql.DB

const (
	DB_USER     = "maropost"
	DB_PASSWORD = "Maro123!"
	DB_NAME     = "test"
)

func InitTimeScale() {
	var dbcon *sql.DB

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	dbcon, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic("Error! Unable to Connect to DB:" + err.Error())
	}

	// result, err := DB.Exec("create database test")
	// if err != nil {
	// 	fmt.Println("Error! " + err.Error())
	// }
	//
	// fmt.Println(result)

	//createTable := "CREATE TABLE user_test (recorded_at TIMESTAMPTZ NOT NULL,account_id INTEGER,to_email TEXT,from_email TEXT,to_domain TEXT,from_domain TEXT,injection_ip TEXT,remote_ip TEXT,ip_pool TEXT,event TEXT);"
	createTable := "CREATE TABLE user_test ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);"
	result2, eror := dbcon.Exec(createTable)
	if eror != nil {
		fmt.Println("Error! IN creating Table for TimescaleDB  " + eror.Error())
	}
	fmt.Println(result2)
	hyperTable := "SELECT create_hypertable('user_test', 'created_at');"
	result3, eror := dbcon.Exec(hyperTable)
	if eror != nil {
		fmt.Println("Error! IN creating Hypertable " + eror.Error())
	}
	fmt.Println(result3)
}

func StoreDataTimeScaleDB(count int) error {
	var dbcon *sql.DB
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	dbcon, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic("Error! Unable to Connect to DB:" + err.Error())
	}
	startTime := time.Now()
	for i := 0; i < count; i++ {
		data := GenerateData()
		rows, err := dbcon.Query(`INSERT INTO user_test ( account_id, name, full_name,product_name,email,
       email_subject, email_body,user_agent, company, domain_name,gender,language,
      created_at, updated_at) VALUES ($1,$2, $3, $4, $5, $6, $7,$8,$9,$10,$11,$12,$13,$14)`,
			data.AccountID, data.Name, data.FullName, data.ProductName, data.Email, data.EmailSubject,
			data.EmailBody, data.UserAgent, data.Company, data.DomainName, data.Gender, data.Language, data.CreatedAt, data.Updatedat)
		if err != nil {
			log.Println("Error! unable to insert data into TimeScaledb", err)
			return err
		}
		for rows.Next() {
			fmt.Println(rows)
			// err := rows.Scan(&accountid)
			// if err != nil {
			// 	fmt.Printf("Not able to fetch required data for classification: %v", err)
			// }
		}
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	fmt.Println("Write Operation Finished for Count::" + "   " + string(count) + "   " + "in Following Seconds")
	fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
	return nil
}

func FetchDataTimescaledb(accountid int) {
	var dbcon *sql.DB
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	dbcon, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic("Error! Unable to Connect to DB:" + err.Error())
	}
	startTime := time.Now()
	rows, err := dbcon.Query(`SELECT account_id, name, full_name,product_name,email,email_subject, email_body,user_agent, company,
    domain_name,gender,language, created_at from user_test WHERE account_id = $1`, accountid)
	if err != nil {
		fmt.Printf("Not able to fetch data from Database: %v", err)
	}
	fmt.Println("rows")
	for rows.Next() {
		fmt.Println("rows")
		fmt.Println(rows)
		// err := rows.Scan(&accountid)
		// if err != nil {
		// 	fmt.Printf("Not able to fetch required data for classification: %v", err)
		// }
		//	fmt.Println("account_ID", accountid)
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	fmt.Println("Read Operation Finished in Following Seconds")
	fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
}

func FetchDataTimescaledbComplex(accountid int) {
	var dbcon *sql.DB
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	dbcon, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic("Error! Unable to Connect to DB:" + err.Error())
	}
	startTime := time.Now()
	rows, err := dbcon.Query(`SELECT COUNT(account_id,email,company) from user_test WHERE account_id = $1 created_at>$2 created_at<$2 GROUP BY account_id, created_at`, accountid, startTime)
	if err != nil {
		fmt.Printf("Not able to fetch data from Database: %v", err)
	}
	for rows.Next() {
		fmt.Println(rows)
		// err := rows.Scan(&accountid)
		// if err != nil {
		// 	fmt.Printf("Not able to fetch required data for classification: %v", err)
		// }
		//	fmt.Println("account_ID", accountid)
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	fmt.Println("Write Operation Finished in Following Seconds")
	fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
}
