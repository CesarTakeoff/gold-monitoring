package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ResulStr struct {
	IntDate      string
	IntTimestamp string
	IntId        string
	IntName      string
	IntData      string
}
type resulOutput []ResulStr

// func main() {

// 	db := postgressConnection()
// 	defer db.Close()
// 	rec := getResult(db)
// 	fmt.Println(rec)

// 	addMonitoringResul(db, resulStr{})

// }

func PostgressConnection(dbParams map[string]string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbParams["host"], dbParams["port"], dbParams["user"], dbParams["password"], dbParams["dbname"])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func AddMonitoringResul(db *sql.DB, goldMonitoringValues ResulStr) {

	Qs := fmt.Sprintf("insert into gold_interface_data values ('%s', '%s', '%s', '%s','%s');",
		goldMonitoringValues.IntId,
		goldMonitoringValues.IntDate,
		goldMonitoringValues.IntTimestamp,
		goldMonitoringValues.IntId,
		goldMonitoringValues.IntData)
	fmt.Println(Qs)
	_, err := db.Exec(Qs)
	if err != nil {
		panic(err)
	}
}

func GetResult(db *sql.DB) resulOutput {
	Qs := fmt.Sprintf("Select * from gold_interface_data")

	rows, err := db.Query(Qs)
	if err != nil {
		log.Fatal("Couln't get data from table")
	}
	defer rows.Close()
	retVal := resulOutput{}
	for rows.Next() {
		fields := ResulStr{}
		err := rows.Scan(&fields.IntId, &fields.IntDate, &fields.IntTimestamp, &fields.IntName, &fields.IntData)
		if err != nil {
			log.Fatal("Error scanning rows", err)
		}
		retVal = append(retVal, fields)
	}
	return retVal
}
