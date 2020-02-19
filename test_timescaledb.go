package main

import (
  "database/sql"
  "fmt"
_ "github.com/lib/pq"
"log"
)

var DB *sql.DB

const (
	DB_USER     = "maropost"
	DB_PASSWORD = "Maro123!"
	DB_NAME     = "postgres"
)

func StartTimeScale() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, "timescaledb")
	DB, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic("Error! Unable to Connect to DB:" + err.Error())
	}

  DBqry := "create database test"
  _, err := DB.Exec(DBqry)
  if err != nil {
   fmt.Println("Error! " + err.Error())
 }
createTable:="CREATE TABLE IF NOT EXISTS user ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);"
_, eror := DB.Exec(createTable)
if eror != nil {
 fmt.Println("Error!  " + eror.Error())
}
hyperTable:= "SELECT create_hypertable('user', 'created_at');"
_, eror := DB_obj.Exec(hyperTable)
if eror != nil {
  fmt.Println("Error! " + eror.Error())
}
}
