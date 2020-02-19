package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	DB_USER     = "maropost"
	DB_PASSWORD = "Maro123!"
	DB_NAME     = "test"
)

func InitTimeScale() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	DB, err := sql.Open("postgres", dbinfo)
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
	result2, eror := DB.Exec(createTable)
	if eror != nil {
		fmt.Println("Error! IN creating Table for TimescaleDB  " + eror.Error())
	}
	fmt.Println(result2)
	hyperTable := "SELECT create_hypertable('user_test', 'created_at');"
	result3, eror := DB.Exec(hyperTable)
	if eror != nil {
		fmt.Println("Error! IN creating Hypertable " + eror.Error())
	}
	fmt.Println(result3)
}

func StoreDataTimeScaleDB(count int) error {
	startTime := time.Now()

	query := "INSERT INTO bounce_reports(recorded_at,account_id,to_email,from_email,to_domain,from_domain,injection_ip,remote_ip,ip_pool,event) VALUES ('" + data["time_stamp"] + "','" + data["account_id"] + "','" + data["rcpt"] + "','" + data["from"] + "','" + toDomain + "','" + fromDomain + "','" + data["ip"] + "','" + data["remote_ip"] + "','ip_pool','" + data["event"] + "');"
	_, err := DB.Query(query)
	if err != nil {
		ServiceLogger.Println("Error! ocurred while data inserting: %v", err)
	}

	for i := 0; i < count; i++ {
		data := GenerateData()
		if err := session_cassandra.Query(`INSERT INTO test.user ( account_id, name, full_name,product_name,email,
     email_subject, email_body,user_agent, company, domain_name,gender,language,
    created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
			data.AccountID, data.Name, data.FullName, data.ProductName, data.Email, data.EmailSubject,
			data.EmailBody, data.UserAgent, data.Company, data.DomainName, data.Gender, data.Language, data.CreatedAt, data.Updatedat).Exec(); err != nil {
			log.Println("Error! unable to insert data into scylladb", err)
			return err
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
