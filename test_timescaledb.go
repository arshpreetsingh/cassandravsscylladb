package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "password"

	DB_NAME = "postgres"
)

func InitTimeScale() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescale")
	db := sqlx.MustConnect("postgres", dbinfo)
	defer db.Close()

	createTable := "CREATE TABLE user_test ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);"
	res, err := db.Queryx(createTable)
	if err != nil {
		panic(err)
	}
	res.Close()
	hyperTable := "SELECT create_hypertable('user_test', 'created_at');"
	res2, err := db.Queryx(hyperTable)
	if err != nil {
		panic(err)
	}
	res2.Close()
}

func StoreDataTimeScaleDB(count int) (float64, error) {
	startTime := time.Now()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescale")
	db := sqlx.MustConnect("postgres", dbinfo)
	defer db.Close()
	data := GenerateData()
	// query := fmt.Sprintf("INSERT INTO user_test ( account_id, name, full_name,product_name,email, email_subject, email_body,user_agent, company, domain_name,gender,language) VALUES (%d, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", data.AccountID, data.Name, data.FullName,
	// 	data.ProductName, data.Email, data.EmailSubject, data.EmailBody, data.UserAgent, data.Company, data.DomainName, data.Gender, data.Language)
	query := fmt.Sprintf("INSERT INTO user_test(account_id,name,full_name,product_name,email,email_body,user_agent, company, domain_name,gender,language,created_at) VALUES (%d,'username','full_name','productName','email_bodyy','user agenttttt','maropost','maropost.com','male','english','hello',Now());", data.AccountID)
	res, err := db.Queryx(query)
	if err != nil {
		return 2.2, err
	}
	res.Close()
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	return diff, nil
}

func StoreDataTimeScaleDBWorker(wg *sync.WaitGroup) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescale")
	db := sqlx.MustConnect("postgres", dbinfo)
	defer db.Close()
	for job := range Jobs {
		data := GenerateData()
		query := fmt.Sprintf("INSERT INTO user_test(account_id,name,full_name,product_name,email,email_body,user_agent, company, domain_name,gender,language) VALUES (%d,'username','full_name','productName','email_bodyy','user agenttttt','maropost','maropost.com','male','english','hello');", data.AccountID)
		res, err := db.Queryx(query)
		if err != nil {
			fmt.Println("Error! unable to insert data into cassandra", err)
		}
		res.Close()
		time.Sleep(1 * time.Second)
		output := JobResult{job, err}
		Results <- output
	}
	wg.Done() // Only return when Done...! Means when "results" channel will receive Value!!
}

func FetchDataTimescaledb(accountid int) (float64, error) {
	startTime := time.Now()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescale")
	db := sqlx.MustConnect("postgres", dbinfo)
	defer db.Close()
	query := fmt.Sprintf("SELECT account_id, name,email,company from user_test WHERE account_id = %d", accountid)
	res, err := db.Queryx(query)
	if err != nil {
		return 2.2, err
	}
	res.Close()
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	return diff, nil
}

func FetchDataTimescaledbComplex(accountid int) {
	var dbcon *sql.DB
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescale")
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
