package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	// model1
	// query := "select * from customer"
	// rows, err := db.QueryContext(ctx, query)

	query := "select * from customer where email = ? and year(dob) = ?"
	rows, err := db.QueryContext(ctx, query, "zea@dev.com", "2022")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, balance int
		var name string
		var email sql.NullString
		var dob sql.NullTime
		var createdAt time.Time
		var rating float64
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &dob, &married)
		if err != nil {
			panic(err)
		}
		fmt.Printf("------\nID : %d\nName : %s\nBalance : %d\nRating : %.2f\nCreated At : %s\nMarried : %t\n",
			id, name, balance, rating, createdAt.Format("2006-01-02 15:04:05"), married)

		if dob.Valid {
			fmt.Printf("DOB : %s\n", dob.Time.Format("2006-01-02"))
		} else {
			fmt.Println("DOB : NULL")
		}

		if email.Valid {
			fmt.Printf("Email : %s\n", email.String)
		} else {
			fmt.Println("Email : NULL")
		}
	}
	rows.Close()
}

func TestGetUsers(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	// model1
	query := "select * from users"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var password, username string

		err := rows.Scan(&id, &username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Printf("------\nID : %d\nUsername : %s\nPassword : %s\n",
			id, username, password)
	}
	rows.Close()
}
func TestCreate(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	// query := `insert into customer (name, email, balance, rating, dob, married) values
	// ('Fazea Audrilia','zea@dev.com', 10000, 95.0, '2021-10-10', false),
	// ('Yuvita','yuvita@dev.com', 900000, 92.0, '1996-10-10', true),
	// ('Ugik Dev','ugik.dev@dev.com', 500000, 60.0, '1996-10-10', true)`
	query := `insert into customer (name, email, balance, rating, dob, married) values 
	('Kakung',null, 10000, 81.3, null, false)`
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Query Finis")
}

func TestCreateGood(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()
	username := "admin2"
	password := "pwd"
	query := `insert into users (username, password) values 
	( ?, ?)`
	result, err := db.ExecContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Insert ID ", insertId)
}

func TestCreateGoodTwo(t *testing.T) {
	//jika mengunakan prepair statement lebih hemat jika data banyak sekali eksekusi
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	query := `insert into users (username, password) values 
	( ?, ?)`
	statement, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 1; i <= 10; i++ {
		username := "user" + strconv.Itoa(i)
		password := "pwdu" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, username, password)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("looping", i, "Last Insert ID ", insertId)
	}
	// insertId, err := result.LastInsertId()

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	query := `insert into users (username, password) values 
	( ?, ?)`
	// transaction
	for i := 31; i <= 35; i++ {
		username := "user" + strconv.Itoa(i)
		password := "pwdu" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, query, username, password)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("looping", i, "Last Insert ID ", insertId)
	}
	// end transactiin

	// err = tx.Commit()

	// untuk membatalkan semua transaksi
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
