package golangdatabase

import (
	"database/sql"
	"fmt"
	"time"
)

/*
*

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/dbname")
		// beralam lama sebuah koneksi bertahan
	db.SetConnMaxLifetime(time.Minute * 30)

	// berapa lama koenski yang sudah tidak gunakan akan dihapus
	db.SetConnMaxIdleTime(time.Minute * 3)

	// max connection
	db.SetMaxOpenConns(50)

	// berapa jumlah conecction dibuat untuk berjaga" atau standby jika trafik naik
	db.SetMaxIdleConns(10)
*/
func GetConnection() *sql.DB {
	fmt.Println("Get Connection Running")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_learn?parseTime=true")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!!")
	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	return db
}
