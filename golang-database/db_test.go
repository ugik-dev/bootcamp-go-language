package golangdatabase

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	fmt.Println("Test Connection Running")
	db, err := sql.Open("mysql", "root:''@tcp(localhost:3306)/go_learn")
	// db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!!")
	defer db.Close()
	// See "Important settings" section.
	// beralam lama sebuah koneksi bertahan
	db.SetConnMaxLifetime(time.Minute * 3)

	// berapa lama koenski yang sudah tidak gunakan akan dihapus
	db.SetConnMaxIdleTime(time.Minute * 3)

	// max connection
	db.SetMaxOpenConns(50)

	// berapa jumlah conecction dibuat untuk berjaga" atau standby jika trafik naik
	db.SetMaxIdleConns(10)

}
