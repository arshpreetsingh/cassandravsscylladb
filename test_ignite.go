package main

import (
	"database/sql"

	"fmt"

	_ "github.com/amsokol/ignite-go-client/sql"
)

func InitIgnite() {

	// open connection
	//protocol://host:port/cache?param1=value1&param2=value2&paramN=valueN
	fmt.Println("connection started for Apache ignite")
	db, err := sql.Open("ignite", "tcp://apacheignite:10800/ExampleDB?version=1.1.0&timeout=5000")
	if err != nil {
		fmt.Println("failed to open connection: %v", err)
	}
	fmt.Println("connection made for ignite")
	createTable := "CREATE TABLE user_test ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);"
	result, eror := db.Exec(createTable)
	if eror != nil {
		fmt.Println("Error! IN creating Table for TimescaleDB  " + eror.Error())
	}
	fmt.Println(result)
	defer db.Close()
}
