package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Madhavi@4"
	dbname   = "Data"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close database
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return // Exit the program if the connection fails
	}

	fmt.Println("Connected!")
	// //hardCoded
	// insertData := `insert into "students"("name", "roll_number") values('Jacob', 20)`
	// _, err1 := db.Exec(insertData)
	// //Dynamic
	// insertData1 := `insert into "students"("name", "roll_number") values($1, $2)`
	// _, err2 := db.Exec(insertData1, "Madhavi", 1)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// if err2 != nil {	// 	fmt.Println(err2)
	// 	return
	// }

	updateData := `update "students" set "name"=$1 where "roll_number"=$2`
	_, err3 := db.Exec(updateData, "Madhavi Asam", 20)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	deleteData := `delete from "students" where "roll_number"=$1`
	_, err4 := db.Exec(deleteData, 20)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
}
